package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"context"
	"github.com/juani-castore/goWeb/2ndLesson/internal/producto"
)
const (
	puerto = ":8080"
)
var (
	valueContext any = "user"
)


func main() {

	storage := Producto.Storage{
		Products: loadData(),
	}

	storage.PrintInfo()
	
	router := gin.Default()

	// health check
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/productosParams", func(ctx *gin.Context) {
		ctx.Query("id")
		nameQuery := ctx.Query("name")
		ctx.Query("quantity")
		ctx.Query("code_value")
		ctx.Query("expiration")
		ctx.Query("is_published")
		ctx.Query("price")

		// si viene el parametro vacio, devuelvo todos los productos
		if nameQuery == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"data": storage.GetAll(ctx),
			})}

	
		/*
		if nameQuery == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"mensaje": "nombre invalido",
			})
			return
		}
		*/
		// si tengo parametros, busco por nombre
		data := storage.GetProductByName(nameQuery)

		if data == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"mensaje": "no se encontraron productos",
			})
			return
		}


		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})


	// get Products endpoint
	router.GET("/productos/search", func(ctx *gin.Context) {

		precioQuery := ctx.Query("priceGt")
		user := ctx.Query("user")

		if precioQuery != "" {
			precio, err := strconv.ParseFloat(precioQuery, 64)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mensaje": "precio invalido",
				})
				return
			}

			data := storage.GetProductosMayorPrecio(precio)
			ctx.JSON(http.StatusOK, gin.H{
				"data": data,
			})
			return
		}

		nuevoContexto := addToContext(ctx, user)

		ctx.JSON(http.StatusOK, gin.H{
			"data": storage.GetAll(nuevoContexto),
		})
	})

	router.Run(puerto)

}

func loadData() []producto.Producto {
	productos := []producto.Producto{
		{
			ID:          1,
			Name:        "Banana",
			CodeValue:   "AABBCCC",
			Quantity:    10,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.0,
		},
		{
			ID:          2,
			Name:        "Manzana",
			CodeValue:   "AABBDDD",
			Quantity:    5,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.0,
		},
		{
			ID:          3,
			Name:        "Pera",
			CodeValue:   "AAZZZCCC",
			Quantity:    8,
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.0,
		},
	}

	return productos
}

func addToContext(ctx context.Context, user string) context.Context {
	nuevoContexto := context.WithValue(ctx, valueContext, user)
	return nuevoContexto
}