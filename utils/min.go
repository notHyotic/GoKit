package utils

import "cmp"

// Min returns the smaller of two comparable values.
func Min[T cmp.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}