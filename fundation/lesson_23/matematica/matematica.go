package matematica

// Funcões que começam com letra maíscula são exportadas por default, letras mínusculas ficam disponíveis
// somente dentro do pacote
func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) Andar() string {
	return "Carro Andando"
}
