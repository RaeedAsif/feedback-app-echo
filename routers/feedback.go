package routers

import (
	"github.com/RaeedAsif/feedback-app-echo/controllers"
	"github.com/RaeedAsif/feedback-app-echo/middleware"
	"github.com/RaeedAsif/feedback-app-echo/store"
	"github.com/labstack/echo/v4"
)

func initFeedbackRouters(e *echo.Echo) {
	e.POST("/feedback", middleware.AuthJWTMiddleware(createFeedback))
	e.GET("/feedbacks", middleware.AuthJWTMiddleware(getUserFeedbacks))
}

// createFeedback godoc
// @Summary      POST API which creates feedback
// @Description  post feedback
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ResponseSuccess{id=int}
// @Failure      403  {object}  json.ResponseError
// @Failure      500  {object}  json.ResponseError
// @Param request body models.FeedbackInput true "query params"
// @Security     ApiKeyAuth
// @Router       /feedback [post]
func createFeedback(c echo.Context) error {
	fbController := controllers.NewFeedbackController(store.DB)
	return fbController.CreateFeedback(c)
}

// getUserFeedbacks godoc
// @Summary      GET API which returs users feedbacks
// @Description  get feedbacks
// @Accept       json
// @Produce      json
// @Param 	  	 type query string false "string valid"
// @Param 	  	 page query int true "int valid"
// @Success      200  {object}  json.ResponseSuccess{data=models.FeedbackListResponse}
// @Failure      404  {object}  json.ResponseError
// @Failure      403  {object}  json.ResponseError
// @Failure      500  {object}  json.ResponseError
// @Security     ApiKeyAuth
// @Router       /feedbacks [get]
func getUserFeedbacks(c echo.Context) error {
	fbController := controllers.NewFeedbackController(store.DB)
	return fbController.GetFeedbacks(c)
}
