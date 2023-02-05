package pkg

import "testing"

func TestMessage(t *testing.T) {
	expected := "Service Image is running."

	if Message() != expected {
		t.Errorf(`expected to be "%s"`, expected)
	}
}
