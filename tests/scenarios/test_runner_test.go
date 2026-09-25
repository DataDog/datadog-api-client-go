/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package scenarios

import (
	"compress/zlib"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestServerTransportDecompressesDeflateResponse(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if got := request.Header.Get("x-openapi-test-session"); got != "session-id" {
			t.Errorf("unexpected test session: %q", got)
		}
		writer.Header().Set("Content-Encoding", "deflate")
		compressed := zlib.NewWriter(writer)
		if _, err := compressed.Write([]byte(`{"ok":true}`)); err != nil {
			t.Errorf("write compressed response: %v", err)
		}
		if err := compressed.Close(); err != nil {
			t.Errorf("close compressed response: %v", err)
		}
	}))
	defer testServer.Close()

	request, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	response, err := (testServerTransport{session: "session-id"}).RoundTrip(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(body), `{"ok":true}`; got != want {
		t.Fatalf("unexpected response body: got %q, want %q", got, want)
	}
	if got := response.Header.Get("Content-Encoding"); got != "" {
		t.Fatalf("unexpected content encoding after decompression: %q", got)
	}
}
