package goiac

import (
	"bytes"
	"text/template"

	"github.com/spf13/viper"
)

func ReadAndTemplate(path string) (*viper.Viper, error) {
	var tpl bytes.Buffer

	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}
	if err := t.Execute(&tpl, G); err != nil {
		return nil, err
	}
	cfg := viper.New()
	cfg.SetConfigType("yaml")
	parsed := bytes.NewReader(tpl.Bytes())
	err = cfg.ReadConfig(parsed)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
