package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Listen(port int) {
	var err error
	http.HandleFunc("GET /",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			log.Println("Received request:\n=======> Method:", r.Method, "\n=======> Host:", r.URL.String(), "\n=======> Client:", r.RemoteAddr)
			r.Header.Set("X-Forwarded-For", r.RemoteAddr)

			if isForbiddenHost(r.URL.Host) {
				// w.WriteHeader(http.StatusForbidden)
				log.Println("Website content not allowed.", r.URL.Host)
				http.Error(w, "Website content not allowed.", http.StatusForbidden)
				return
			}

			if containsBannedWord(r.URL.Host) {
				log.Println("Website content not allowed.", r.URL.Host)
				http.Error(w, "Website content not allowed.", http.StatusForbidden)
				return
			}

			res, err := Request(r.URL.String())
			if err != nil {
				http.Error(w, "Failed to fetch target URL", http.StatusBadGateway)
				return
			}
			if res.StatusCode == 200 {
				defer res.Body.Close()

				for key, values := range res.Header {
					for _, value := range values {
						w.Header().Add(key, value)
					}
				}
				w.WriteHeader(res.StatusCode)

				log.Println("Forwarding Response to:", r.RemoteAddr)
				io.Copy(w, res.Body)
			}
		})
	err = http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil)
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
