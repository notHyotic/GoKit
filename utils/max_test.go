package utils

import "testing"

func TestMax(t *testing.T) {
    if Max(10, 20) != 20 {
        t.Errorf("Expected 20, got %d", Max(10, 20))
    }
}
