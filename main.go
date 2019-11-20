package main

import (
	
	"encoding/json"
	"fmt"
	"net/http"
	"seminario/SeminarioGo/otracarpeta"

	"github.com/gin-gonic/gin"
)

var agencia otracarpeta.Agencia

func main() {
	fmt.Printf("quien?\n")

	c := otracarpeta.SumarValores(4, 5)
	auto1 := otracarpeta.NewAuto("1A", "Verde", "Ford")
	auto2 := otracarpeta.NewAuto("2B", "Gris", "Ford")
	auto3 := otracarpeta.NewAuto("3C", "Plateado", "VMW")
	auto4 := otracarpeta.NewAuto("4D", "Negro", "VMW")
	agencia = otracarpeta.NewAgencia()

	//Agrego autos
	agencia.AddAuto(auto1, auto2, auto3, auto4)
	fmt.Println(agencia.GetAutos())
	//Borro el auto en la posicion pasada
	agencia.RemoveAuto(1)
	//Edito Color
	agencia.EditColor("1A", "Verde Claro")
	fmt.Println(agencia.GetAutos())
	fmt.Println(c)

	//Punteros
	var m *int
	n := 10
	l := 20
	m = &n
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(l)
	fmt.Println(&l)

	//Json
	myString := "Hola como andas"

	fmt.Println("MyString", myString)

	miAuto := auto1

	miAutoMarshalled, _ := json.Marshal(miAuto)

	fmt.Println("MiAutoMarshalled", miAutoMarshalled)

	//Api
	router := gin.Default()

	router.GET("/hello", GetHello)

	router.GET("/agencia/:patente", BuscarAuto)

	if err := router.Run(); err != nil {
		fmt.Println(err)
	}

}

//CONTROLLER
func BuscarAuto(ctx *gin.Context)  {
	patente := ctx.Param("patente")
	result := agencia.BuscarAuto(patente)
	ctx.JSON(http.StatusOK, gin.H{
		"Resultado": result,
	})
}
/////

func GetHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
