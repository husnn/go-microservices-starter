package utils

import (
	"fmt"
	"github.com/namsral/flag"
)

var BaseApiUrl = flag.String("base_url", "localhost:30000", "internal api base url")
var ClientUrl = flag.String("client_url", "http://localhost:3000", "client url")

func Url(version, endpoint string) string {
	if version != "" {
		return fmt.Sprintf("%s/%s/%s", *BaseApiUrl, version, endpoint)
	}
	return fmt.Sprintf("%s/%s", *BaseApiUrl, endpoint)
}
