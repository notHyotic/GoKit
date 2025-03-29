package utils

import "cmp"

// Max returns the larger of two comparable values.
func Max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}