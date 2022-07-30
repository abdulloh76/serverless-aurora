package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/abdulloh76/serverless-aurora/domain"
	"github.com/abdulloh76/serverless-aurora/store"
	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayV2Handler struct {
	users *domain.Users
}

func NewAPIGatewayV2Handler(d *domain.Users) *APIGatewayV2Handler {
	return &APIGatewayV2Handler{
		users: d,
	}
}

func (h *APIGatewayV2Handler) AllHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	allUsers, err := h.users.AllUsers()
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, allUsers), nil
}

func (h *APIGatewayV2Handler) GetHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	user, err := h.users.GetUser(id)

	if errors.Is(err, store.ErrUserNotFound) {
		return errResponse(http.StatusNotFound, err.Error()), nil
	}
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, user), nil
}

func (h *APIGatewayV2Handler) CreateHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	if strings.TrimSpace(event.Body) == "" {
		return errResponse(http.StatusBadRequest, "empty request body"), nil
	}

	newUser, err := h.users.Create([]byte(event.Body))
	if err != nil {
		if errors.Is(err, domain.ErrJsonUnmarshal) {
			return errResponse(http.StatusBadRequest, err.Error()), nil
		} else {
			return errResponse(http.StatusInternalServerError, err.Error()), nil
		}
	}

	return response(http.StatusCreated, newUser), nil
}

func (h *APIGatewayV2Handler) PutHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	if strings.TrimSpace(event.Body) == "" {
		return errResponse(http.StatusBadRequest, "empty request body"), nil
	}

	product, err := h.users.ModifyUser(id, []byte(event.Body))
	if err != nil {
		if errors.Is(err, domain.ErrJsonUnmarshal) || errors.Is(err, store.ErrUserNotFound) {
			return errResponse(http.StatusBadRequest, err.Error()), nil
		} else {
			return errResponse(http.StatusInternalServerError, err.Error()), nil
		}
	}

	return response(http.StatusCreated, product), nil
}

func (h *APIGatewayV2Handler) DeleteHandler(event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	id, ok := event.PathParameters["id"]
	if !ok {
		return errResponse(http.StatusBadRequest, "missing 'id' parameter in path"), nil
	}

	err := h.users.DeleteUser(id)
	if err != nil {
		if errors.Is(err, store.ErrUserNotFound) {
			return errResponse(http.StatusBadRequest, err.Error()), nil
		} else {
			return errResponse(http.StatusInternalServerError, err.Error()), nil
		}
	}

	return response(http.StatusOK, nil), nil
}

func response(code int, object interface{}) events.APIGatewayV2HTTPResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func errResponse(status int, body string) events.APIGatewayV2HTTPResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}
}
