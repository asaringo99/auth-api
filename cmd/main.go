package main

import (
	"net/http"

	"go.uber.org/fx"

	login_controller "github.com/asaringo99/authapi/auth/controller/login"
	signup_controller "github.com/asaringo99/authapi/auth/controller/signup"
	token_controller "github.com/asaringo99/authapi/auth/controller/token"
	"github.com/asaringo99/authapi/auth/model"
	"github.com/asaringo99/authapi/auth/repository"
	login_usecase "github.com/asaringo99/authapi/auth/usecase/login"
	signup_usecase "github.com/asaringo99/authapi/auth/usecase/signup"
	token_usecase "github.com/asaringo99/authapi/auth/usecase/token"
	"github.com/asaringo99/authapi/db"
	"github.com/asaringo99/authapi/http/handler"
	"github.com/asaringo99/authapi/http/router"
)

func main() {
	fx.New(
		fx.Provide(
			db.NewAuthDbConnection,
			token_controller.NewTokenController,
			token_usecase.NewTokenUsecase,
			login_controller.NewLoginController,
			login_usecase.NewLoginUsecase,
			signup_controller.NewSignUpController,
			signup_usecase.NewSignUpUsecase,
			repository.NewAuthRepository,
			model.NewAuthModel,
			router.NewEchoServer,
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"routes"`),
			),
			handler.AsHandler(handler.NewSignUpHandler),
			handler.AsHandler(handler.NewLoginHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
