package utilities

// 指定された範囲のsliceを返す
func Range[T int | uint](start, end T) []T {
	if start > end {
		return nil
	}
	s := make([]T, end-start)
	for i := range s {
		s[i] = start + T(i)
	}
	return s
}

// 指定した値がsliceに含まれているかを返す
func Contains[T byte | ~int | ~uint | string | bool](slice *[]T, value T) bool {
	for _, v := range *slice {
		if v == value {
			return true
		}
	}
	return false
}
