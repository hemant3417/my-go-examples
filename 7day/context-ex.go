package main

import (
	"context"
	"fmt"
)

func main() {
	c := context.Background()

	fmt.Println("context \t", c)
	fmt.Println("context error \t", c.Err)
	fmt.Printf("context type \t%T\n", c)
	fmt.Println("----------")
	ctx, cancel := context.WithCancel(c)
	fmt.Println("context \t", ctx)
	fmt.Println("context error \t", ctx.Err())
	fmt.Printf("context type \t%T\n", ctx)
	fmt.Println("context \t", cancel)
	fmt.Printf("context type \t%T\n", cancel)
	fmt.Println("----------")
	cancel()
	fmt.Println("context \t", ctx)
	fmt.Println("context error \t", ctx.Err())
	fmt.Printf("context type \t%T\n", ctx)
	fmt.Println("context \t", cancel)
	fmt.Printf("context type \t%T\n", cancel)
	fmt.Println("----------")

}
