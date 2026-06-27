package handler

import (
	"backend/internal/service"
	"net/http"
)

type DriverHandler struct {
	driverService *service.DriverService
}

func NewDriverHandler(driverService *service.DriverService) *DriverHandler {
	return &DriverHandler{
		driverService: driverService,
	}
}

type CreateDriverReqiest struct{
	name string
	
}

func (h *DriverHandler) CreateaRider(w http.ResponseWriter, r *http.Request) {
    
}
