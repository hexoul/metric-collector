package imon

import "net/http"

// GetEntities returns an entity list for metrics
func GetEntities() []string {
	return nil
}

// MakeRequestContext creates a HTTP request context to get information
func MakeRequestContext(metric string) (*http.Request, error) {
	return nil, nil
}
