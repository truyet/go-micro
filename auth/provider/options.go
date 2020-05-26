package provider

import "golang.org/x/oauth2"

// Option returns a function which sets an option
type Option func(*Options)

// Options a provider can have
type Options struct {
	// Type of provider, e.g. oauth
	Type string `json:"type"`
	// ClientID is the application's ID.
	ClientID string `json:"client_id"`
	// ClientSecret is the application's secret.
	ClientSecret string `json:"client_secret"`
	// AuthURL for the provider
	AuthURL string `json:"auth_url"`
	// TokenURL for the provider
	TokenURL string `json:"token_url"`
	// AuthStyle to use on token request
	AuthStyle oauth2.AuthStyle
	// Redirect url incase of UI
	Redirect string `json:"redirect"`
	// Scope of the oauth request
	Scope string `json:"scope"`
}

// Credentials is an option which sets the client id and secret
func Credentials(id, secret string) Option {
	return func(o *Options) {
		o.ClientID = id
		o.ClientSecret = secret
	}
}

// AuthURL sets the auth url option
func AuthURL(a string) Option {
	return func(o *Options) {
		o.AuthURL = a
	}
}

// TokenURL sets the token url option
func TokenURL(a string) Option {
	return func(o *Options) {
		o.TokenURL = a
	}
}

// AuthStyle sets the oauth style option
func AuthStyle(a oauth2.AuthStyle) Option {
	return func(o *Options) {
		o.AuthStyle = a
	}
}

// Redirect sets the Redirect option
func Redirect(r string) Option {
	return func(o *Options) {
		o.Redirect = r
	}
}

// Scope sets the oauth scope
func Scope(s string) Option {
	return func(o *Options) {
		o.Scope = s
	}
}
