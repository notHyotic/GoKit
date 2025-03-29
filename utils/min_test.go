package utils

import "testing"

func TestMin(t *testing.T) {
    if Min(10, 20) != 10 {
        t.Errorf("Expected 10, got %d", Min(10, 20))
    }
}