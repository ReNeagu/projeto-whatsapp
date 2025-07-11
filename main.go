package main

import (
	"log"

	"github.com/ReNeagu/projeto-whatsapp/internal/db"
	"github.com/ReNeagu/projeto-whatsapp/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Conectar() // conecta ao PostgreSQL
	r := gin.Default()

	//Definirei os endpoints depois
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	//Rota do webhook
	r.POST("/webhook", handlers.HandleWebhook)

	r.GET("/leads", handlers.ListarLeads)

	log.Println("Servidor rodando na porta 8080")
	r.Run(":8080") // inicia na porta 8080

}
