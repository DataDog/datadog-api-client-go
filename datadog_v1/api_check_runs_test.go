package datadog_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"gotest.tools/assert"
)

func TestSubmitCheckRuns(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var checks []datadog.CheckRun
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&checks)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(400)
			fmt.Fprintf(w, `{"errors": ["Decode error: %s"]}`, err)
			return
		}
		for _, check := range checks {
			if check.Check == "" {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(400)
				fmt.Fprintln(w, `{"errors": ["Missing field: check"]}`)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "ok"}`)
	}))
	defer srv.Close()
	config := datadog.NewConfiguration()
	config.BasePath = srv.URL
	client := datadog.NewAPIClient(config)

	submitCRResp, _, err := client.CheckRunsApi.SubmitCheckRuns(context.TODO(), []datadog.CheckRun{
		datadog.CheckRun{
			Check: "app.up",
		},
	})
	assert.NilError(t, err)
	assert.Equal(t, "ok", submitCRResp.GetStatus())

	submitCRResp, _, err = client.CheckRunsApi.SubmitCheckRuns(context.TODO(), []datadog.CheckRun{
		datadog.CheckRun{},
	})
	respErr := err.(datadog.GenericOpenAPIError).Model().(datadog.Error400)
	assert.Equal(t, "Missing field: check", respErr.GetErrors()[0])
}
