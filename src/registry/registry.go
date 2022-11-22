package registry

import (
	"fmt"
	"github.com/namsral/flag"
	"os"
	"boilerplate/utils"
)

var ports = map[string]int{
	"gateway": 30000,
	"users":   30010,
	"guard":   30020,
	"auth":    30030,
}

func ServiceAddress(name string) string {
	if addr, ok := ports[name]; ok {
		return utils.MaybePrependLocalhost(
			fmt.Sprintf(":%d", addr))
	}
	return ""
}

func ClientAddress(name string) string {
	addr := flag.String(fmt.Sprintf(
		"%s_address", name), "", "service address")
	if *addr != "" {
		return *addr
	}

	if _, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST"); ok {
		// Running inside K8s
		return fmt.Sprintf("%s:%d", name, ports[name])
	}

	return ServiceAddress(name)
}
