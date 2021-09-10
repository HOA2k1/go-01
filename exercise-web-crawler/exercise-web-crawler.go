package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Tìm và trả về nội dung của URL và một đoạn của URL tìm thấy trên trang đó
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]bool
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}

func (u SafeCounter) IsCrawled(url string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.v[url]; ok == false {
		u.v[url] = true
		return false
	}
	return true
}

// Thu thập thông tin sử dụng trình tìm nạp để thu thập thông tin một cách đệ quy các trang bắt đầu bằng url, có độ sâu tối đa
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	// Kiểm tra xem URL đã được thu thập thông tin chưa
	if c.IsCrawled(url) == true {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(2 * time.Second)
}

// fakeFetcher Trả về kết quả soạn trước.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/cmd/ii",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		"cmd",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/cmd/ii",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
