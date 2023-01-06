package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type reqUrl struct {
	w   *sync.WaitGroup
	m   *sync.Mutex
	res []*http.Response
	suc bool
}

func (r *reqUrl) setFails() {
	r.m.Lock()
	r.suc = false
	r.res = []*http.Response{}
	r.m.Unlock()
}

func (r *reqUrl) isFails() bool {
	r.m.Lock()
	defer r.m.Unlock()
	return r.suc
}

func (r *reqUrl) appendRes(res *http.Response) {
	r.m.Lock()
	r.res = append(r.res, res)
	r.m.Unlock()
}

func (r *reqUrl) requestUrl(url string, cl http.Client) {
	res, err := cl.Get(url)
	defer r.w.Done()
	if err != nil {
		r.setFails()
		return
	}
	if !r.isFails() {
		return
	}
	r.appendRes(res)
}

func callUrls(urls []string) *reqUrl {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int8 = 0
	// can't be one second Timeout
	client := http.Client{Timeout: 3 * time.Second}

	var req reqUrl = reqUrl{
		w:   &wg,
		m:   &mu,
		suc: true,
	}

	for _, v := range urls {
		wg.Add(1)
		go req.requestUrl(v, client)
		if int8(runtime.NumCPU()/4) == i {
			wg.Wait()
			i = 0
		}
		i++
	}
	wg.Wait()

	return &req
}

func main() {
	res := callUrls([]string{
		"https://gobyexample.com/",
		"https://gobyexample.com/",
		"https://stackoverflow.com/",
		"https://stackoverflow.com/",
	})
	if res.suc {
		for _, v := range res.res {
			time.Sleep(1 * time.Second)
			fmt.Println("Status: ", v.Status)
			fmt.Println("Status Code: ", v.StatusCode)
			fmt.Println("Header", v.Header)
			fmt.Println("Body: ", v.Body)
		}
	} else {
		fmt.Println("One Url Get Error!")
	}
}
