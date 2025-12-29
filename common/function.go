package common

// TernaryAny 三元运算
func TernaryAny[T any](ok bool, okDo T, noOkDo T) T {
	if ok {
		return okDo
	} else {
		return noOkDo
	}
}
