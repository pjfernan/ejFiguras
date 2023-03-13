package figuras

type Circulo struct {
	Radio float32
}

func (c *Circulo) Area() float32 {
	const pi = 3.141592654
	return pi * c.Radio * c.Radio
}

func (c *Circulo) Perimetro() float32 {
	const pi = 3.141592654
	return pi * c.Radio * 2
}
