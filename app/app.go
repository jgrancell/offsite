package app

import (
	"strings"

	"github.com/jgrancell/offsite/configuration"
)

type App struct {
	Configuration *configuration.Configuration
	Version       string
	Flags         map[string]*Flag
	RunType       string
}

type Flag struct {
	Name  string
	Type  string
	Value interface{}
}

func (a *App) Load(args []string, version string) error {
	var err error
	a.Version = version

	err = a.ParseFlags(args)
	if err != nil {
		return err
	}

	a.Configuration = &configuration.Configuration{}
	if _, ok := a.Flags["mode"]; ok {
		err = a.Configuration.Load(a.Flags["mode"].Value.(string))
	} else {
		err = a.Configuration.Load("portable")
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *App) ParseFlags(args []string) error {
	var skipNext bool

	a.Flags = make(map[string]*Flag)
	for i := 0; i < len(args); i++ {
		if skipNext {
			skipNext = false
		} else {
			if strings.Index(args[i], "-") == 0 {
				// Checking to see if boolean flag or assignment flag
				flagName := strings.TrimLeft(args[i], "-")

				// Looking to see if boolean (-), or assignment (--)
				if strings.Index(args[i], "--") == 0 {
					a.Flags[flagName] = &Flag{
						Name:  flagName,
						Type:  "assignment",
						Value: args[i+1],
					}
					skipNext = true
				} else {
					a.Flags[flagName] = &Flag{
						Name:  flagName,
						Type:  "boolean",
						Value: true,
					}
				}
			}
		}
	}
	return nil
}
