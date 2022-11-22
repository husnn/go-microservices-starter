package env

import (
	"github.com/namsral/flag"
	"testing"
)

const (
	Dev     = "dev"
	Staging = "staging"
	Prod    = "prod"
)

const (
	prodBase = "https://boilerplate.pk"
	prodApi  = "https://api.boilerplate.pk"
)

var env = flag.String("env", Dev, "Environment")

func IsDev() bool {
	return *env == Dev
}

func IsStaging() bool {
	return *env == Staging
}

func IsProd() bool {
	return *env == Prod
}

func SetForTesting(t testing.TB, newEnv string) {
	oldEnv := *env
	*env = newEnv
	t.Cleanup(func() {
		*env = oldEnv
	})
}

func BaseURL() string {
	if IsDev() {
		return "http://127.0.0.1:9000"
	}

	if IsStaging() {
		return "https://staging.boilerplate.pk"
	}

	return prodBase
}
