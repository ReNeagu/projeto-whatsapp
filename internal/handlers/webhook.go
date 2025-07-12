package handlers

import (
	"encoding/csv"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ReNeagu/projeto-whatsapp/internal/db"
	"github.com/ReNeagu/projeto-whatsapp/internal/utils"
	"github.com/gin-gonic/gin"
)

type WhatsAppPayload struct {
	From    string `json:"from"`    //numero do telefone
	Name    string `json:"name"`    //nome do contato
	Message string `json:"message"` //mensagem enviada
	Time    int64  `json:"time"`    //timestamp UNIX
}

// Lead representa um registro no banco
type Lead struct {
	ID        int       `json:"id"`
	Phone     string    `json:"phone"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Vendedor  string    `json:"vendedor"`
}

// Função que lida com POST /webhook
func HandleWebhook(c *gin.Context) {
	var payload WhatsAppPayload

	//Tentativa de leitura do JSON recebido
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Erro ao ler JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}
	//se não vier timestamp, usamos o atual
	if payload.Time == 0 {
		payload.Time = time.Now().Unix()
	}
	//escolhe o vendedor
	vendedor := utils.ProximoVendedor()

	_, err := db.DB.Exec(`
		INSERT INTO leads (phone, name, message, created_at, vendedor)
		VALUES ($1, $2, $3, $4, $5)
		`, payload.From, payload.Name, payload.Message, time.Unix(payload.Time, 0), vendedor)

	if err != nil {
		log.Println("Erro ao inserir no banco:", err)
		c.JSON(500, gin.H{"error": "erro ao salvar no banco"})
		return
	}

	log.Println("Novo Contato Recebido:")
	log.Println("De:", payload.From)
	log.Println("Nome:", payload.Name)
	log.Println("Mensagem:", payload.Message)
	log.Println("Data:", time.Unix(payload.Time, 0).Format(time.RFC3339))
	log.Println("Vendedor atribuído:", vendedor)

	//Verificar se Name, From ou Message estão vazios:
	if payload.From == "" || payload.Name == "" || payload.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios ausentes"})
		return
	}

	//Envia a resposta final
	// c.JSON(http.StatusOK, gin.H{"status": "recebido com sucesso", "vendedor": vendedor})

	//Envia a resposta final
	c.JSON(200, gin.H{
		"status":   "recebido com sucesso",
		"vendedor": vendedor,
	})

}

// GET /leads
func ListarLeads(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, phone, name, message, created_at, vendedor FROM leads ORDER BY created_at DESC`)
	if err != nil {
		log.Println("Erro ao buscar leads:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar leads"})
		return
	}
	defer rows.Close()

	var leads []Lead
	for rows.Next() {
		var lead Lead
		err := rows.Scan(&lead.ID, &lead.Phone, &lead.Name, &lead.Message, &lead.CreatedAt, &lead.Vendedor)
		if err != nil {
			log.Println("Erro ao ler linha:", err)
			continue
		}
		leads = append(leads, lead)
	}

	c.JSON(http.StatusOK, leads)

}

// GET /painel - renderiza HTML com os leads
func MostrarPainel(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, phone, name, message, created_at, vendedor FROM leads ORDER BY created_at DESC `)
	if err != nil {
		log.Println("Erro ao buscar leads:", err)
		c.String(http.StatusInternalServerError, "Erro ao buscar os dados")
		return
	}
	defer rows.Close()

	var leads []Lead
	for rows.Next() {
		var lead Lead
		if err := rows.Scan(&lead.ID, &lead.Phone, &lead.Name, &lead.Message, &lead.CreatedAt, &lead.Vendedor); err != nil {
			log.Println("Erro ao ler linha:", err)
			continue
		}
		leads = append(leads, lead)
	}
	//renderiza o template com os leads
	c.HTML(http.StatusOK, "painel.html", leads)
}

// ExportarLeadsCSV gera um CSV com todos os leads e força o download
func ExportarLeadsCSV(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, phone, name, message, created_at, vendedor FROM leads ORDER BY created_at DESC`)
	if err != nil {
		log.Println("Erro ao buscar leads para exportação:", err)
		c.String(http.StatusInternalServerError, "Erro ao buscar dados")
		return
	}
	defer rows.Close()

	// Define headers para forçar download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=leads.csv")
	c.Header("Content-Type", "text/csv")

	// Escreve o CSV direto na resposta HTTP
	writer := csv.NewWriter(c.Writer)

	// Cabeçalho do CSV
	writer.Write([]string{"ID", "Telefone", "Nome", "Mensagem", "Data", "Vendedor"})

	// Linhas do CSV
	for rows.Next() {
		var lead Lead
		err := rows.Scan(&lead.ID, &lead.Phone, &lead.Name, &lead.Message, &lead.CreatedAt, &lead.Vendedor)
		if err != nil {
			log.Println("Erro ao ler linha:", err)
			continue
		}

		writer.Write([]string{
			strconv.Itoa(lead.ID),
			lead.Phone,
			lead.Name,
			lead.Message,
			lead.CreatedAt.Format("2006-01-02 15:04:05"),
			lead.Vendedor,
		})
	}

	writer.Flush()

}
