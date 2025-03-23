package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
  "gorm.io/gorm"
	"os"
	"time"
	"fmt"
)

var db *gorm.DB

func dbConnect(dialector gorm.Dialector, config gorm.Option, count int) (err error) {
	x := 0
	for x < count {
		if db, err = gorm.Open(dialector, config); err != nil {
			time.Sleep(time.Second * 2)
			x++
			fmt.Printf("retry count:%v\n", x)
			continue
		}
		fmt.Println("Connected to database")
		break
	}
	return err
}

type item struct {
	gorm.Model
	Text string `gorm:"size:256;not null" json:"text"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// GET /api/items
	api := r.Group("/api")
	api.GET("/items", func(c *gin.Context) {
		var items []item
		db.Find(&items)
		c.JSON(http.StatusOK, items)
	})
	// POST /api/items
	api.POST("/items", func(c *gin.Context) {
		var item item
		c.BindJSON(&item)
		if result := db.Create(&item); result.Error != nil {
			panic(result.Error)
		}
		c.JSON(http.StatusCreated, item)
	})
	// PUT /api/items/:id
	api.PUT("/items/:id", func(c *gin.Context) {
		var item item
		id := c.Param("id")
		if result := db.First(&item, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Item not found",
			})
			return
		}
		c.BindJSON(&item)
		if result := db.Save(&item); result.Error != nil {
			panic(result.Error)
		}
		c.JSON(http.StatusOK, item)
	})
	// DELETE /api/items/:id
	api.DELETE("/items/:id", func(c *gin.Context) {
		id := c.Param("id")
		result := db.Delete(&item{}, id)
		if result.Error != nil {
			panic(result.Error)
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Item not found",
			})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})
	return r
}

func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	dialector := mysql.Open(dsn)
	option := &gorm.Config{}
	if err := dbConnect(dialector, option, 10); err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&item{})
	r := setupRouter()
	r.Run(":" + os.Getenv("API_PORT"))
}
