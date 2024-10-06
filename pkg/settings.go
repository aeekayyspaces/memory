// Package pkg Shared library for the memory application.
package pkg

import (
	"github.com/aeekayyspaces/memory/pkg/base"
)

type Settings struct {
	base.Singleton
}

func (s *Settings) Init() *Settings {
	baseSettings := make(map[string]interface{})
	baseSettings["database_connection_string"] = "postgres://localhost:5432"

	return &Settings{
		Name: "MemorySettings",
		Data: baseSettings,
	}	
}
