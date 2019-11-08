package otracarpeta

type Auto struct {
	Patente   string
	Color     string
	Modelo    string
	Velocidad int
}

func NewAuto (Patente, Color, Modelo string) Auto{
	return Auto{
		Patente: Patente,
		Color: Color,
		Modelo: Modelo,
	}
}

func Acelerar(a *Auto) () {
	a.Velocidad += 10
}

func Frenar(a *Auto) () {
	a.Velocidad -= 10
	
}



