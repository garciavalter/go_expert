package main

func main() {
	a := 10
	b := 2
	c := 3
	d := 4
	//os operadores comparativos são iguais javascript (&& || != ==)
	if a > b && c > a {
		println(a)
	} else {
		println(b)
	}
	//compara o valor de a e imprime e executa a ação do case
	switch a {
	case 1:
		println(a)
	case 2:
		println(b)
	case 3:
		println(c)
	default:
		println(d)
	}

}
