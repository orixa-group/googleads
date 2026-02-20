package googleads

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"
	grpccreds "google.golang.org/grpc/credentials"
)

const (
	target = "googleads.googleapis.com:443"
)

// Config holds the configuration for the Google Ads API connection.
type Config struct {
	ClientID        string
	ClientSecret    string
	RefreshToken    string
	DeveloperToken  string
	LoginCustomerID string
}

// Client is the struct that represents the Google Ads API client.
type Client struct {
	conn *grpc.ClientConn
}

// Connect creates a new Google Ads API client.
func Connect(ctx context.Context, config Config) (*Client, error) {
	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     google.Endpoint,
	}

	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(grpccreds.NewTLS(nil)),
		grpc.WithPerRPCCredentials(&credentials{
			tokenSource: oauthConfig.TokenSource(ctx, &oauth2.Token{
				RefreshToken: config.RefreshToken,
			}),
			developerToken:  config.DeveloperToken,
			loginCustomerID: config.LoginCustomerID,
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Google Ads: %w", err)
	}

	return &Client{conn}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

type credentials struct {
	tokenSource     oauth2.TokenSource
	developerToken  string
	loginCustomerID string
}

func (c *credentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get oauth token: %w", err)
	}

	md := map[string]string{
		"authorization":   "Bearer " + token.AccessToken,
		"developer-token": c.developerToken,
	}

	if c.loginCustomerID != "" {
		md["login-customer-id"] = c.loginCustomerID
	}

	return md, nil
}

func (c *credentials) RequireTransportSecurity() bool {
	return true
}
