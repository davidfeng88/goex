// Uses the task package to fetch URLs
package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/davidfeng88/goex/task"
)

var begin time.Time

type httpTask struct {
	url     string
	ok      bool
	elapsed time.Duration
}

func (h *httpTask) Process() {
	if h.url == "" {
		h.ok = false
		return
	}

	start := time.Now()
	resp, err := http.Get(h.url)
	h.elapsed = time.Since(start)
	if err != nil {
		h.ok = false
		return
	}

	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false
}

func (h *httpTask) Output() {
	fmt.Printf("%s %t %s\n", h.url, h.ok, h.elapsed)
}

type httpFactory struct{}

func (f *httpFactory) Create(line string) task.Task {
	h := &httpTask{}
	h.url = line
	return h
}

func main() {
	count := flag.Int("count", 10, "Number of workers")
	flag.Parse()
	f := &httpFactory{}
	task.Run(f, *count)
}
