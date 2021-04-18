package main

import (
	"blackspeak/article"
	"fmt"
)

func main() {
	a, err := article.RandomArticle(20)
	if err != nil {
		panic("error")
	}
	fmt.Println(a)
}
