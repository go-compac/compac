package compac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNlFromPtr(t *testing.T) {
	zeroValue := 0
	nonZeroValue := 1

	ptrToZeroValue := &zeroValue
	ptrToNonZeroValue := &nonZeroValue

	type testCase struct {
		name string
		args *int
		want Nl[int]
	}
	tests := []testCase{
		{
			name: "nil ptr",
			args: nil,
			want: Nl[int]{
				Data:  0,
				Valid: false,
			},
		},
		{
			name: "ptr to zero value",
			args: ptrToZeroValue,
			want: Nl[int]{
				Data:  0,
				Valid: true,
			},
		},
		{
			name: "ptr to non zero value",
			args: ptrToNonZeroValue,
			want: Nl[int]{
				Data:  1,
				Valid: true,
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NlFromPtr(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNlFromValue(t *testing.T) {
	type testCase struct {
		name string
		args int
		want Nl[int]
	}
	tests := []testCase{
		{
			name: "zero value",
			args: 0,
			want: Nl[int]{
				Data:  0,
				Valid: true,
			},
		},
		{
			name: "non zero value",
			args: 1,
			want: Nl[int]{
				Data:  1,
				Valid: true,
			},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NlFromValue(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
