package ejFiguras

type Cuadrado struct {
	Ancho, Alto float32
}

func (c *Cuadrado) Area() float32 {

	return c.Ancho * c.Alto
}

func (c *Cuadrado) Perimetro() float32 {

	return (c.Ancho + c.Alto) * 2
}
