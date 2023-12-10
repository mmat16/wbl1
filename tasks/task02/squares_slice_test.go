package squeares_slice

import (
	"os"
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSquareSliceCh(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareSliceCh(tt.args.nums)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareSliceWG(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareSliceWG(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareSliceChOrdered(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareSliceChOrdered(tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquareSliceMu(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "first",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "empty slice",
			args: args{
				nums: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SquareSliceMu(tt.args.nums)
			slices.Sort(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquareSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

var old = os.Stdout

func BenchmarkSquareSliceChan(b *testing.B) {
	file, _ := os.Open("output.txt")
	os.Stdout = file
	for n := 0; n < b.N; n++ {
		SquareSliceCh(slice)
	}
	os.Stdout = old
	file.Close()
}

func BenchmarkSquareSliceWG(b *testing.B) {
	file, _ := os.Open("output.txt")
	os.Stdout = file
	for n := 0; n < b.N; n++ {
		SquareSliceWG(slice)
	}
	os.Stdout = old
	file.Close()
}

func BenchmarkSquareSliceChanOrdered(b *testing.B) {
	file, _ := os.Open("output.txt")
	os.Stdout = file
	for n := 0; n < b.N; n++ {
		SquareSliceChOrdered(slice)
	}
	os.Stdout = old
	file.Close()
}

func BenchmarkSquareSliceMu(b *testing.B) {
	file, _ := os.Open("output.txt")
	os.Stdout = file
	for n := 0; n < b.N; n++ {
		SquareSliceMu(slice)
	}
	os.Stdout = old
	file.Close()
}
