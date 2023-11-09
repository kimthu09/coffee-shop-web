package common

import "testing"

//func TestValidateDateString(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateDateString(tt.args.s); got != tt.want {
//				t.Errorf("ValidateDateString() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateEmail(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateEmail(tt.args.s); got != tt.want {
//				t.Errorf("ValidateEmail() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestValidateEmptyString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateEmptyString(tt.args.s); got != tt.want {
				t.Errorf("ValidateEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestValidateId(t *testing.T) {
//	type args struct {
//		id *string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateId(tt.args.id); got != tt.want {
//				t.Errorf("ValidateId() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateNegativeNumber(t *testing.T) {
//	type args struct {
//		number interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateNegativeNumber(tt.args.number); got != tt.want {
//				t.Errorf("ValidateNegativeNumber() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateNotNegativeNumber(t *testing.T) {
//	type args struct {
//		number interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateNotNegativeNumber(tt.args.number); got != tt.want {
//				t.Errorf("ValidateNotNegativeNumber() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateNotNilId(t *testing.T) {
//	type args struct {
//		id *string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateNotNilId(tt.args.id); got != tt.want {
//				t.Errorf("ValidateNotNilId() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateNotPositiveNumber(t *testing.T) {
//	type args struct {
//		number interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateNotPositiveNumber(tt.args.number); got != tt.want {
//				t.Errorf("ValidateNotPositiveNumber() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidatePassword(t *testing.T) {
//	type args struct {
//		pass *string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidatePassword(tt.args.pass); got != tt.want {
//				t.Errorf("ValidatePassword() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidatePhone(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidatePhone(tt.args.s); got != tt.want {
//				t.Errorf("ValidatePhone() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidatePositiveNumber(t *testing.T) {
//	type args struct {
//		number interface{}
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidatePositiveNumber(tt.args.number); got != tt.want {
//				t.Errorf("ValidatePositiveNumber() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestValidateUrl(t *testing.T) {
//	type args struct {
//		s string
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := ValidateUrl(tt.args.s); got != tt.want {
//				t.Errorf("ValidateUrl() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
