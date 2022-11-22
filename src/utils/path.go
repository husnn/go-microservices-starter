package utils

import (
	"boilerplate/env"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProjectPath() string {
	return os.Getenv("PROJECT_PATH")
}

func JoinProjectPath(paths ...string) string {
	return filepath.Join(append([]string{ProjectPath()}, paths...)...)
}

func MaybePrependLocalhost(addr string) string {
	s := strings.Split(addr, ":")
	// If no port, but host is provided.
	if len(s) <= 1 && len(s[0]) > 0 {
		return addr
	}

	// In development, prepend host if missing.
	if len(s[0]) < 1 && env.IsDev() {
		host := "127.0.0.1"

		if _, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST"); ok {
			host = "0.0.0.0"
		}

		return fmt.Sprintf("%s%s", host, addr)
	}

	return addr
}
