package middleware

import (
	"encoding/json"
	"net/http"
	"simple_restful_api_golang/model/api"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "SECRET" == request.Header.Get("X-API-KEY") {
		// ok
		// continue to other proces
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		standartResponse := api.StandartResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		writer.Header().Add("Content-Type", "application/json")
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(standartResponse)
		if err != nil {
			panic(err)
		}
	}
}
