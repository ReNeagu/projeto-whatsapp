package utils

import "sync"

//Simulação de vendedores cadastrados
var Vendedores = []string{
	"vendedor_1",
	"vendedor_2",
	"vendedor_3",
	"vendedor_4",
}

var contador int
var mu sync.Mutex // garante concorrência segura

//Função que retorna o próximo vendedor (round-robin)
func ProximoVendedor() string {
	mu.Lock()
	defer mu.Unlock()

	vendedor := Vendedores[contador]
	contador = (contador + 1) % len(Vendedores)
	return vendedor
}
