package main

import (
	ping "github.com/QiuDDGKD/golang-test/views"
	"github.com/gin-gonic/gin"
)

func runserver() {
	router := gin.Default()
	router.GET("/ping", ping.Get)
	router.POST("/ping", ping.Post)
	router.Run()
}

type TestStruct struct {
	numbers []int
}

func test() {
	a := TestStruct{
		numbers: []int{1, 2, 3},
	}
	b := a
	println("a ptr: ", &a)
	println("b ptr: ", &b)
	println("a.numbers[0] ptr: ", &a.numbers[0])
	println("b.numbers[0] ptr: ", &b.numbers[0])
}

func main() {
	test()
}
