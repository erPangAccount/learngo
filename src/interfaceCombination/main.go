package main

import (
	"fmt"
	"interfaceCombination/mock"
)

type Retriever interface {
	MyGet(url string) string
}

type Poster interface {
	MyPost(url string, form map[string]int) string
}

func get(r Retriever, url string) string {
	return r.MyGet(url)
}

func post(poster Poster, url string, form map[string]int) string {
	return poster.MyPost(url, form)
}

type Session interface {
	Poster
	Retriever
	Put(url string, form map[string]int) string
}

func add(session Session, url string, form map[string]int) string {
	fmt.Println(session.MyPost(url, form))
	return session.MyGet(url)
}

func update(session Session, url string, form map[string]int) string {
	fmt.Println(session.Put(url, form))
	return session.MyGet(url)
}

func main() {
	var retriever Retriever
	retriever = &mock.User{}
	fmt.Println(get(retriever, "a")) //[]

	var poster Poster
	poster = &mock.User{}
	fmt.Println(post(poster, "a", map[string]int{ // ok
		"id": 123,
	}))

	var sessioner Session
	sessioner = &mock.User{}
	fmt.Println(add(sessioner, "a", map[string]int{ // ok
		"id": 1234,
	}))
	//ok
	//[1234 1234]

	fmt.Println(update(sessioner, "a", map[string]int{ // ok
		"id": 12345,
	}))
	//ok
	//[12345 1234]

	user := &mock.User{
		List: []int{1, 2, 3, 4, 5},
	}
	fmt.Println(user) //User.List = [1 2 3 4 5]
}
