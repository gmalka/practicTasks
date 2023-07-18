package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("Incorrect Method", http.MethodPost, r.Method)
		}
		if r.URL.String() != "/test" {
			t.Errorf("Ожидался путь %s, получен %s", "/test", r.URL.String())
		}
		// Отправляем ответ
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "abc")
	}))
	defer ts.Close()

	urls := strings.Split(ts.URL, "/")

	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "regular request",
			args: args{
				args: []string{"", "abc", urls[len(urls) - 1] + "/test"},
			},
			want: "abc",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				args: []string{"abc", urls[len(urls) - 1] + "/test"},
			},
			want: "",
			wantErr: true,
		},
		{
			name: "http error",
			args: args{
				args: []string{"abc", urls[len(urls) - 1] + "/test2"},
			},
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := MakeRequest(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
