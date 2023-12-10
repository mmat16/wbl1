package sum_slice_squares

import "testing"

func TestSumSliceSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				nums: []int{2, 3, 4},
			},
			want: 29,
		},
		{
			name: "second",
			args: args{
				nums: []int{15, 33, 29},
			},
			want: 2155,
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSliceSquares(tt.args.nums); got != tt.want {
				t.Errorf("SumSliceSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumSliceSquaresMu(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				nums: []int{2, 3, 4},
			},
			want: 29,
		},
		{
			name: "second",
			args: args{
				nums: []int{15, 33, 29},
			},
			want: 2155,
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumSliceSquaresMu(tt.args.nums); got != tt.want {
				t.Errorf("SumSliceSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

func BenchmarkSumSliceSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSliceSquares(nums)
	}
}

func BenchmarkSumSliceSquaresMu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSliceSquaresMu(nums)
	}
}

func BenchmarkSumDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumDefer(nums)
	}
}
