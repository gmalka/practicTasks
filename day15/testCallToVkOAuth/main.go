package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GoTo localhost:8080/vk to get UserInfo

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		req, _ := http.NewRequest("GET", "https://oauth.vk.com/access_token", nil)

		q := req.URL.Query()
		q.Add("client_id", "51712905")
		q.Add("client_secret", "wziqnKD7XwWgqnkEu9zy")
		q.Add("redirect_uri", "http://localhost:8080")
		q.Add("code", code)

		req.URL.RawQuery = q.Encode()

		resp, _ := http.DefaultClient.Do(req)
		defer resp.Body.Close()
		s := struct {
			Access_token string `json:"access_token"`
		}{}
		b, _ := io.ReadAll(resp.Body)
		json.Unmarshal(b, &s)

		req, _ = http.NewRequest("GET", "https://api.vk.com/method/users.get", nil)
		req.Header.Add("Authorization", "Bearer "+s.Access_token)
		q = req.URL.Query()
		q.Add("v", "5.131")
		req.URL.RawQuery = q.Encode()

		resp, _ = http.DefaultClient.Do(req)
		b, _ = io.ReadAll(resp.Body)
		w.Write(b)

		return
	})

	http.HandleFunc("/vk", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodGet {
			fmt.Println("No GET")
			return
		}

		url, _ := url.Parse("https://oauth.vk.com/authorize")

		q := url.Query()
		q.Add("client_id", "51712905")
		q.Add("redirect_uri", "http://localhost:8080")
		q.Add("display", "page")
		q.Add("response_type", "code")
		url.RawQuery = q.Encode()

		http.Redirect(w, r, url.String(), http.StatusFound)
	})

	http.ListenAndServe(":8080", nil)
}