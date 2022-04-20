package test

import (
	"Network-be/Server/db"
	"Network-be/Server/handlers"
	"Network-be/config"
	"testing"
)

func TestSendVerifyByEmail(t *testing.T) {
	config.Init()
	err := handlers.SendVerifyByEmail("test@eson.ninja",
		handlers.CreateContent("Ticket-ID-sssssssssssss"))
	t.Log(err)
}

func TestNumber(t *testing.T) {
	config.Init()
	db.InitDB()
	handlers.CreateVerifyTicket("eson", "test@eson.ninja")
}
