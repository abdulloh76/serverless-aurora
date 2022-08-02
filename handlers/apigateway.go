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

type APIGatewayHandler struct {
	users *domain.Users
}

func NewAPIGatewayHandler(d *domain.Users) *APIGatewayHandler {
	return &APIGatewayHandler{
		users: d,
	}
}

func (h *APIGatewayHandler) AllHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	allUsers, err := h.users.AllUsers()
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, allUsers), nil
}

func (h *APIGatewayHandler) GetHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func (h *APIGatewayHandler) CreateHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func (h *APIGatewayHandler) PutHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func (h *APIGatewayHandler) DeleteHandler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func response(code int, object interface{}) events.APIGatewayProxyResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func errResponse(status int, body string) events.APIGatewayProxyResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)

	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}
}
