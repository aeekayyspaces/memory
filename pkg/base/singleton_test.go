package base

import (
	"sync"
	"testing"
)

// Reset the singleton for testing purposes
func resetSingleton() {
	once = sync.Once{} // reset sync.Once for each test case
	instance = nil     // reset the instance for each test case
}

// TestSingletonInstance tests whether the same instance is returned on multiple calls
func TestSingletonInstance(t *testing.T) {
	resetSingleton() // Reset singleton before testing

	// First call to GetInstance
	s1 := &Singleton{}
	instance1 := s1.GetInstance()

	// Second call to GetInstance
	s2 := &Singleton{}
	instance2 := s2.GetInstance()

	if instance1 != instance2 {
		t.Errorf("Expected the same instance but got different instances")
	}

	if instance1 == nil || instance2 == nil {
		t.Errorf("Singleton instance should not be nil")
	}
}

// TestSingletonDataAccess tests whether data within the singleton is consistent across instances
func TestSingletonDataAccess(t *testing.T) {
	resetSingleton() // Reset singleton before testing

	// Access the instance and modify data
	singleton := &Singleton{}
	instance := singleton.GetInstance()

	// Set data in the singleton
	instance.Data = map[string]interface{}{
		"key1": "value1",
	}

	// Access data through another instance and verify it is the same data
	singleton2 := &Singleton{}
	data := singleton2.GetInstance().GetData()

	if data["key1"] != "value1" {
		t.Errorf("Expected 'key1' to be 'value1', but got %v", data["key1"])
	}
}

// TestSingletonThreadSafety tests the thread-safety of the singleton using goroutines
func TestSingletonThreadSafety(t *testing.T) {
	resetSingleton() // Reset singleton before testing

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(100)

	var instance1 *Singleton
	var instance2 *Singleton

	// Spawn multiple goroutines to access the singleton
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()

			singleton := &Singleton{}
			if instance1 == nil {
				instance1 = singleton.GetInstance()
			} else {
				instance2 = singleton.GetInstance()
				if instance1 != instance2 {
					t.Errorf("Expected the same instance but got different instances")
				}
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}


