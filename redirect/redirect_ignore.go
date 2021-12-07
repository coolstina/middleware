package redirect

// Ignore will skip the redirect detection without triggering redirection events.
type Ignore map[string]bool

// URIIsIgnore Checks whether the URI is set to ignore.
func (ignore Ignore) URIIsIgnore(uri string) bool {
	access.Lock()
	defer access.Unlock()

	for k, v := range ignore {
		if k == uri && v {
			return true
		}
	}
	return false
}
