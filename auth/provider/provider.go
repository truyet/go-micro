// Package provider is an external auth provider e.g oauth
package provider

// Provider is an auth provider
type Provider interface {
	// String returns the name of the provider
	String() string
	// Options returns the options of a provider
	Options() Options
}
