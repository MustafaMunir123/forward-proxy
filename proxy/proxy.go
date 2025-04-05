package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Listen(port int) {
	http.HandleFunc("GET /",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			log.Println("Received request:\n=======> Method:", r.Method, "\n=======> URL:", r.URL.String())
			res, err := Request(r.URL.String())
			if err != nil {
				http.Error(w, "Failed to fetch target URL", http.StatusBadGateway)
				return
			}
			if res.StatusCode == 200 {
				defer res.Body.Close()
				w.WriteHeader(res.StatusCode)

				for key, values := range res.Header {
					for _, value := range values {
						w.Header().Add(key, value)
					}
				}
				log.Println("Forwarding Response to:", r.RemoteAddr)
				io.Copy(w, res.Body)
			}
		})
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
	if err != nil {
		log.Fatal("Error Occured", err)
	}
}

func Request(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Invalid Request: %s \n=======> ERROR: %v", url, err)
		return nil, err
	}
	return res, nil
}
