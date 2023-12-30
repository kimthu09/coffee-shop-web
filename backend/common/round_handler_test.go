package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomRound(t *testing.T) {
	type args struct {
		num *float32
	}

	num1 := float32(0.0006)
	numResult1 := float32(0.001)

	num2 := float32(0.006)
	numResult2 := float32(0.006)

	tests := []struct {
		name string
		args args
		want *float32
	}{
		{
			name: "Round 0.0006 successfully",
			args: args{
				num: &num1,
			},
			want: &numResult1,
		},
		{
			name: "Round 0.006 successfully",
			args: args{
				num: &num2,
			},
			want: &numResult2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CustomRound(tt.args.num)

			assert.Equal(
				t,
				tt.want,
				tt.args.num,
				"CustomRound() num = %v, want %v", tt.args.num, tt.want)
		})
	}
}

func TestRoundToInt(t *testing.T) {
	type args struct {
		num float32
	}

	num1 := float32(0.0006)
	numResult1 := 0

	num2 := float32(0.49)
	numResult2 := 0

	num3 := float32(0.5)
	numResult3 := 1

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Round to int 0.0006 successfully",
			args: args{
				num: num1,
			},
			want: numResult1,
		},
		{
			name: "Round to int 0.49 successfully",
			args: args{
				num: num2,
			},
			want: numResult2,
		},
		{
			name: "Round to int 0.5 successfully",
			args: args{
				num: num3,
			},
			want: numResult3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RoundToInt(tt.args.num)

			assert.Equal(
				t,
				tt.want,
				got,
				"CustomRound() num = %v, want %v", got, tt.want)
		})
	}
}
