package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaginationDBArgs(t *testing.T) {
	tests := []struct {
		name string
		page int
		size int
		want struct {
			Limit  int32
			Offset int32
		}
	}{
		{
			name: "Test Case 1: Page 1, Size 10",
			page: 1,
			size: 10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  10,
				Offset: 0,
			},
		},
		{
			name: "Test Case 2: Page 2, Size 10",
			page: 2,
			size: 10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  10,
				Offset: 10,
			},
		},
		{
			name: "Test Case 3: Page 3, Size 10",
			page: 3,
			size: 10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  10,
				Offset: 20,
			},
		},
		{
			name: "Test Case 4: Page 1, Size 20",
			page: 1,
			size: 20,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  20,
				Offset: 0,
			},
		},
		{
			name: "Test Case 5: Page 2, Size 20",
			page: 2,
			size: 20,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  20,
				Offset: 20,
			},
		},
		{
			name: "Test Case 6: Page 0, Size 10",
			page: 0,
			size: 10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  0,
				Offset: 0,
			},
		},
		{
			name: "Test Case 7: Page -1, Size 10",
			page: -1,
			size: 10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  0,
				Offset: 0,
			},
		},
		{
			name: "Test Case 8: Page 1, Size 0",
			page: 1,
			size: 0,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  0,
				Offset: 0,
			},
		},
		{
			name: "Test Case 9: Page 1, Size -10",
			page: 1,
			size: -10,
			want: struct {
				Limit  int32
				Offset int32
			}{
				Limit:  0,
				Offset: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createPaginationDBArgs(tt.page, tt.size)
			assert.Equal(t, tt.want, got)
		})
	}
}
