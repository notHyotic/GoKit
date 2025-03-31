package utils

func Ternary(condition bool, trueValue, falseValue any) any {
    if condition {
        return trueValue
    }
    return falseValue
}