package configuration

import (
	"os"
)

func (c *Configuration) FindConfigFile() error {
	// We support system-wide, single-user, and portable installs
	// This function looks for the correct configuration file
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Checking to see if there's an offsite.conf file in the current working directory
	portable_path, err := c.CheckConfigExists(pwd)
	if err == nil {
		// Returning the local configuration file and no error
		c.ConfigPath = portable_path
		return nil
	}

	// TODO: Enable these checks
	/*
		// Checking to see if we have XDG specific configuration environment variables
		if local_path, err := os.UserConfigDir(); err == nil {
			if path, err := c.CheckConfigExists(local_path); err == nil {
				c.ConfigPath = path
				return nil
			}
		}
	*/

	c.ConfigPath = portable_path
	return os.ErrNotExist
}

func (c *Configuration) CheckConfigExists(path string) (string, error) {
	fullPath := path + string(os.PathSeparator) + "offsite.conf"
	if _, err := os.Stat(fullPath); err != nil {
		return fullPath, err
	}
	return fullPath, nil
}
