package common

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewSuccessResponse(t *testing.T) {
	type args struct {
		data   interface{}
		paging interface{}
		filter interface{}
	}

	data := mock.Anything
	paging := Paging{Page: 1, Limit: 10, Total: 11}
	filter := "filter"

	tests := []struct {
		name string
		args args
		want *SuccessRes
	}{
		{
			name: "Get success response from data, paging and filter successfully",
			args: args{
				data:   &data,
				paging: &paging,
				filter: &filter,
			},
			want: &SuccessRes{
				Data:   &data,
				Paging: &paging,
				Filter: &filter,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSuccessResponse(tt.args.data, tt.args.paging, tt.args.filter)

			assert.Equal(
				t,
				tt.want,
				got,
				"NewSuccessResponse() = %v, want %v", got, tt.want)
		})
	}
}

func TestSimpleSuccessResponse(t *testing.T) {
	type args struct {
		data interface{}
	}

	data := mock.Anything

	tests := []struct {
		name string
		args args
		want *SuccessRes
	}{
		{
			name: "Get success response from data successfully",
			args: args{
				data: &data,
			},
			want: &SuccessRes{
				Data:   &data,
				Paging: nil,
				Filter: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleSuccessResponse(tt.args.data)

			assert.Equal(
				t,
				tt.want,
				got,
				"SimpleSuccessResponse() = %v, want %v", got, tt.want)
		})
	}
}
