package goiac

type Stage struct {
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description,omitempty"`
	Provider    string `mapstructure:"provider"`
	// Options for provider.
	Options *interface{} `mapstructure:"options,omitempty"`
}
