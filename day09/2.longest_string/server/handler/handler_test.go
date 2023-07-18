package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StringHandlerMok struct{}

func (s StringHandlerMok) ServeString(str string) string {
	if str == "a" {
		return "a"
	}
	if str == "b" {
		return "b"
	}

	return "c"
}

func TestHandler_HandleString(t *testing.T) {
	type fields struct {
		sh StringHandler
	}
	type args struct {
		h Handler
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK_1",
			args: args{
				Newhandler(StringHandlerMok{}),
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/api/substring", bytes.NewReader([]byte("a")))
			res := httptest.NewRecorder()
			tt.args.h.InitRouter().ServeHTTP(res, req)

			if got := res.Result().StatusCode; got != tt.want {
				t.Errorf("handler.HandleString() = %v, want %v", got, tt.want)

			}
		})
	}
}
