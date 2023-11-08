package authlib

import "fmt"

func GetUrlPathAuthorityCode(urlPath, method string) string {
	return fmt.Sprintf("urlpaths:%s#%s", urlPath, method)
}
