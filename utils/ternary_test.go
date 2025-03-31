package utils

import (
    "testing"
)

func TestTernary(t *testing.T) {
    tests := []struct {
        condition bool
        trueValue any
        falseValue any
        expected any
    }{
        {true, 10, 20, 10},
        {false, 10, 20, 20},
        {true, "hello", "world", "hello"},
        {false, "hello", "world", "world"},
    }

    for _, test := range tests {
        t.Run("Test", func(t *testing.T) {
            result := Ternary(test.condition, test.trueValue, test.falseValue)
            if result != test.expected {
                t.Errorf("For condition %v, expected %v but got %v", test.condition, test.expected, result)
            }
        })
    }
}