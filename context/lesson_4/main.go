package main

import (
	"context"
	"fmt"
)

func main() {
	//usando contexto para passar parametros.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

// por convenção as variáveis de contexto sempre são o primeiro
// parametro nas funções.
func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
