package controller

import "fmt"

type Handler struct {
}

func NewHanlder() *Handler {
	return &Handler{}
}

func (h *Handler) Hello() {
	fmt.Println("Hello World")
}
