package userhandler

import (
	"context"
	"fmt"
	"time"

	userdb "github.com/relipocere/cafebackend/internal/database/user"
	"github.com/relipocere/cafebackend/internal/model"
	"github.com/relipocere/cafebackend/internal/service/security"
)

const (
	logInBadCredsMessage = "Wrong username or password"
	sessionDuration      = 60 * 24 * time.Hour // 60 days
)

// LogInRequest is the request for LogIn.
type LogInRequest struct {
	Username string
	Password string
}

// LogInResponse is the response of LogIn.
type LogInResponse struct {
	Token string
}

// LogIn handles the login scenario.
func (h *Handler) LogIn(ctx context.Context, req LogInRequest) (LogInResponse, error) {
	resp := LogInResponse{}

	validationErr := validateLogInRequest(req)
	if validationErr != nil {
		return resp, validationErr
	}

	username := req.Username
	password := req.Password
	now := h.now()

	user, err := h.userRepo.Get(ctx, h.db, userdb.GetByUsername(username))
	if err != nil {
		return resp, fmt.Errorf("getting user by username '%s': %w", username, err)
	}

	if user == nil{
		return resp, model.Error{
			Message: logInBadCredsMessage,
			Code:    model.ErrorCodeUnauthenticated,
		}
	}

	err = h.logInValidateCredentials(ctx, *user, password)
	if err != nil {
		return resp, err
	}


	existingSession := user.Session
	if existingSession != nil && existingSession.ExpiresAt.After(now) {
		resp.Token = existingSession.ID
		return resp, nil
	}

	sessionID, err := h.logInCreateSession(ctx, username, now)
	if err != nil {
		return resp, err
	}

	resp.Token = sessionID
	return resp, nil
}

func (h *Handler) logInValidateCredentials(ctx context.Context, user model.User, password string) error {
	isSamePassword := security.IsSameHash(password, user.PasswordHash, user.Salt)
	if !isSamePassword {
		return model.Error{
			Message: logInBadCredsMessage,
			Code:    model.ErrorCodeUnauthenticated,
		}
	}

	return nil
}

func (h *Handler) logInCreateSession(ctx context.Context, username string, now time.Time) (string, error) {
	sessionID, err := security.GenerateSessionID()
	if err != nil {
		return "", fmt.Errorf("session id generation: %w", err)
	}

	err = h.userRepo.SetSession(ctx, h.db, username, model.Session{
		ID: sessionID,
		ExpiresAt: now.Add(sessionDuration),
	})
	if err != nil {
		return "", fmt.Errorf("session creation: %w", err)
	}

	return sessionID, nil
}

func validateLogInRequest(req LogInRequest) error {
	if req.Username == "" {
		return model.Error{
			Message: model.ErrMessageMissingFieldRequired("Username"),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	if req.Password == "" {
		return model.Error{
			Message: model.ErrMessageMissingFieldRequired("Password"),
			Code:    model.ErrorCodeBadRequest,
		}
	}

	return nil
}
