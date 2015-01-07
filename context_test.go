package deliver

import (
	"testing"
)

func TestSetGetContext(t *testing.T) {
	context := NewContext()
	key := "key"

	if value := context.Get(key); value != nil {
		t.Errorf("Context value should not exist for key '%v'", key)
	}

	context.Set(key, "value")

	if value := context.Get(key).(string); value != "value" {
		t.Errorf("Context value mismatch: '%v'", value)
	}
}

func TestGetOkContext(t *testing.T) {
	context := NewContext()
	key := "key"

	if _, ok := context.GetOk(key); ok {
		t.Errorf("Context value should not exist for key '%v'", key)
	}

	context.Set(key, "value")

	if value, ok := context.GetOk(key); ok {
		if value != "value" {
			t.Errorf("Context value mismatch: '%v'", value)
		}
	} else {
		t.Errorf("Context should have key '%v'", key)
	}
}

func TestHasContext(t *testing.T) {
	context := NewContext()
	key := "key"

	if context.Has(key) {
		t.Errorf("Context should not have key '%v'", key)
	}

	context.Set(key, "value")

	if !context.Has(key) {
		t.Errorf("Context should have key '%v'", key)
	}
}