package config

import (
	"fmt"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	m, err := EnvRead("./testdata/.env.local")
	if err != nil {
		t.Fatal(t)
	}

	for k, v := range m {
		fmt.Println(fmt.Sprintf("k:%s v:%s", k, v))
	}
}

func TestReadOSEnv(t *testing.T) {
	_ = os.Setenv("FOO_ENV", "foo")
	m, err := EnvRead("./testdata/.env.local")
	if err != nil {
		t.Fatal(t)
	}

	if m["foo_env"] != "foo" {
		t.Fatal("foo_env must get foo but get:" + m["foo_env"])
	}
}
