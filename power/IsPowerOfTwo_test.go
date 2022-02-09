package power

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example1",
			args: args{n: 1},
			want: true,
		},
		{
			name: "Example2",
			args: args{n: 16},
			want: true,
		},
		{
			name: "Example3",
			args: args{n: 3},
			want: false,
		},
		{
			name: "Example4",
			args: args{n: 1024},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("IsPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
