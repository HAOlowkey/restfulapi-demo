package client

import "context"

func NewDefaultClientConfig() *ClientConfig {
	return &ClientConfig{
		Addr:           "127.0.0.1:18050",
		Authentication: &Authentication{},
	}
}

type ClientConfig struct {
	Addr string
	*Authentication
}

const (
	ClientHeaderKey = "client-id"
	ClientSecretKey = "client-secret"
)

// Authentication todo
type Authentication struct {
	clientID     string
	clientSecret string
}

func (a *Authentication) SetClientCredentials(clientID, clientSecret string) {
	a.clientID = clientID
	a.clientSecret = clientSecret
}

// GetRequestMetadata todo
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{
		ClientHeaderKey: a.clientID,
		ClientSecretKey: a.clientSecret,
	}, nil
}

// RequireTransportSecurity todo
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
