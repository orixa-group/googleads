package googleads

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

type mockTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m *mockTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}

func TestCredentials_GetRequestMetadata(t *testing.T) {
	tests := []struct {
		name            string
		token           *oauth2.Token
		tokenErr        error
		developerToken  string
		loginCustomerID string
		want            map[string]string
		wantErr         bool
	}{
		{
			name:            "Success with loginCustomerID",
			token:           &oauth2.Token{AccessToken: "fake-access-token"},
			tokenErr:        nil,
			developerToken:  "dev-token-123",
			loginCustomerID: "123-456-7890",
			want: map[string]string{
				"authorization":     "Bearer fake-access-token",
				"developer-token":   "dev-token-123",
				"login-customer-id": "123-456-7890",
			},
			wantErr: false,
		},
		{
			name:            "Success without loginCustomerID",
			token:           &oauth2.Token{AccessToken: "fake-access-token"},
			tokenErr:        nil,
			developerToken:  "dev-token-123",
			loginCustomerID: "",
			want: map[string]string{
				"authorization":   "Bearer fake-access-token",
				"developer-token": "dev-token-123",
			},
			wantErr: false,
		},
		{
			name:            "Error while fetching token",
			token:           nil,
			tokenErr:        errors.New("token expired or invalid"),
			developerToken:  "dev-token-123",
			loginCustomerID: "",
			want:            nil,
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &credentials{
				tokenSource:     &mockTokenSource{token: tt.token, err: tt.tokenErr},
				developerToken:  tt.developerToken,
				loginCustomerID: tt.loginCustomerID,
			}

			got, err := c.GetRequestMetadata(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequestMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRequestMetadata() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentials_RequireTransportSecurity(t *testing.T) {
	c := &credentials{}
	if got := c.RequireTransportSecurity(); got != true {
		t.Errorf("RequireTransportSecurity() = %v, want %v", got, true)
	}
}

func TestConnect(t *testing.T) {
	ctx := context.Background()
	config := Config{
		ClientID:        "client-id",
		ClientSecret:    "client-secret",
		RefreshToken:    "refresh-token",
		DeveloperToken:  "dev-token",
		LoginCustomerID: "login-id",
	}

	cClient, err := Connect(ctx, config)
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer cClient.Close()

	if cClient == nil {
		t.Fatal("Connect returned nil client")
	}

	if cClient.conn == nil {
		t.Fatal("client connection is nil")
	}
}
