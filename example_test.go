package intercom_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-intercom"
)

//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := intercom.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("INTERCOM_ACCESS_TOKEN"),
			)
			req.Header.Set("Intercom-Version", "2.17")

			return nil
		},
	)

	client, err := intercom.NewClientWithResponses(
		"https://api.intercom.io",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.IdentifyAdminWithResponse(
		context.Background(),
		&intercom.IdentifyAdminParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
