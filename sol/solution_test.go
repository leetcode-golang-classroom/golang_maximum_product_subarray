package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{2, 3, -2, 4}
	for idx := 0; idx < b.N; idx++ {
		maxProduct(nums)
	}
}
func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [2,3,-2,4]",
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "nums = [-2,0,-1]",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
