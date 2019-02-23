# goex

This repo contains my exercises and experiments in Golang. If you already have Golang environment setup (e.g. `$GOPATH`):

* To get the whole repo, run `go get github.com/davidfeng88/goex/...`.
* To see the doc in browser: first run `godoc -http=:999`, then in browser go to `http://localhost:999/pkg/github.com/davidfeng88/goex/`.

## poetry

A package used in John Graham-Cumming's [
Introduction to Go Programming](http://shop.oreilly.com/product/0636920035305.do) course. It demostrates the basics of Golang, including types, slices, interfaces, `strings`, `fmt`, `bufio`, etc.

* To run the tests: `go test github.com/davidfeng88/goex/poetry` (Use `-v` for verbose)

## poetryserver

Also a package used in John Graham-Cumming's [
Introduction to Go Programming](http://shop.oreilly.com/product/0636920035305.do) course. It uses the `poetry` package, and demostrates the usages of `encoding/json`, `flag`, `log`, `net/http`, `strconv`, `sync`, `time`, etc. The server loads config file and poem file, and serves the poem.

To run the server:

1. `go install github.com/davidfeng88/goex/poetryserver`
2. Copy the sample config file `config`, and the poem file `words` to your `$GOPATH`
3. Whe working directory `$GOPATH`, run `bin/poetryserver`. You should see server logs in terminal.
4. In terminal, run `curl http://127.0.0.1:8088/poem\?name=words` or go to `http://localhost:8088/poem?name=words` in browser.

## gopl

Exercises from *the Go Programming Language*.
