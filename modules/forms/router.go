package forms

import (
	"mdgkb/ankets-server/modules/forms/handlers/answervariants"
	"mdgkb/ankets-server/modules/forms/handlers/fieldfills"
	"mdgkb/ankets-server/modules/forms/handlers/fields"
	"mdgkb/ankets-server/modules/forms/handlers/formfills"
	"mdgkb/ankets-server/modules/forms/handlers/forms"
	"mdgkb/ankets-server/modules/forms/handlers/formsections"
	"mdgkb/ankets-server/modules/forms/handlers/selectedanswervariants"

	answervariantsRouter "mdgkb/ankets-server/modules/forms/routing/answervariants"
	fieldfillsRouter "mdgkb/ankets-server/modules/forms/routing/fieldfills"
	fieldsRouter "mdgkb/ankets-server/modules/forms/routing/fields"
	formFillsRouter "mdgkb/ankets-server/modules/forms/routing/formfills"
	formsRouter "mdgkb/ankets-server/modules/forms/routing/forms"
	formsectionsRouter "mdgkb/ankets-server/modules/forms/routing/formsections"
	selectedanswervariantsRouter "mdgkb/ankets-server/modules/forms/routing/selectedanswervariants"

	"github.com/gin-gonic/gin"
	helperPack "github.com/pro-assistance/pro-assister/helper"
)

func InitRoutes(api *gin.RouterGroup, helper *helperPack.Helper) {
	formsections.Init(helper)
	formsectionsRouter.Init(api.Group("/form-sections"), formsections.H)

	forms.Init(helper)
	formsRouter.Init(api.Group("/forms"), forms.H)

	fields.Init(helper)
	fieldsRouter.Init(api.Group("/fields"), fields.H)

	fieldfills.Init(helper)
	fieldfillsRouter.Init(api.Group("/field-fills"), fieldfills.H)

	answervariants.Init(helper)
	answervariantsRouter.Init(api.Group("/answer-variants"), answervariants.H)

	formfills.Init(helper)
	formFillsRouter.Init(api.Group("/form-fills"), formfills.H)

	selectedanswervariants.Init(helper)
	selectedanswervariantsRouter.Init(api.Group("/selected-answer-variants"), selectedanswervariants.H)
}
