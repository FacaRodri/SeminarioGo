package main
import (
	"fmt"
	
	"seminario/otracarpeta"
)



func main() {
	fmt.Printf("quien?\n")

	c := otracarpeta.SumarValores(4, 5) 
	auto1 := otracarpeta.NewAuto("1A", "Verde", "Ford")
	auto2 := otracarpeta.NewAuto("2B", "Gris", "Ford")
	auto3 := otracarpeta.NewAuto("3C", "Plateado", "VMW")
	auto4 := otracarpeta.NewAuto("4D", "Negro", "VMW")
	agencia := otracarpeta.NewAgencia()

	//Agrego autos
	agencia.AddAuto(auto1, auto2, auto3, auto4)
	fmt.Println(agencia.GetAutos())
	//Borro el auto en la posicion pasada
	agencia.RemoveAuto(1)
	//Edito Color
	agencia.EditColor("1A", "Verde Claro")
	fmt.Println(agencia.GetAutos())
	fmt.Println(c)
}