package fizzbuzz

import "testing"

func TestReplace(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "divisible by 3 and 5",
			args: args{
				num: 30,
			},
			want: "FizzBuzz",
		},
		{
			name: "not divisible by 3 and 5",
			args: args{
				num: 26,
			},
			want: "26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.num); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}
