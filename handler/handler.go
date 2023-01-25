package handler

import "fmt"

type GetHandler struct {
}

func NewGetHanlder() *GetHandler {
	return &GetHandler{}
}

func (h *GetHandler) Hello() {
	fmt.Println("Hello World")
}
