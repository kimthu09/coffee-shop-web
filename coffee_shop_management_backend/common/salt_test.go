package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenSalt(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
	}{
		{
			name: "Test case 1",
			args: args{
				length: 10,
			},
			wantLen: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Nil(t, tt.wantLen, len(GenSalt(tt.args.length)))
		})
	}
}

func Test_randSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randSequence(tt.args.n); got != tt.want {
				t.Errorf("randSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
