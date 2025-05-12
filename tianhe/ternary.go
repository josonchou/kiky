package tianhe

// Ternary 是一个泛型三元表达式工具函数，根据条件返回不同的值
// T 和 U 分别是真值和假值的类型，返回类型根据条件选择
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
