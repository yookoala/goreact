package goreact

import (
	"github.com/robertkrimen/otto"

	"fmt"
	"log"
	"testing"
)

func TestJs(t *testing.T) {
	vm := otto.New()
	script := renderReactElem("CommentBox", "comments", "_test/js/main.js")

	_, err := vm.Run(script)
	if err != nil {
		log.Fatal(err)
	}

	val, err := vm.Get("comments")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mustStr(val.ToString()))
}

func TestJsx(t *testing.T) {
	vm := otto.New()
	script := renderReactElem("CommentBox", "comments", "_test/js/main.jsx")

	_, err := vm.Run(script)
	if err != nil {
		log.Fatal(err)
	}

	val, err := vm.Get("comments")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mustStr(val.ToString()))
}
