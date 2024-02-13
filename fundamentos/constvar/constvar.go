package main

// é possível usar um alias para os imports para ter nomes personalizados.
// caso seja necessário importar algo e não utilizar basta usar o caracter _.
import (
	"fmt"
	m "math"
)

/*
Em go não se faz necessário inferir os tipos porém é possível conforme exemplos a baixo.
*/
func main() {
	const PI float64 = 3.1415 //constantes em go.
	var raio = 3.2            //valor inferido pelo navegador.

	area := PI * m.Pow(raio, 2) //forma reduzida de se criar uma var.

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	) // forma de se criar varias constantes em um unico bloco.

	var (
		c = 3
		d = 4
	) // forma de se criar varias variaveis em um unico bloco.

	fmt.Println(c, d) //TODO o compilador do go não permite variaveis criadas e não utilizadas gerando i, erro de compilação

	// é possível criar e atribuir ao mesmo tempo duas ou mais variaveis. Ou mesma Linha!
	var e, f bool = true, false
	//na forma reduzida
	g, h := true, "false"

	fmt.Println(e, f, g, h)
}
