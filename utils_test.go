package goreact

import (
	"testing"
)

func TestJs(t *testing.T) {
	html, err := renderElemToString("CommentBox", "_test/js/main.js")
	if err != nil {
		t.Error(err.Error())
	}
	if html == "" {
		t.Error("Empty html result")
	}
	t.Log("TestJs: \n", html)
}

func TestJsx(t *testing.T) {
	html, err := renderElemToString("CommentBox", "_test/js/main.jsx")
	if err != nil {
		t.Error(err.Error())
	}
	if html == "" {
		t.Error("Empty html result")
	}
	t.Log("TestJsx: \n", html)
}
