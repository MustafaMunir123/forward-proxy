package proxy

import (
	"io"
	"log"
	"net/http"
	"time"
)

type CachedResponse struct {
	Header   http.Header
	Body     []byte
	Status   int
	ExpireAt time.Time
}

var cache = make(map[string]*CachedResponse)

func cacheNewResponse(url string, res http.Response) *CachedResponse {
	currentTime := time.Now()
	expireAt := currentTime.Add(5 * time.Minute)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error Occured while caching response: ", err)
	}

	cachedRes := &CachedResponse{
		Header:   res.Header,
		Body:     body,
		Status:   res.StatusCode,
		ExpireAt: expireAt,
	}
	cache[url] = cachedRes
	return cachedRes

}

func getCachedResponse(url string) *CachedResponse {
	currentTime := time.Now()

	if cachedResponse, ok := cache[url]; ok {
		if cachedResponse.ExpireAt.Before(currentTime) {
			log.Println("Cache Expires, fetching server")
			delete(cache, url)
		} else {
			log.Println("Serving from cache:", url)
			return cachedResponse
		}
	}
	return nil
}
