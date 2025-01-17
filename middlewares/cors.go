package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func Cors() gin.HandlerFunc {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permette tutte le origini, puoi personalizzarlo
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	// Restituisce un middleware gin che usa cors
	return func(c *gin.Context) {
		// Adattiamo il cors per usarlo con Gin
		corsHandler.HandlerFunc(http.ResponseWriter(c.Writer), c.Request)
		c.Next()
	}
}
