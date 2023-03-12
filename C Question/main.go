package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Requester struct {
	urls   *[]string
	mu     *sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	res    []*http.Response
}

func (re *Requester) request(url string) (*http.Response, error) {
	ctx, cancel := context.WithCancel(re.ctx)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	defer cancel()

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (re *Requester) handelReq(url string) {
	res, err := re.request(url)
	if err != nil {
		fmt.Println(err.Error())
		re.cancel()
		return
	}

	re.mu.Lock()
	re.res = append(re.res, res)
	if len(re.res) == len(*re.urls) {
		defer re.cancel()
	}
	re.mu.Unlock()
}

func (re *Requester) Run() bool {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	re.ctx = ctx
	re.cancel = cancel

	for _, v := range *re.urls {
		go re.handelReq(v)
	}

	<-re.ctx.Done()
	return len(re.res) == len(*re.urls)
}

func NewRequester(urls *[]string) *Requester {
	var mu sync.Mutex
	return &Requester{
		urls: urls,
		mu:   &mu,
	}
}

func main() {
	r := NewRequester(&[]string{
		"https://stackoverflow.com/",
		"https://stackoverflow.com/",
		"https://gobyexample.com/",
		"https://gobyexample.com/",
	})

	if !r.Run() {
		fmt.Println("one Url can't get data!")
	} else {
		for _, v := range r.res {
			time.Sleep(1 * time.Second)
			fmt.Println("Status: ", v.Status)
			fmt.Println("Status Code: ", v.StatusCode)
			fmt.Println("Header", v.Header)
			fmt.Println("Body: ", v.Body)
			fmt.Println()
		}
	}
}
