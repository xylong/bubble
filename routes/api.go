package routes

import (
	v1 "bubble/app/http/controllers/api/v1"
	"github.com/xylong/bingo"
)

func init() {
	registerController(v1.NewUserCtrl())
}

var Controllers = make([]bingo.Controller, 0)

func registerController(controller bingo.Controller) {
	Controllers = append(Controllers, controller)
}
