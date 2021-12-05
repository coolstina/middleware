package redirect

// TriggerEvent trigger redirection event.
type TriggerEvent struct {
	StatusCode            int    // Http status code.
	RedirectURI           string // redirect uri.
	RedirectHeaderMessage string // Redirect header message,
}
