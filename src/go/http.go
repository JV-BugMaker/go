package main

import (
	"sync"
	"net/http"
	"io/ioutil"
)
func main() {
	var wg sync.WaitGroup
	for i:=0;i<3;i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			resp, err := http.Get("http://www.baidu.com")
			if err != nil {
				// handle error
				println(err)
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
				println(err)
			}

			println(body)
		}(i)
	}

	wg.Wait()

	println("main done")
}

