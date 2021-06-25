package middleware

//See docs for context.WithValue
type contextKey int

const (
	oidcStateKey contextKey = iota
	launchTokenKey
	resourcesKey
)
