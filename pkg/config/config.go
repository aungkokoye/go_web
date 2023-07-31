package config

import "text/template"

// hold application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
