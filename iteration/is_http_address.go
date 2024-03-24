package iteration

import "strings"

func IsHttpAddress(httpAddress string) bool {
	if strings.HasPrefix(httpAddress, "www.") && strings.HasSuffix(httpAddress, ".com") {
		return true
	} else {
		return false
	}

}
