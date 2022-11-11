package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

var incomingURLs = []string{
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
}

func httpGetBoby(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBoby)
	var n sync.WaitGroup
	for _, url := range incomingURLs {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
