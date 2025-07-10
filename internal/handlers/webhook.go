package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/ReNeagu/projeto-whatsapp/internal/utils"
	"github.com/gin-gonic/gin"
)

type WhatsAppPayload struct {
	From    string `json:"from"`    //numero do telefone
	Name    string `json:"name"`    //nome do contato
	Message string `json:"message"` //mensagem enviada
	Time    int64  `json:"time"`    //timestamp UNIX
}

// Função que lida com POST /webhook
func HandleWebhook(c *gin.Context) {
	var payload WhatsAppPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Erro ao ler JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if payload.Time == 0 {
		payload.Time = time.Now().Unix()
	}

	vendedor := utils.ProximoVendedor()

	log.Println("Novo Contato Recebido:")
	log.Println("De:", payload.From)
	log.Println("Nome:", payload.Name)
	log.Println("Mensagem:", payload.Message)
	log.Println("Data:", time.Unix(payload.Time, 0).Format(time.RFC3339))
	log.Println("Vendedor atribuído:", vendedor)

	c.JSON(http.StatusOK, gin.H{"status": "recebido com sucesso", "vendedor": vendedor})

}
