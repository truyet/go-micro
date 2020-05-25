package provider

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
	// Endpoint for the provider
	Endpoint string `json:"endpoint"`
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

// Endpoint sets the endpoint option
func Endpoint(e string) Option {
	return func(o *Options) {
		o.Endpoint = e
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
