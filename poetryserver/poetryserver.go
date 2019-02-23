package main

import (
	"encoding/json"
	"flag"
	_ "fmt"
	"log"
	"net/http"
	"os"
	_ "sort"
	"strconv"
	"sync"
	"time"

	"github.com/davidfeng88/goex/poetry"
	// add a _ before it so that compiler won't complain if it's not used
)

type protectedCache struct {
	sync.Mutex // no field name, but Lock and Unlock will be available on the struct
	c          map[string]poetry.Poem
}

var cache protectedCache

type config struct {
	Route       string   // in the JSON file it could be lowercase route
	BindAddress string   `json:"addr"`
	ValidPoems  []string `json:"valid"`
}

type poemWithTitle struct {
	Title     string // when writing json, only uppercase fields are exported
	Body      poetry.Poem
	WordCount string
	TheCount  int
}

var c config

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	poemName := r.Form["name"][0]

	p, ok := cache.c[poemName]

	if !ok {
		http.Error(w, "Not found (invalid)", http.StatusNotFound)
		return
	}

	log.Printf("User requested poem %s\n", poemName)

	// fmt.Fprintf(w, "%s\n", p) output text poem
	// sort.Sort(p[0])
	// output json poem
	pwt := poemWithTitle{poemName, p,
		strconv.FormatInt(int64(p.NumWords()), 16), p.NumThe()}
	enc := json.NewEncoder(w)
	enc.Encode(pwt)
}

func main() {
	// stage 1:
	// p = Poem{{"This is a poam"}}
	// v, c := p.Stats()
	// fmt.Printf("Vowels: %d, Consonnants: %d\n", v, c)
	// fmt.Printf("Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())

	// stage 2:
	// LoadPoem from file
	// p, err := poetry.LoadPoem("words")
	// if err != nil {
	// 	fmt.Printf("Error loading poem: %s\n", err)
	// }
	// fmt.Printf("%s\n", p)
	// fmt.Printf("%#v\n", p) // show the types

	// stage 3:
	// poem server

	// config the logger
	log.SetFlags(log.Lmicroseconds)

	// config flag method 1
	configFilename := flag.String("conf", "config", "Name of config file")

	// config flag method 2
	// var configFilename string
	// flag.StringVar(&configfilename, "conf", "config", "Name of config file")

	flag.Parse()

	f, err := os.Open(*configFilename)
	if err != nil {
		log.Fatalf("Failed to open file config\n") // will os.Exit(1)
	}

	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	f.Close()
	if err != nil {
		log.Fatalf("Bad JSON\n")
	}

	cache.c = make(map[string]poetry.Poem)

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, name := range c.ValidPoems {
		wg.Add(1)
		go func(n string) {
			// protect the cache map, only one goroutine can access it at one time
			cache.Lock()
			defer cache.Unlock()

			cache.c[n], err = poetry.LoadPoem(n)
			if err != nil {
				log.Fatalf("Failed to load poem %s\n", n)
			}

			wg.Done()
		}(name)
	}
	// wait until all reading is finished before serving
	wg.Wait()

	elapsed := time.Since(startTime)
	// alternative:
	// elapsed := time.Now().Sub(startTime)
	log.Printf("Loading took %s\n", elapsed)

	log.Printf("Server started at %s\n", time.Now().Format(time.Kitchen))

	http.HandleFunc(c.Route, poemHandler)
	http.ListenAndServe(c.BindAddress, nil)
	// http://localhost:8088/poem in browser or
	// curl http://127.0.0.1:8088/poem\?name=abc
}

/*
// in command line
export GOPATH=`pwd`
go install hello
bin/hello
*/
