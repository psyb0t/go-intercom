// Command patch-spec adapts Intercom schemas for Go client generation.
package main

import (
	"maps"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	outputSpecPath = "api/openapi.fixed.yaml"
	outputSpecMode = 0o600
	postMethod     = "post"
)

func main() {
	input, err := os.ReadFile("api/openapi.yaml")
	if err != nil {
		panic(err)
	}

	var spec map[string]any

	if err := yaml.Unmarshal(input, &spec); err != nil {
		panic(err)
	}

	patchPaths(spec)
	patchEnumNames(spec)
	patchContactBodies(spec)
	hoistUnionBodies(spec)

	output, err := yaml.Marshal(spec)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(outputSpecPath, output, outputSpecMode); err != nil {
		panic(err)
	}
}

func patchPaths(spec map[string]any) {
	paths := objectAt(spec, "paths")

	// These placeholders lack declarations in the upstream document.
	// Required string parameters let oapi-codegen retain both operations.
	for _, path := range []string{
		"/export/reporting_data/{job_identifier}",
		"/download/reporting_data/{job_identifier}",
	} {
		item, exists := paths[path]
		if !exists {
			continue
		}

		pathItem, ok := item.(map[string]any)
		if !ok {
			panic("reporting path must be an object")
		}

		params, _ := pathItem["parameters"].([]any)
		pathItem["parameters"] = append(params, map[string]any{
			"name":     "job_identifier",
			"in":       "path",
			"required": true,
			"schema":   map[string]any{"type": "string"},
		})
	}
}

func patchEnumNames(spec map[string]any) {
	// Inline type enums otherwise collide with the ticket_type and
	// conversation_attribute_list_type components. Wire names stay unchanged.
	for schema, typeName := range map[string]string{
		"conversation_attribute_list": "ConversationAttributeListObjectType",
		"ticket":                      "TicketObjectType",
	} {
		property := objectAt(
			spec,
			"components", "schemas", schema, "properties", "type",
		)
		property["x-go-type-name"] = typeName
	}
}

func objectAt(value map[string]any, keys ...string) map[string]any {
	for _, key := range keys {
		next, ok := value[key].(map[string]any)
		if !ok {
			panic("expected spec object at " + key)
		}

		value = next
	}

	return value
}

func patchContactBodies(spec map[string]any) {
	// A one-branch request union creates a distinct body type that loses
	// its JSON marshaler and sends {}. A direct reference keeps the same
	// schema while retaining the component's fields and marshaling methods.
	for path, method := range map[string]string{
		"/contacts":              postMethod,
		"/contacts/{contact_id}": "put",
	} {
		schema := objectAt(
			spec,
			"paths", path, method, "requestBody", "content",
			"application/json", "schema",
		)

		branches, ok := schema["oneOf"].([]any)
		if !ok || len(branches) != 1 {
			continue
		}

		branch, ok := branches[0].(map[string]any)
		if !ok {
			panic("contact body branch must be an object")
		}

		delete(schema, "oneOf")

		maps.Copy(schema, branch)
	}
}

func hoistUnionBodies(spec map[string]any) {
	schemas := objectAt(spec, "components", "schemas")

	// Inline unions lose their JSON methods when request body types are
	// defined. Component references produce aliases that retain those methods.
	for _, body := range []struct {
		path   string
		method string
		name   string
	}{
		{
			path:   "/fin/escalate",
			method: postMethod,
			name:   "client_body_escalate_fin_conversation",
		},
		{
			path:   "/conversations/{conversation_id}/parts",
			method: postMethod,
			name:   "client_body_manage_conversation",
		},
		{
			path:   "/tags",
			method: postMethod,
			name:   "client_body_create_tag",
		},
		{
			path:   "/tickets/{ticket_id}/reply",
			method: postMethod,
			name:   "client_body_reply_ticket",
		},
	} {
		media := objectAt(
			spec, "paths", body.path, body.method,
			"requestBody", "content", "application/json",
		)

		schema := objectAt(media, "schema")
		_, isOneOf := schema["oneOf"]
		_, isAnyOf := schema["anyOf"]

		if !isOneOf && !isAnyOf {
			continue
		}

		if _, exists := schemas[body.name]; exists {
			panic("request body component already exists: " + body.name)
		}

		schemas[body.name] = schema
		media["schema"] = map[string]any{
			"$ref": "#/components/schemas/" + body.name,
		}
	}
}
