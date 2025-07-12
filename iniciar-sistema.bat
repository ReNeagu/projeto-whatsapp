@echo off
title Iniciando Sistema WhatsApp Leads
echo --------------------------------------
echo Iniciando PostgreSQL...
echo --------------------------------------

REM Inicie o PostgreSQL (se não estiver rodando como serviço automático)
REM Se estiver instalado como serviço, pode pular esta parte
REM Se usa o pgAdmin + Postgre 17 com serviço automático, comente esta linha:
net start postgresql-x64-17

echo --------------------------------------
echo Iniciando Aplicação Go (.exe)
echo --------------------------------------
start "" whatsapp-leads.exe

timeout /t 3
start http://localhost:8080/painel

exit
