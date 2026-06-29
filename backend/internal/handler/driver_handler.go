package handler

import (
	"backend/internal/service"
	"net/http"
	"strings"
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

func (h *DriverHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {
     
	authHeader := r.Header.Get("Authorization")
	if authHeader == ""{
		util.ErrJson(w, ErrHeaderMissing)
	}

	authToken:= strings.SplitN(authHeader, " ")

}
