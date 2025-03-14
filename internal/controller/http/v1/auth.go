package v1

// import (
// 	"net/http"

// 	"ai-seller/internal/entity"
// 	"ai-seller/internal/usecase"
// 	"ai-seller/pkg/logger"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/gofiber/fiber/v2"
// )

// type authRoutes struct {
// 	t usecase.Auth
// 	l logger.Interface
// 	v *validator.Validate
// }

// func NewAuthRoutes(apiV1Group *gin.RouterGroup, t usecase.Auth, l logger.Interface) {
// 	r := &authRoutes{t, l, validator.New(validator.WithRequiredStructEnabled())}

// 	authGroup := apiV1Group.Group("/auth")
// 	{
// 		authGroup.GET("/history", r.history)
// 		authGroup.POST("/do-translate", r.doTranslate)
// 	}
// }


// // @Summary     Show history
// // @Description Show all translation history
// // @ID          history
// // @Tags  	    translation
// // @Accept      json
// // @Produce     json
// // @Success     200 {object} historyResponse
// // @Failure     500 {object} response
// // @Router      /translation/history [get]
// func (r *authRoutes) history(ctx *fiber.Ctx) error {
// 	translations, err := r.t.CreateUser(ctx.UserContext())
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - history")

// 		return errorResponse(ctx, http.StatusInternalServerError, "database problems")
// 	}

// 	return ctx.Status(http.StatusOK).JSON(historyResponse{translations})
// }

// type doTranslateRequest struct {
// 	Source      string `json:"source"       validate:"required"  example:"auto"`
// 	Destination string `json:"destination"  validate:"required"  example:"en"`
// 	Original    string `json:"original"     validate:"required"  example:"текст для перевода"`
// }

// // @Summary     Translate
// // @Description Translate a text
// // @ID          do-translate
// // @Tags  	    translation
// // @Accept      json
// // @Produce     json
// // @Param       request body doTranslateRequest true "Set up translation"
// // @Success     200 {object} entity.Translation
// // @Failure     400 {object} response
// // @Failure     500 {object} response
// // @Router      /translation/do-translate [post]
// func (r *authRoutes) doTranslate(ctx *fiber.Ctx) error {
// 	var request doTranslateRequest

// 	if err := ctx.BodyParser(&request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")

// 		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
// 	}

// 	if err := r.v.Struct(request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")

// 		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
// 	}

// 	translation, err := r.t.CreateUser(
// 		ctx.UserContext(),
// 		entity.Translation{
// 			Source:      request.Source,
// 			Destination: request.Destination,
// 			Original:    request.Original,
// 		},
// 	)
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")

// 		return errorResponse(ctx, http.StatusInternalServerError, "translation service problems")
// 	}

// 	return ctx.Status(http.StatusOK).JSON(translation)
// }
