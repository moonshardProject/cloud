package main

import (
	"github.com/amirrezaask/mongo_repo"
	"github.com/amirrezaask/sentinel"
)

func main() {
	r := sentinel.New(8080, sentienl.InputTasks{sentinel.InputJSONTask}, sentinel.OutputTasks{sentienl.OutputJSONTask})

	r.POST("/user/register")
	r.POST("/user/login")
	r.POST("/module/add")
	r.POST("/module/send")
	r.POST("module/news")
	r.POST("module/remove")
	// r.Start()
}
