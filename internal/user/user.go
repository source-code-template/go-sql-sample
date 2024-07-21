package user

import (
	"context"
	"database/sql"
	"net/http"

	v "github.com/core-go/core/v10"
	"github.com/core-go/search/query"

	"go-service/internal/user/handler"
	"go-service/internal/user/model"
	"go-service/internal/user/repository/adapter"
	"go-service/internal/user/service"
)

type UserTransport interface {
	Search(w http.ResponseWriter, r *http.Request)
	All(w http.ResponseWriter, r *http.Request)
	Load(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Patch(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(db *sql.DB, logError func(context.Context, string, ...map[string]interface{})) (UserTransport, error) {
	validator, err := v.NewValidator()
	if err != nil {
		return nil, err
	}

	buildQuery := query.UseQuery[model.User, *model.UserFilter](db, "users")
	userRepository, err := adapter.NewUserAdapter(db, buildQuery)
	if err != nil {
		return nil, err
	}
	userService := service.NewUserService(db, userRepository)
	userHandler := handler.NewUserHandler(userService, validator.Validate, logError, nil)
	return userHandler, nil
}
