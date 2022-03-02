package advanced

import "testing"

func TestRingbuffer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ringbuffer()
		})
	}
}
