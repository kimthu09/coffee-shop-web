package common

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenSalt(t *testing.T) {
	for _, length := range []int{-5, 0, 10, 20} {
		result := GenSalt(length)

		expectedLength := length
		if length < 0 {
			expectedLength = 50
		}

		assert.Equal(
			t,
			expectedLength,
			len(result),
			"Expected length = %d, got %d", len(result), expectedLength)

		lettersStr := string(letters)

		for _, char := range result {
			isExist := strings.IndexRune(lettersStr, char) != -1
			assert.Equal(
				t,
				true,
				isExist,
				"Invalid character in result: %c", char)
		}
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
			assert.Equalf(t, tt.want, randSequence(tt.args.n), "randSequence(%v)", tt.args.n)
		})
	}
}
