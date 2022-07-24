package goiac

import (
	"fmt"

	"github.com/spf13/viper"
)

type Host struct {
	Roles    map[string]string `mapstructure:"roles"`
	Platform string            `mapstructure:"platform,omitempty"`
	Arch     string            `mapstructure:"arch,omitempty"`
	Os       string            `mapstructure:"os,omitempty"`
}

type Project struct {
	Name       string              `mapstructure:"name"`
	Descrition string              `mapstructure:"description,omitempty"`
	Platform   string              `mapstructure:"platform,omitempty"`
	Arch       string              `mapstructure:"arch,omitempty"`
	Roles      map[string][]string `mapstructure:"roles,omitempty"`
	Hosts      map[string]Host     `mapstructure:"hosts,omitempty"`
	Modules    []*Module
}

// Read module from directory.
func ReadProject(path string, hostName string) (*Project, error) {
	var moduleList []string
	cfg := viper.New()
	initPath := path + "/init.yaml"
	cfg.SetConfigFile(initPath)
	cfg.SetConfigType("yaml")
	err := cfg.ReadInConfig()
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't read project.")
		return nil, err
	}
	project := Project{}
	err = cfg.Unmarshal(&project)
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't parse project file.")
		return nil, err
	}

	// Name is mandatory field.
	if project.Name == "" {
		Log.Error().Str("path", path).Msg("Name is mandatory filed for project")
		return nil, fmt.Errorf("Name is mandatory filed for project '%v'", path)
	}

	// Roles is mandatory field.
	if project.Roles == nil {
		Log.Error().Str("path", path).Msg("Roles is mandatory filed for project")
		return nil, fmt.Errorf("Roles is mandatory filed for project '%v'", path)
	}

	// Build modules list from host and roles. Always use hostname from host (hostname -s) but not from -H.
	if project.Roles != nil {
		if m, ok := project.Roles["all"]; ok {
			moduleList = append(moduleList, m...)
		}
		if host, ok := project.Hosts[hostName]; ok {
			if host.Roles != nil {
				for _, r := range host.Roles {
					if m, ok := project.Roles[r]; ok {
						moduleList = append(moduleList, m...)
					} else {
						Log.Error().Str("host", hostName).Msgf("Role '%v' not found in roles list", r)
						return nil, fmt.Errorf("Role '%v' not found in roles list", r)
					}
				}
			}
		}
	}
	if len(moduleList) == 0 {
		Log.Error().Str("host", hostName).Msgf("There is no modules for host")
		return nil, fmt.Errorf("There is no modules for host")
	}

	for i, mn := range moduleList {
		if m, err := ReadModule(path+"/modules/"+mn, &project, i); err != nil {
			return nil, err
		} else {
			project.Modules = append(project.Modules, m)
		}
	}

	return &project, nil
}
