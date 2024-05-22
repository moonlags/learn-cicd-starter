package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey sashaaaaaa")

	key, err := GetAPIKey(headers)
	if key != "sashaaaaaa" || err != nil {
		t.Fatalf(`GetAPIKey(headers) = %q, %v, want "sashaaaaaa", error`, key, err)
	}
}
