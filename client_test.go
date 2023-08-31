package goperset_test

import (
	"github.com/justmike1/goperset"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := goperset.NewClient("https://superset.domain.net/")
	t.Logf(client.BasePath)
}

func TestGetAccessTokens(t *testing.T) {
	client := goperset.NewClient("https://superset.domain.net/") // Client also has a context
	authToken, csrfToken, err := goperset.GetAccessTokens(client, "admin", "admin")
	if err != nil {
		t.Errorf("Error getting access token: %v", err)
		return
	}
	t.Logf("Auth token: %s", authToken)
	t.Logf("CSRF token: %s", csrfToken)
}
