package pub

import (
	"fmt"

	"github.com/HuKeping/rbtree"
)

func TestPublishMethod(str string) {
	fmt.Println(str, "Printing this")
	useTree()
}

func useTree() {
	rbt := rbtree.New()
	fmt.Println(rbt)
}
