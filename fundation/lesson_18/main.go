package main

import "fmt"

// interfaces vazias
// quando a interface não tem métodos ela é implementada em todo mundo
// type x interface{}

func main() {
	//interfaces vazias burlam a tipagem do go
	var x interface{} = 10
	var y interface{} = "Hello World"
	showType(x)
	showType(y)
	y = 20
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v \n", t, t)
}
