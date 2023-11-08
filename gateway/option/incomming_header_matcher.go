package gateway

const (
	HeaderAccessToken = "Access-Token"
)

func IncommingHeaderMatcher(key string) (string, bool) {
	switch key {
	case HeaderAccessToken:
		return HeaderAccessToken, true
	default:
		return key, false
	}
}
