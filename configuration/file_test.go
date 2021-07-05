package configuration

import (
	"os"
	"testing"
)

func TestFindConfigFile(t *testing.T) {
	pwd, _ := os.Getwd()
	c := Configuration{}
	if c.ConfigPath != "" {
		t.Errorf("expected default config path to be empty, got %s", c.ConfigPath)
	}

	err := c.FindConfigFile()
	if err == nil {
		t.Errorf("expected error to equal os.ErrNotExist, got nil")
	}

	if !os.IsNotExist(err) {
		t.Errorf("expected os.ErrNotExist error, got %s", err.Error())
	}

	if c.ConfigPath == "" {
		t.Errorf("expected config path to equal ./offsite.conf, got blank")
	}

	// Moving into our test workdir to get a success
	os.Chdir("../testdata/configuration")
	c = Configuration{}
	if c.ConfigPath != "" {
		t.Errorf("expected default config path to be empty, got %s", c.ConfigPath)
	}

	err = c.FindConfigFile()
	if err != nil {
		t.Errorf("expected error to nil, got %s", err.Error())
	}

	if c.ConfigPath == "" {
		t.Errorf("expected config path to equal ./offsite.conf, got blank")
	}

	os.Chdir(pwd)
}

func TestCheckConfigExists(t *testing.T) {
	path := "../testdata/configuration"

	c := Configuration{}
	fullPath, err := c.CheckConfigExists(path)
	if err != nil {
		t.Errorf("expected no error when checking if config exists, got: %s", err.Error())
	}

	expectedPath := "../testdata/configuration/offsite.conf"
	if fullPath != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, fullPath)
	}
}
