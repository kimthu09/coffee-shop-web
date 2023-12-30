package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	type args struct {
		s string
	}

	validEmail := "a@gmail.com"
	invalidEmail := "invalid"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate email successfully",
			args: args{
				s: validEmail,
			},
			want: true,
		},
		{
			name: "Validate email failed",
			args: args{
				s: invalidEmail,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateEmail(tt.args.s)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateEmail() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateEmptyString(t *testing.T) {
	type args struct {
		s string
	}

	validString := ""
	invalidString := "a"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate empty string successfully",
			args: args{
				s: validString,
			},
			want: true,
		},
		{
			name: "Validate empty string failed",
			args: args{
				s: invalidString,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateEmptyString(tt.args.s)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateEmptyString() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateId(t *testing.T) {
	type args struct {
		id *string
	}
	validId := "012345678901"
	emptyId := ""
	longId := "0123456789012"
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate id successfully",
			args: args{
				id: &validId,
			},
			want: true,
		},
		{
			name: "Validate id successfully because id is nil",
			args: args{
				id: nil,
			},
			want: true,
		},
		{
			name: "Validate id successfully because id has length equal 0",
			args: args{
				id: &emptyId,
			},
			want: true,
		},
		{
			name: "Validate id failed because id is too long",
			args: args{
				id: &longId,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateId(tt.args.id)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateId() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateNegativeNumberFloat(t *testing.T) {
	type args struct {
		number float32
	}

	validFloat := float32(-0.000000000001)
	invalidFloat := float32(0)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate negative number float successfully",
			args: args{
				number: validFloat,
			},
			want: true,
		},
		{
			name: "Validate negative number float failed",
			args: args{
				number: invalidFloat,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNegativeNumberFloat(tt.args.number)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateNegativeNumberFloat() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateNegativeNumberInt(t *testing.T) {
	type args struct {
		number int
	}
	validInt := -1
	invalidInt := 0

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate negative number int successfully",
			args: args{
				number: validInt,
			},
			want: true,
		},
		{
			name: "Validate negative number int failed",
			args: args{
				number: invalidInt,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNegativeNumberInt(tt.args.number)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateNegativeNumberInt() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateNotNilId(t *testing.T) {
	type args struct {
		id *string
	}
	validId := "012345678901"
	emptyId := ""
	longId := "0123456789012"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate id successfully",
			args: args{
				id: &validId,
			},
			want: true,
		},
		{
			name: "Validate id failed because id is nil",
			args: args{
				id: nil,
			},
			want: false,
		},
		{
			name: "Validate id failed because id has length equal 0",
			args: args{
				id: &emptyId,
			},
			want: false,
		},
		{
			name: "Validate id failed because id is too long",
			args: args{
				id: &longId,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNotNilId(tt.args.id)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateNotNilId() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidateNotPositiveNumberInt(t *testing.T) {
	type args struct {
		number int
	}
	validInt := 0
	invalidInt := 1

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate not positive number int successfully",
			args: args{
				number: validInt,
			},
			want: true,
		},
		{
			name: "Validate not positive number int failed",
			args: args{
				number: invalidInt,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNotPositiveNumberInt(tt.args.number)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidateNotPositiveNumberInt() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidatePassword(t *testing.T) {
	type args struct {
		pass *string
	}

	validPass := "123456"
	invalidPass := "12345"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate password successfully",
			args: args{
				pass: &validPass,
			},
			want: true,
		},
		{
			name: "Validate password failed because password's length is not enough",
			args: args{
				pass: &invalidPass,
			},
			want: false,
		},
		{
			name: "Validate password failed because password is nil",
			args: args{
				pass: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePassword(tt.args.pass)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidatePassword() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidatePhone(t *testing.T) {
	type args struct {
		s string
	}

	validPhone := "0123456789"
	validPhone1 := "01234567890"
	invalidPhone := "012"
	invalidPhone1 := "012345a678"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate phone with 10 numbers successfully",
			args: args{
				s: validPhone,
			},
			want: true,
		},
		{
			name: "Validate phone with 11 numbers successfully",
			args: args{
				s: validPhone1,
			},
			want: true,
		},
		{
			name: "Validate phone with non-11-numbers and non-10-numbers failed",
			args: args{
				s: invalidPhone,
			},
			want: false,
		},
		{
			name: "Validate phone with phone has character failed",
			args: args{
				s: invalidPhone1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePhone(tt.args.s)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidatePhone() = %v, want %v", got, tt.want)
		})
	}
}

func TestValidatePositiveNumberInt(t *testing.T) {
	type args struct {
		number int
	}

	validInt := 1
	invalidInt := 0

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Validate positive number int successfully",
			args: args{
				number: validInt,
			},
			want: true,
		},
		{
			name: "Validate positive number int failed",
			args: args{
				number: invalidInt,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidatePositiveNumberInt(tt.args.number)

			assert.Equal(
				t,
				tt.want,
				got,
				"ValidatePositiveNumberInt() = %v, want %v", got, tt.want)
		})
	}
}
