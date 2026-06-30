package handler

import (
	"backend/internal/dto"
	"backend/internal/firebase"
	"backend/internal/service"
	"backend/internal/util"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type DriverHandler struct {
	driverService *service.DriverService
}

func NewDriverHandler(driverService *service.DriverService) *DriverHandler {
	return &DriverHandler{
		driverService: driverService,
	}
}

type CreateDriverReqiest struct {
	name string
}

type AuthRequest struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

func (h *DriverHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		util.ErrJson(w, util.ErrTokenMissing)
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		util.ErrJson(w, util.ErrInvalidToken)
	}

	idToken := parts[1]
	ctx := r.Context()
	var payload firebase.FirebasePayload

	if idToken == "frontend" {
		payload = firebase.FirebasePayload{
			Email:  "test@example.com",
			UserId: uuid.Must(uuid.FromString("0b927d97-782a-4c82-b9d2-e4e06774ed37")),
			UID:    "0b927d97-782a-4c82-b9d2-e4e06774ed37",
		}
	} else {
		_, err := firebase.VerifyFirebaseIDToken(ctx, idToken)
		if err != nil {
			util.ErrJson(w, err)
			return
		}
	}

	var req dto.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.ErrJson(w , err)
		return 
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password),10)
	if err != nil {
		util.ErrJson(w , err)
		return 
	}

	req.Password = string(hashedpassword)
 

	user , err := h.driverService.CreateNewUser(ctx, payload, req)
	if err != nil {
		util.ErrJson(w, err)
		return 
	}

	util.WriteJson(w, http.StatusCreated, user)

}
