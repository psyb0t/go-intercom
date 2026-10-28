package intercom_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-intercom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAccessToken = "EXAMPLE-DO-NOT-USE"
	testAPIVersion  = "2.17"
	testContactID   = "contact-example"
	contentTypeJSON = "application/json"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *intercom.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := intercom.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAccessToken)
			req.Header.Set("Intercom-Version", testAPIVersion)

			return nil
		},
	)

	client, err := intercom.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertAuth(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAccessToken, r.Header.Get("Authorization"))
	assert.Equal(t, testAPIVersion, r.Header.Get("Intercom-Version"))
}

func writeJSON(t *testing.T, w http.ResponseWriter, body string) {
	t.Helper()

	w.Header().Set("Content-Type", contentTypeJSON)

	_, err := w.Write([]byte(body))
	assert.NoError(t, err)
}

func TestIdentifyAdmin(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/me", r.URL.Path)
		assertAuth(t, r)

		writeJSON(t, w, `{
			"type":"admin","id":"admin-example","name":"Example Admin",
			"email":"admin@example.com","email_verified":true,
			"app":{"type":"app","id_code":"workspace-example"}
		}`)
	})

	res, err := client.IdentifyAdminWithResponse(
		context.Background(),
		&intercom.IdentifyAdminParams{},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, "admin-example", *res.JSON200.Id)
	assert.Equal(t, "admin@example.com", *res.JSON200.Email)
	assert.True(t, *res.JSON200.EmailVerified)
	require.NotNil(t, res.JSON200.App)
	assert.Equal(t, "workspace-example", *res.JSON200.App.IdCode)
}

func TestShowContact(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/contacts/"+testContactID, r.URL.Path)
		assert.Equal(t, "true", r.URL.Query().Get("include_merge_history"))
		assertAuth(t, r)

		writeJSON(t, w, `{
			"type":"contact","id":"contact-example","role":"user",
			"email":"person@example.com","name":null,"phone":null,
			"created_at":1571672154,"updated_at":1571672154
		}`)
	})

	includeMergeHistory := true
	res, err := client.ShowContactWithResponse(
		context.Background(),
		testContactID,
		&intercom.ShowContactParams{
			IncludeMergeHistory: &includeMergeHistory,
		},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, testContactID, *res.JSON200.Id)
	assert.Equal(t, "person@example.com", *res.JSON200.Email)
	assert.Equal(t, 1571672154, *res.JSON200.CreatedAt)
	assert.Nil(t, res.JSON200.Name)
	assert.Nil(t, res.JSON200.Phone)
}

func TestCreateContact(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/contacts", r.URL.Path)
		assert.Equal(t, contentTypeJSON, r.Header.Get("Content-Type"))
		assertAuth(t, r)

		var body intercom.CreateContactRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))

		if !assert.NotNil(t, body.Email) || !assert.NotNil(t, body.Role) {
			return
		}

		assert.Equal(t, "person@example.com", *body.Email)
		assert.Equal(t, "user", *body.Role)

		writeJSON(t, w, `{
			"type":"contact","id":"contact-example",
			"role":"user","email":"person@example.com"
		}`)
	})

	email := "person@example.com"
	role := "user"
	res, err := client.CreateContactWithResponse(
		context.Background(),
		&intercom.CreateContactParams{},
		intercom.CreateContactJSONRequestBody{Email: &email, Role: &role},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Equal(t, testContactID, *res.JSON200.Id)
	assert.Equal(t, email, *res.JSON200.Email)
}

func TestUpdateContact(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/contacts/"+testContactID, r.URL.Path)
		assert.Equal(t, contentTypeJSON, r.Header.Get("Content-Type"))
		assertAuth(t, r)

		var body intercom.UpdateContactRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))

		if !assert.NotNil(t, body.Name) {
			return
		}

		assert.Equal(t, "Updated name", *body.Name)
		writeJSON(t, w, `{
			"type":"contact","id":"contact-example","name":"Updated name"
		}`)
	})

	name := "Updated name"
	res, err := client.UpdateContactWithResponse(
		context.Background(),
		testContactID,
		&intercom.UpdateContactParams{},
		intercom.UpdateContactJSONRequestBody{Name: &name},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Equal(t, testContactID, *res.JSON200.Id)
	assert.Equal(t, name, *res.JSON200.Name)
}

func TestShowContactUnauthorized(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/contacts/"+testContactID, r.URL.Path)
		assertAuth(t, r)

		w.Header().Set("Content-Type", contentTypeJSON)
		w.WriteHeader(http.StatusUnauthorized)
		writeJSON(t, w, `{
			"type":"error.list",
			"errors":[{"code":"unauthorized","message":"Invalid token"}]
		}`)
	})

	res, err := client.ShowContactWithResponse(
		context.Background(),
		testContactID,
		&intercom.ShowContactParams{},
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.StatusCode())
	assert.Nil(t, res.JSON200)
	require.NotNil(t, res.JSON401)
	assert.Equal(t, "error.list", res.JSON401.Type)
	require.Len(t, res.JSON401.Errors, 1)

	code, err := res.JSON401.Errors[0].Code.AsErrorErrorsCode0()
	require.NoError(t, err)
	assert.Equal(t, "unauthorized", code)
}

