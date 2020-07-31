package main

import (
	"fmt"
	"time"

	"imooc.com/xiangnan/learngo/retriever/mock"
	"imooc.com/xiangnan/learngo/retriever/real"
)

const url = "http://www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "xiangnan", "course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

func download(r Retriever) string {
	return r.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:74.0) Gecko/20100101 Firefox/74.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session:")
	fmt.Println(session(&retriever))
}
