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

	r.LoadHTMLGlob("templates/*") //carrega o template HTML da pasta /templates

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	//Rotas
	r.POST("/webhook", handlers.HandleWebhook)
	r.GET("/leads", handlers.ListarLeads)
	r.GET("/painel", handlers.MostrarPainel)
	r.GET("/exportar", handlers.ExportarLeadsCSV)

	log.Println("Servidor rodando na porta 8080")
	r.Run(":8080") // inicia na porta 8080

}
