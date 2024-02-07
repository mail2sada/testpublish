package main

import (
	"fmt"

	"github.com/HuKeping/rbtree"
	"github.com/mail2sada/testpublish/pub"
)

func main() {
	fmt.Println("Demo: Test publish")
	pub.TestPublishMethod("Hello")
	useTree()
}

func useTree() {
	rbt := rbtree.New()
	fmt.Println(rbt)
}
