package utils

func ToPointer[T any](v T) *T {
	return &v
}

func CheckBoolPointer(v *bool) bool {
	return v != nil && *v
}
