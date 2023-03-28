package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Utilizando gennerics para alternar entre tipos nos métodos
func SomaGenerics[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

// Criando Constraints
type Number interface {
	// O til serve para considerar types herdardos nas constraints
	// Nesse caso do MyNumber
	~int | ~float64
}

func SomaConstraint[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Usando a constraint comparable
func Compara[T comparable](a, b T) bool {
	//comparable compara apenas igualdades, mas não se são maiores ou menores.
	if a == b {
		return true
	}
	return false

}

func main() {
	m := map[string]int{
		"Wesley": 1000,
		"João":   2000,
		"Maria":  3000,
	}
	m2 := map[string]float64{
		"Wesley": 1000.20,
		"João":   2000.30,
		"Maria":  3000.40,
	}
	m3 := map[string]MyNumber{
		"Wesley": 1000,
		"João":   2000,
		"Maria":  3000,
	}
	resultado := SomaInteiro(m)
	resultado2 := SomaFloat(m2)
	resultado3 := SomaGenerics(m)
	resultado4 := SomaGenerics(m2)
	resultado5 := SomaConstraint(m3)
	println(resultado)
	println(resultado2)
	println(resultado3)
	println(resultado4)
	println(resultado5)
	println(Compara(10, 10))
}
