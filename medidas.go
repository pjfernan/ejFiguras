package ejFiguras

import "fmt"

type Figura interface {
	Area() float32
	Perimetro() float32
}

func AreaFigura(figura Figura) {

	fmt.Println("El área de la figura es: ", figura.Area())
}
func PerimetroFigura(figura Figura) {

	fmt.Println("El perimetro de la figura es: ", figura.Perimetro())
}
