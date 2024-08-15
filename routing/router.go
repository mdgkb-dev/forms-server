package routing

import (
	"mdgkb/ankets-server/handlers/auth"
	"mdgkb/ankets-server/handlers/humans"
	"mdgkb/ankets-server/handlers/users"
	"mdgkb/ankets-server/modules/forms"
	authRouter "mdgkb/ankets-server/routing/auth"

	humansRouter "mdgkb/ankets-server/routing/humans"
	usersRouter "mdgkb/ankets-server/routing/users"

	"github.com/gin-gonic/gin"

	helperPack "github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/middleware"
	baseRouter "github.com/pro-assistance/pro-assister/routing"
)

func Init(r *gin.Engine, helper *helperPack.Helper) {
	m := middleware.CreateMiddleware(helper)
	api, apiNoToken := baseRouter.Init(r, helper)
	// api.Use(m.InjectClaims())
	api.Use(m.InjectFTSP())

	auth.Init(helper)
	authRouter.Init(apiNoToken.Group("/auth"), auth.H)

	users.Init(helper)
	usersRouter.Init(api.Group("/users"), users.H)

	humans.Init(helper)
	humansRouter.Init(api.Group("/humans"), humans.H)

	forms.InitRoutes(api, helper)
}
