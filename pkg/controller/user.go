package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"zvaljean/proxy-subscribe-api/pkg/config"
	"zvaljean/proxy-subscribe-api/pkg/entity"
	"zvaljean/proxy-subscribe-api/pkg/log"
	"zvaljean/proxy-subscribe-api/pkg/model"
	"zvaljean/proxy-subscribe-api/pkg/service"
	"zvaljean/proxy-subscribe-api/pkg/utils"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		userService: service.NewUserService(model.NewUserModel(db))}
}

func (u *UserController) FindUserByToken(ctx *gin.Context) {

	token := ctx.Param("token")
	path := ctx.Param("path")

	if len(token) == 0 || len(path) == 0 {
		log.Error("user or conf param empty: user: %s, conf: %s", token, path)
		ctx.String(http.StatusNotFound, "")
		return
	}

	user, err := u.userService.FindUserByToken(token)
	if user == nil || err != nil {
		log.Error("FindUserByToken-error: token: %s, err: %s", token, err)
		ctx.String(http.StatusNotFound, "")
		return
	}

	marshal, _ := json.Marshal(user)
	log.Info("user-data: %s", string(marshal))

	utils.RespStr(ctx, utils.StrtoBase64(user.Data))

}

func (u *UserController) FindUserByTokenV1(ctx *gin.Context) {

	token := ctx.Param("token")
	path := ctx.Param("path")
	ty := ctx.Param("type")

	if len(token) == 0 || len(path) == 0 || len(ty) == 0 {
		log.Error("user or conf param empty: user: %s, conf: %s", token, path)
		ctx.String(http.StatusNotFound, "")
		return
	}

	typ, _ := strconv.Atoi(ty)

	param := &entity.UserDto{
		Type:  typ,
		Token: token,
		Path:  path,
	}

	user, err := u.userService.FindUserByTokenTypePath(param)
	if user == nil || err != nil {
		log.Error("FindUserByToken-error: token: %s, err: %s", token, err)
		ctx.String(http.StatusNotFound, "")
		return
	}

	marshal, _ := json.Marshal(user)
	log.Info("user-data: %s", string(marshal))

	var resp string

	switch typ {
	case config.X_UI:
		resp = utils.StrtoBase64(user.Data)
	case config.HY2:
		resp = user.Data
	}

	utils.RespStr(ctx, resp)

}

func (u *UserController) SetupRouter(engine *gin.Engine) {

	oldApi := engine.Group(config.Proxy)
	{
		oldApi.GET(":token/:path", u.FindUserByToken)
	}

	newApi := engine.Group(config.ApiV1)
	{
		newApi.GET("/:token/:type/:path", u.FindUserByTokenV1)
	}

}
