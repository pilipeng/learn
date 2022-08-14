package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				prices: []int{8, 9, 2, 5, 4, 7, 1},
			},
			want: 5,
		}, {
			name: "2",
			args: args{
				prices: []int{2, 4, 1},
			},
			want: 2,
		}, {
			name: "3",
			args: args{
				prices: []int{3, 2, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
