package stringService

import "testing"

func Test_stringService_ServeString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		s    stringService
		args args
		want string
	}{
		{
			name: "OK_1",
			s: NewStringService(),
			args: args{"abc"},
			want: "abc",
		},
		{
			name: "OK_2",
			s: NewStringService(),
			args: args{"aaaaa"},
			want: "a",
		},
		{
			name: "OK_3",
			s: NewStringService(),
			args: args{"aabcccdefrt"},
			want: "cdefrt",
		},
		{
			name: "OK_4",
			s: NewStringService(),
			args: args{"a"},
			want: "a",
		},
		{
			name: "OK_5",
			s: NewStringService(),
			args: args{""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stringService{}
			if got := s.ServeString(tt.args.str); got != tt.want {
				t.Errorf("stringService.ServeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
