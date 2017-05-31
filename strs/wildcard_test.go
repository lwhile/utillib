package strs

import "testing"

func TestWildcardMatch(t *testing.T) {
	type args struct {
		str     string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{str: "abc", pattern: "*"},
			want: true,
		},
		{
			name: "test2",
			args: args{str: "abc", pattern: "*?"},
			want: true,
		},
		{
			name: "test3",
			args: args{str: "abc", pattern: "*d"},
			want: false,
		},
		{
			name: "test4",
			args: args{str: "abc", pattern: "?"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WildcardMatch(tt.args.str, tt.args.pattern); got != tt.want {
				t.Errorf("WildcardMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
