package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	var token string
	args := os.Args
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		url, _ := url.Parse("http://localhost:8080/oauth")

		q := url.Query()
		q.Add("client_id", args[1])
		q.Add("redirect_uri", "http://localhost:8081/goahead")
		q.Add("scope", "all_token")
		url.RawQuery = q.Encode()

		http.Redirect(w, r, url.String(), http.StatusFound)
	})

	http.HandleFunc("/goahead", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		code := r.URL.Query().Get("code")

		url, _ := url.Parse("http://localhost:8080/authorize")

		q := url.Query()
		q.Add("client_id", args[1])
		q.Add("client_secret", args[2])
		q.Add("code", code)

		url.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, url.String(), nil)
		if err != nil {
			log.Println(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		} else if resp != nil {
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			json.Unmarshal(b, &token)
		}
	})

	http.HandleFunc("/username", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/info", nil)
		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Add("Authorization", token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		} else if resp != nil {
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			w.Write(b)
		}
	})

	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/secret", nil)
		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Add("Authorization", token)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
		} else if resp != nil {
			defer resp.Body.Close()
			b, _ := io.ReadAll(resp.Body)
			w.Write(b)
		}
	})

	http.ListenAndServe(":8081", nil)
}
