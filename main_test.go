package main

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HelloWorld()
		})
	}
}
