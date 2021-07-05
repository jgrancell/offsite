package configuration

import (
	"os"
	"testing"
)

func TestFetchConfig(t *testing.T) {
	pwd, _ := os.Getwd()

	os.Chdir("../testdata/configuration")
	c := Configuration{}
	c.Load("portable")
	c.Save()
	os.Chdir(pwd)
}
