package strs

import (
	"reflect"
	"testing"
)

func TestDelEscape(t *testing.T) {
	type args struct {
		s      []byte
		escape []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{s: []byte("a\bc"), escape: []byte("\b")},
			want: []byte("c"),
		},
		{
			name: "test2",
			args: args{s: []byte("abc\bd\b"), escape: []byte("\b")},
			want: []byte("ab"),
		},
		{
			name: "test3",
			args: args{s: []byte("abca\bsde\b\b\bd\b"), escape: []byte("\b")},
			want: []byte("abc"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DelEscape(tt.args.s, tt.args.escape); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DelEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}
