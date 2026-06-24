package handler

import "backend/internal/service"


type DriverHandler struct {
	driverService *service.DriverService
}

func NewDriverHandler(driverService *service.DriverService) *DriverHandler {
	return &DriverHandler{
		driverService:  driverService,
	}
}