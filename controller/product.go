package controller

import (
	"ecommerce/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CreateProduct(db *gorm.DB) gin.HandlerFunc{
    return   func(ctx *gin.Context) {
        // Konversi string ke float
        price,err := strconv.ParseFloat(ctx.PostForm("price"), 64)

        if err != nil{
            ctx.JSON(400,"Price must be a numbers") // 400 : invalid syntax request to server
        }
        // Konversi String ke Integer
        stock,err := strconv.Atoi(ctx.PostForm("stock"))

        if err != nil{
            ctx.JSON(400,"Stock must be numbers")
        }

        // Jika table products belum di database, maka dibuatkan
        db.AutoMigrate(models.Product{})

        // Memasukkan nilai yang diinput user ke variable newProduct
        newProduct := models.Product{
            Name: ctx.PostForm("name"),
            Price: price,
            Stock: stock,
        }
        
        db.Create(&newProduct) // Simpan data di table
        ctx.JSON(201,newProduct) // Kirim JSON ke user, sukses input
    }
}

func ShowAllProducts(db *gorm.DB) gin.HandlerFunc{
    return func(ctx *gin.Context){
        var products []models.Product
        
        db.Find(&products)

        ctx.JSON(200,products)
    }

}

func GetProduct(db *gorm.DB) gin.HandlerFunc{
    return func(ctx *gin.Context){
        ID := ctx.Param("id")
        var product models.Product

        if err := db.First(&product, "id=?",ID).Error; err != nil {
            ctx.JSON(404, gin.H{"message":"Product ID not found"})
            return
        }

        ctx.JSON(200,product)
    }
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc{
    return func(ctx *gin.Context) {
        ID := ctx.Param("id")
        var product models.Product

        if err := db.First(&product, "id=?", ID).Error;err != nil{
            ctx.JSON(400, gin.H{"message":"Product ID not found"}) // 400 : invalid syntax request to server
        }

        price,err := strconv.ParseFloat(ctx.PostForm("price"), 64)

        if err != nil{
            ctx.JSON(400,"Price must be numbers")
        }

        stock,err := strconv.Atoi(ctx.PostForm("stock"))

        if err != nil{
            ctx.JSON(400,"Stock must be numbers")
        }
        var updatedProduct = models.Product{
            ID      : product.ID,
            Name    : ctx.PostForm("name"),
            Price   : price,
            Stock   : stock,
        }

        db.Model(&product).Updates(updatedProduct)
        ctx.JSON(200, product)
    }
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc{
    return func(ctx *gin.Context) {
        ID := ctx.Param("id")
        var product models.Product

        if err := db.First(&product, "id=?",ID).Error; err != nil{
            ctx.JSON(404, gin.H{"message":"Product ID not found."})
            return
        }
        db.Delete(&product)
        ctx.JSON(200, gin.H{"message":"Product succesfully deleted."})
    }
}