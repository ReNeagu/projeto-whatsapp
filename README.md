WhatsApp Leads Manager - Sistema de Registro e Distribuição de Contatos

Este é um sistema desenvolvido em Go (Golang) com PostgreSQL e interface web, criado para registrar automaticamente contatos recebidos via campanhas de WhatsApp e distribuí-los entre vendedores de forma organizada. O projeto tem foco em portfólio pessoal e demonstração de competências em backend, integração com banco de dados e renderização HTML.

📊 Funcionalidades

✉️ 1. Recebimento automático de contatos (simulado por webhook)

Endpoint POST /webhook recebe dados no formato JSON:

{
  "from": "5599999999999",
  "name": "João Silva",
  "message": "Quero saber mais",
  "time": 1720538134
}

O sistema registra os dados recebidos no banco PostgreSQL

Atribui automaticamente um vendedor de forma rotativa (round-robin)

📅 2. Registro no banco de dados PostgreSQL

Os contatos são armazenados na tabela leads com os campos:

phone, name, message, created_at, vendedor

🔎 3. Painel HTML interativo (/painel)

Rota GET /painel renderiza um painel web responsivo com os leads

➜ Exibe dados em ordem cronológica decrescente

💾 4. Exportação em CSV

Rota GET /exportar gera um arquivo CSV com todos os registros

Pode ser usado para backup ou importação em outros sistemas

🔧 5. Executável standalone

Projeto pode ser compilado em .exe e executado localmente

➜ Acompanha .bat para iniciar PostgreSQL + sistema automaticamente

🚀 Tecnologias Utilizadas

Go (Golang) - backend leve e performático

Gin - framework HTTP

PostgreSQL - banco de dados relacional

HTML + template - para renderização do painel

CSV nativo - para exportação de dados

💡 Possíveis Expansões

Integração com WhatsApp real (via WPPConnect)

Gráfico de leads por vendedor

Login para vendedores

Deploy em servidor externo

👥 Para quem esse projeto é interessante?

Profissionais de tecnologia que querem ver backend funcional

Recrutadores interessados em aplicações reais com Go

Pequenas empresas que precisam organizar leads

Estudantes querendo aprender integração Go + PostgreSQL + HTML

🛎þ Como executar localmente

Requisitos:

Go instalado (v1.20+)

PostgreSQL instalado e rodando

Passos:

# Clone o projeto
git clone https://github.com/SEU-USUARIO/projeto-whatsapp.git
cd projeto-whatsapp

# Configure as variáveis do banco (se necessário no db/db.go)

# Rode diretamente
go run main.go

# Ou crie o .exe
go build -o whatsapp-leads.exe
whatsapp-leads.exe

Acesse:

http://localhost:8080/painel

📘 Licença

Este projeto foi desenvolvido como parte de estudo e portfólio, e está sob a licença MIT. Sinta-se livre para usar, contribuir e adaptar!

✨ Contato

Desenvolvido por Renato Neagu

LinkedIn: linkedin.com/in/renato-neagu

GitHub: github.com/ReNeagu

Se este projeto te ajudou ou te inspirou, deixe uma estrela ⭐ no repositório!

