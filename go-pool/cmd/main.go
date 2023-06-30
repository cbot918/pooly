package main

import (
	"fmt"
	"time"

	pool "github.com/cbot918/go-pool"
)

func main() {
	var (
		workNum = 10
		jobNum  = 100
	)
	p := pool.New(workNum, request)

	go func() {
		for i := 0; i < jobNum; i++ {
			p.Push(i)
		}
		p.Close()
	}()

	p.Run()
}

func request(i interface{}) {
	time.Sleep(time.Second * 1)
	fmt.Println("finish request, url num is:", i.(int))
}
