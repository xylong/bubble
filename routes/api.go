package routes

import (
	v1 "bubble/app/http/controllers/api/v1"
	"bubble/app/http/controllers/api/v1/auth"
	"github.com/xylong/bingo"
)

func init() {
	registerController(v1.NewUserCtrl())
	registerController(auth.NewSignupController())
}

var Controllers = make([]bingo.Controller, 0)

func registerController(controller bingo.Controller) {
	Controllers = append(Controllers, controller)
}
