package goiac

import (
	"fmt"
)

// Working with modules.
type Module struct {
	Name        string   `mapstructure:"name"`
	Description string   `mapstructure:"description,omitempty"`
	Version     int      `mapstructure:"version"`
	Author      string   `mapstructure:"author,omitempty"`
	License     string   `mapstructure:"license,omitempty"`
	ScriptFiles []string `mapstructure:"scripts"`
	S           []*Script
	Project     *Project
	Idx         int
}

// Read module from directory.
func ReadModule(path string, project *Project, idx int) (*Module, error) {
	initPath := path + "/init.yaml"
	cfg, err := ReadAndTemplate(initPath)
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't read module file.")
		return nil, err
	}

	module := Module{}
	err = cfg.Unmarshal(&module)
	if err != nil {
		Log.Error().Err(err).Str("path", path).Msg("Can't parse module file.")
		return nil, err
	}
	if project != nil {
		module.Project = project
	} else {
		p := Project{
			Name: "w/o project",
		}
		module.Project = &p
	}
	// Name is mandatory field.
	if module.Name == "" {
		Log.Error().Str("path", path).Msg("Name is mandatory filed for module")
		return nil, fmt.Errorf("Name is mandatory filed for module '%v'", path)
	}

	// Stages is mandatory field.
	if module.ScriptFiles == nil {
		Log.Error().Str("path", path).Msg("Scripts is mandatory filed for module")
		return nil, fmt.Errorf("Sripts is mandatory filed for module '%v'", path)
	}

	for i, sn := range module.ScriptFiles {
		if s, err := ReadScript(path+"/"+sn, &module, i); err != nil {
			return nil, err
		} else {
			module.S = append(module.S, s)
		}
	}

	return &module, nil
}
