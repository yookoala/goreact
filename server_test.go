package main

import (
	"github.com/robertkrimen/otto"

	"fmt"
	"log"
	"testing"
)

func TestServer(t *testing.T) {
	vm := otto.New()
	vm.Run(renderReactComp("CommentBox", "comments", "static/js/main.js"))
	val, err := vm.Get("comments")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mustStr(val.ToString()))
}
