package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)
const (
	puerto = ":8080"
)

type Producto struct {
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Precio float32 `json:"precio"`
	Stock int `json:"stock"`
	Codigo string `json:"codigo"`
	Publicado bool `json:"publicado"`
	FechaDeCreacion string `json:"fechaCreacion"`
}

func main() {

	router := gin.Default()
	var productos, err = getProductos()
	
	router.GET("/productos", func(c *gin.Context){

		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"productos": productos,
		})

	})

	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(puerto)
}


func getProductos() ([]Producto, error){
	
	var arrayProductos = []Producto{
		Producto{Id: 1, Nombre: "Producto 1", Precio: 1000, Stock: 100, Codigo: "0001", Publicado: true, FechaDeCreacion: "2021-10-10"},
		Producto{Id: 2, Nombre: "Producto 2", Precio: 2000, Stock: 200, Codigo: "0002", Publicado: true, FechaDeCreacion: "2021-10-10"},
		Producto{Id: 3, Nombre: "Producto 3", Precio: 3000, Stock: 300, Codigo: "0003", Publicado: true, FechaDeCreacion: "2021-10-10"},
		Producto{Id: 4, Nombre: "Producto 4", Precio: 4000, Stock: 400, Codigo: "0004", Publicado: true, FechaDeCreacion: "2021-10-10"},
		Producto{Id: 5, Nombre: "Producto 5", Precio: 5000, Stock: 500, Codigo: "0005", Publicado: true, FechaDeCreacion: "2021-10-10"},
		Producto{Id: 6, Nombre: "Producto 6", Precio: 6000, Stock: 600, Codigo: "0006", Publicado: true, FechaDeCreacion: "2021-10-10"},
	}
	jsonData, err := json.Marshal(arrayProductos)
	if err != nil {
		return nil, err
	}
	// imprimo el json para ver que se genero correctamente
	fmt.Println(string(jsonData))

	return arrayProductos, err
}










