package main

import (
	"fmt"
	"strings"
)

type equation struct {
	a, b int
}

func (e equation) String() string {
	return fmt.Sprintf("%d • %d", e.a, e.b)
}

func (e equation) Product() int {
	return e.a * e.b
}

func (e equation) DisplayQuotient() string {
	ans := e.a * e.b
	return fmt.Sprintf("%d / %d = %d", ans, e.a, e.b)
}

func (e equation) DisplayProduct() string {
	ans := e.a * e.b
	return fmt.Sprintf("%d • %d = %d", e.a, e.b, ans)
}

type productList []*equation

func (list productList) DisplayProduct() string {
	var buff []string
	for _, e := range list {
		buff = append(buff, e.DisplayProduct())
	}
	return strings.Join(buff, "\n")
}

func (list productList) String() string {
	var buff []string
	for _, e := range list {
		buff = append(buff, e.DisplayQuotient())
	}
	return strings.Join(buff, "\n")
}

func main() {
	products := make(map[int]productList)
	fmt.Println(products) // FIXME commit = death
	for i := 2; i <= 12; i++ {
		for j := 2; j <= 12; j++ {
			e := &equation{i, j}
			products[e.Product()] = append(products[e.Product()], e)
			fmt.Println(e.DisplayQuotient())
			fmt.Println(e.DisplayProduct())
		}
	}
	for k, v := range products {
		fmt.Printf("Quotients of %d\n", k)
		fmt.Println(v.String())
		fmt.Printf("Products of %d\n", k)
		fmt.Println(v.DisplayProduct())
	}
}
