package pkg

import (
	"testing"
)

// TestSettingsInit tests the Init() method of the Settings struct.
func TestSettingsInit(t *testing.T) {
	settings := &Settings{}
	initializedSettings := settings.Init()

	// Check if the initialized object is not nil
	if initializedSettings == nil {
		t.Errorf("Expected initialized settings but got nil")
	}

	// Check if the settings name is correctly set
	if initializedSettings.Name != "MemorySettings" {
		t.Errorf("Expected Name 'MemorySettings' but got %s", initializedSettings.Name)
	}

	// Check if the settings data map contains the expected database connection string
	expectedConnectionString := "postgres://localhost:5432"
	if connectionString, ok := initializedSettings.Data["database_connection_string"]; !ok {
		t.Errorf("Expected 'database_connection_string' to be present in the data map")
	} else if connectionString != expectedConnectionString {
		t.Errorf("Expected 'database_connection_string' to be %s but got %s", expectedConnectionString, connectionString)
	}
}

// TestSingletonInstanceConsistency tests if the settings instance data is consistent after initialization.
func TestSingletonInstanceConsistency(t *testing.T) {
	settings := &Settings{}
	initializedSettings1 := settings.Init()
	initializedSettings2 := settings.Init()

	// Ensure that both instances point to the same underlying data (singleton behavior)
	if initializedSettings1 != initializedSettings2 {
		t.Errorf("Expected both settings instances to be the same (singleton pattern) but they are different")
	}
}

