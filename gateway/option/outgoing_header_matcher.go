package gateway

import (
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func OutgoingHeaderMatcher(key string) (string, bool) {
	if key == "set-cookie" {
		return key, true
	} else if key == HeaderAccessToken {
		return HeaderAccessToken, true
	}
	return fmt.Sprintf("%s%s", runtime.MetadataHeaderPrefix, key), false
}