func TestIdentifyAdminMalformedResponse(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(t, w, `{"type":`)
	})

	res, err := client.IdentifyAdminWithResponse(
		context.Background(),
		&intercom.IdentifyAdminParams{},
	)

	var syntaxError *json.SyntaxError

	require.ErrorAs(t, err, &syntaxError)
	assert.Nil(t, res)
}

func TestIdentifyAdminCanceled(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(_ http.ResponseWriter, _ *http.Request) {
		t.Error("a canceled request must not reach the server")
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	res, err := client.IdentifyAdminWithResponse(
		ctx,
		&intercom.IdentifyAdminParams{},
	)
	require.ErrorIs(t, err, context.Canceled)
	assert.Nil(t, res)
}

func TestReportingExportStatus(t *testing.T) {
	t.Parallel()

	const jobID = "job-example"

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/export/reporting_data/"+jobID, r.URL.Path)
		assertAuth(t, r)

		writeJSON(t, w, `{
			"job_identifier":"job-example","status":"complete",
			"download_url":"","download_expires_at":""
		}`)
	})

	res, err := client.GetExportReportingDataJobIdentifierWithResponse(
		context.Background(),
		jobID,
		&intercom.GetExportReportingDataJobIdentifierParams{},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)
	assert.Equal(t, jobID, *res.JSON200.JobIdentifier)
	assert.EqualValues(t, "complete", *res.JSON200.Status)
}

func TestReportingDownload(t *testing.T) {
	t.Parallel()

	const (
		jobID = "job-example"
		body  = "id,value\nexample,1\n"
	)

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/download/reporting_data/"+jobID, r.URL.Path)
		assert.Equal(t, "application/octet-stream", r.Header.Get("Accept"))
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/octet-stream")

		_, err := w.Write([]byte(body))
		assert.NoError(t, err)
	})

	const accept = intercom.
		GetDownloadReportingDataJobIdentifierParamsAcceptApplicationoctetStream

	params := intercom.GetDownloadReportingDataJobIdentifierParams{
		Accept: accept,
	}
	res, err := client.GetDownloadReportingDataJobIdentifierWithResponse(
		context.Background(),
		jobID,
		&params,
	)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode())
	assert.Equal(t, body, string(res.Body))
}

func TestPatchedEnumWireNames(t *testing.T) {
	t.Parallel()

	var attributes intercom.ConversationAttributeList

	require.NoError(t, json.Unmarshal(
		[]byte(`{"type":"list","data":[]}`),
		&attributes,
	))
	require.NotNil(t, attributes.Type)
	assert.Equal(t,
		intercom.ConversationAttributeListObjectTypeList,
		*attributes.Type,
	)

	var ticket intercom.Ticket

	require.NoError(t, json.Unmarshal(
		[]byte(`{"type":"ticket","id":"ticket-example"}`),
		&ticket,
	))
	require.NotNil(t, ticket.Type)
	assert.Equal(t, intercom.TicketObjectTypeTicket, *ticket.Type)

	encoded, err := json.Marshal(ticket)
	require.NoError(t, err)
	assert.JSONEq(t,
		`{"type":"ticket","id":"ticket-example"}`,
		string(encoded),
	)
}

func TestUnionRequestBodyJSON(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		body any
		json string
	}{
		{
			name: "escalate existing conversation",
			body: new(intercom.EscalateFinConversationJSONRequestBody),
			json: `{
				"conversation_id":"conversation-example",
				"context":"Please transfer this conversation"
			}`,
		},
		{
			name: "escalate user",
			body: new(intercom.EscalateFinConversationJSONRequestBody),
			json: `{
				"user":{"id":"user-example","email":"person@example.com"},
				"context":"Please transfer this user"
			}`,
		},
		{
			name: "close conversation",
			body: new(intercom.ManageConversationJSONRequestBody),
			json: `{
				"type":"admin","admin_id":"admin-example",
				"message_type":"close","body":"Completed"
			}`,
		},
		{
			name: "create tag",
			body: new(intercom.CreateTagJSONRequestBody),
			json: `{"name":"example-tag"}`,
		},
		{
			name: "admin ticket reply",
			body: new(intercom.ReplyTicketJSONRequestBody),
			json: `{
				"type":"admin","admin_id":"admin-example",
				"message_type":"note","body":"Investigating",
				"skip_notifications":true
			}`,
		},
		{
			name: "contact ticket reply",
			body: new(intercom.ReplyTicketJSONRequestBody),
			json: `{
				"type":"user","user_id":"user-example",
				"message_type":"comment","body":"Additional context"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			require.NoError(t, json.Unmarshal([]byte(tc.json), tc.body))

			encoded, err := json.Marshal(tc.body)
			require.NoError(t, err)
			assert.JSONEq(t, tc.json, string(encoded))
		})
	}
}
