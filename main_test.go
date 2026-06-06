package main

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd")
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsd")
	os.Setenv("REDIS_DSN", "localhost:6379")
	// init() runs automatically and we're just making sure it passes
}
