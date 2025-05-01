package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/Unchana19/go-learn/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,min=3,max=255,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type UserWithToken struct {
	User  store.User `json:"user"`
	Token string     `json:"token"`
}

// RegisterUser godoc
//
//	@Summary	Register a new user
//	@Tags		authentication
//	@Accept		json
//	@Produce	json
//	@Param		payload	body		RegisterUserPayload	true	"User payload"
//	@Success	201		{object}	UserWithToken
//	@Failure	400		{object}	error
//	@Router		/v1/authentication/register [post]
func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequest(w, r, err)
		return
	}

	user := &store.User{
		Username: payload.Username,
		Email:    payload.Email,
		IsActive: false,
	}

	// Hash password
	if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	plainToken := uuid.New().String()

	hash := sha256.Sum256([]byte(plainToken))
	hashToken := hex.EncodeToString(hash[:])

	// Create user
	err := app.store.Users.CreateAndInvite(ctx, user, hashToken, app.config.mail.exp)
	if err != nil {
		switch err {
		case store.ErrDuplicateUsername:
			app.badRequest(w, r, err)
			return
		case store.ErrDuplicateEmail:
			app.badRequest(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	userWithToken := UserWithToken{*user, plainToken}

	if err := app.jsonResponse(w, http.StatusCreated, userWithToken); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
