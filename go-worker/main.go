package main

import (
	"net/http"
	"time"

	tm "github.com/0xBradock/go-worker/taskmanager"
)

func main() {
	// Setup
	repo := tm.NewInMemoryRepo()
	manager := tm.NewTaskManager(repo)

	// Create bound task functions
	sendEmail := tm.EmailTasks{SendEmail: tm.CreateTask(manager, tm.SendEmail, tm.TasksOptions{
		Retries: 2,
		Timeout: 5 * time.Second,
	})}

	hndl := NewPurchaseHandler(sendEmail)
	router := http.NewServeMux()
	router.HandleFunc("/purchase", hndl.purchaseHandler)
}

type PurchaseHandler struct {
	emailTask tm.EmailTasks
}

func NewPurchaseHandler(et tm.EmailTasks) *PurchaseHandler {
	return &PurchaseHandler{emailTask: et}
}

func (h *PurchaseHandler) purchaseHandler(w http.ResponseWriter, r *http.Request) {
	// do work
	h.emailTask.SendEmail(tm.EmailParams{
		To:      "",
		From:    "",
		Subject: "",
		Content: "",
	})
}
