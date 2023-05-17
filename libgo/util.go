package util

import "os"

// Require that a environment variable exist and be non-empty.
// Returns the value associated with the variable.
func MustOsGetEnv(key string) string {
	res := os.Getenv(key)
	if res == "" {
		panic("env var " + key + " must not be empty")
	}
	return res
}
