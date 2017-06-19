package beego_middleware

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func FunctionBeforeStatic(ctx *context.Context) {
	beego.Info("beego.BeforeStatic: Before finding the static file")
}
func FunctionBeforeRouter(ctx *context.Context) {
	beego.Info("beego.BeforeRouter: Executing Before finding router")
}
func FunctionBeforeExec(ctx *context.Context) {
	beego.Info("beego.BeforeExec: After finding router and before executing the matched Controller")
}
func FunctionAfterExec(ctx *context.Context) {
	beego.Info("beego.AfterExec: After executing Controller")
}
func FunctionFinishRouter(ctx *context.Context) {
	beego.Info("beego.FinishRouter: After finishing router")
}

func InitMiddleware() {
	beego.InsertFilter("*", beego.BeforeStatic, FunctionBeforeStatic, false)
	beego.InsertFilter("*", beego.BeforeRouter, FunctionBeforeRouter, false)
	beego.InsertFilter("*", beego.BeforeExec, FunctionBeforeExec, false)
	beego.InsertFilter("*", beego.AfterExec, FunctionAfterExec, false)
	beego.InsertFilter("*", beego.FinishRouter, FunctionFinishRouter, false)
}
