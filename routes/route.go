package routes

import (
	"database/sql"
	api "janio-id-backend/api"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetupRoutes(d *sql.DB) {
	db = d
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("id/orders", GetOrdersIDProvider)
	// router.POST("id/order/update", PostStatus)

	router.Run("localhost:8443")
}

func GetOrdersIDProvider(c *gin.Context) {
	orders, err := api.GetOrdersIDProvider(db)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Orders not found"})
	}
	c.IndentedJSON(http.StatusOK, orders)
}

/*func PostStatus(c *gin.Context) {
	var status models.IDOrderStatus

	c.BindJSON(&status)

	orderStatus, errS := api.PostStatusIDProvider(db, status)
	if errS != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errS.Error()})
	}
	c.IndentedJSON(http.StatusOK, orderStatus)
}*/
