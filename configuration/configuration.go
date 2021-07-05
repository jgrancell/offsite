package configuration

import (
	"io/ioutil"
	"os"

	toml "github.com/pelletier/go-toml/v2"
)

type Configuration struct {
	ConfigPath string
	Peers      string
	Port       int
	Server     string
	Targets    []*Target
	Vaults     []*Vault
}

type Vault struct {
	Path               string
	AllocatedSize      int64
	UsedSize           int64
	AvailableSize      int64
	RequiresEncryption bool
}

type Target struct {
	Name            string
	Path            string
	Objects         []*Object
	TotaSize        int64
	ReplicaStatus   bool
	ReplicaCount    int64
	Encrypted       bool
	DesiredReplicas int64
}

type Object struct {
	Path   string
	Name   string
	ShaSum string
	Data   []byte
}

func (c *Configuration) Load(mode string) error {
	err := c.FindConfigFile()
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(c.ConfigPath)
			if err != nil {
				return err
			}
			_ = file.Close()
		} else {
			return err
		}
	}

	err = c.FetchConfig()
	return err
}

func (c *Configuration) FetchConfig() error {
	contents, err := ioutil.ReadFile(c.ConfigPath)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(contents, c); err != nil {
		return err
	}
	return nil
}

func (c *Configuration) Save() error {
	s, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.ConfigPath, s, 0600)
	return err
}
