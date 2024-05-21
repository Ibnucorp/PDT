package main

import (
	"ecommerce/controller"
	"ecommerce/db"

	"github.com/gin-gonic/gin"
)

func main(){
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    db,err := db.InitDB()

    if err != nil{
        panic(err)
    }

    router.GET("/", func(ctx *gin.Context){
        ctx.JSON(200, gin.H{
            "message":"Hello World",
        })
    })


    router.POST("/products", controller.CreateProduct(db))
    router.GET("/products",controller.ShowAllProducts(db))
    router.GET("/products/:id",controller.GetProduct(db))
    router.PUT("/products/:id",controller.UpdateProduct(db))
    router.DELETE("/products/:id",controller.DeleteProduct(db))
    router.Run(":8080")
}