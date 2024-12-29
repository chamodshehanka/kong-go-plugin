package handlers

import (
	"github.com/Kong/go-pdk"
)

type Config struct {
	Message string
}

type Handler struct{}

func New() interface{} {
	return &Handler{}
}

func (h *Handler) Access(kong *pdk.PDK) {
	err := kong.Log.Info("Access handler invoked......")
	if err != nil {
		kong.Log.Err("Failed to log: %v", err)
		return
	}

	// Add your custom logic here
	err = kong.Response.SetHeader("X-Custom-Header", "CustomValue")
	if err != nil {
		kong.Log.Err("Failed to set header: %v", err)
		return
	}

	kong.Log.Info("Access handler completed......")
}
