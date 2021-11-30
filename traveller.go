package traveller

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Traveller struct {
}

func (t *Traveller) Run() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./website", true)))

	api := router.Group("/api")
	api.GET("/time", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"time": time.Now(),
		})
	})

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("error in running traveller: %s", err.Error())
	}
}
