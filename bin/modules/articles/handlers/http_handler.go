package handlers

import (
	"fmt"
	"kumparan-backend-position-interview/bin/modules/articles/models/binding"
	repository "kumparan-backend-position-interview/bin/modules/articles/repositories"
	"kumparan-backend-position-interview/bin/modules/articles/usecases"
	database "kumparan-backend-position-interview/bin/pkg/databases"
	"kumparan-backend-position-interview/bin/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// HTTPHandler struct
type HTTPHandler struct {
	commandUsecase usecases.ArticleCommandUsecaseInterface
	queryUsecase   usecases.ArticleQueryUsecaseInterface
}

// New initiation
func New() *HTTPHandler {
	db := database.InitSQLite()

	commandDb := repository.NewDBCommand(db)
	queryDb := repository.NewDBQuery(db)
	commandUsecase := usecases.NewArticleCommandUsecase(commandDb, queryDb)
	queryUsecase := usecases.NewArticleQueryUsecase(commandDb, queryDb)

	return &HTTPHandler{
		commandUsecase: commandUsecase,
		queryUsecase:   queryUsecase,
	}
}

// Mount function
func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.POST("/v1/articles", h.Create)
	echoGroup.GET("/v1/articles", h.GetList)
	echoGroup.GET("/v1/articles/search", h.Search)
}

// Command
func (h *HTTPHandler) Create(c echo.Context) error {
	payload := new(binding.Create)
	if err := utils.BindValidate(c, payload); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErr {
				errorMessage := utils.GenerateErrorMessage(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
				return utils.Response(nil, errorMessage, http.StatusBadRequest, c)
			}
		}
	}
	result := h.commandUsecase.Create(c.Request().Context(), payload)
	if result.Error != nil {
		utils.ResponseError(result.Error, c)
	}
	return utils.Response(nil, "Success", http.StatusOK, c)
}

func (h *HTTPHandler) GetList(c echo.Context) error {
	filtered := map[string]any{}
	if author := c.QueryParam("author"); author != "" {
		filtered["author_id"] = author
	}
	result := h.queryUsecase.GetList(c.Request().Context(), filtered)
	if result.Error != nil {
		utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Success", http.StatusOK, c)
}

func (h *HTTPHandler) Search(c echo.Context) error {
	payload := new(binding.Search)
	b, _ := c.Get("payload").([]byte)
	fmt.Println("stringg", string(b))
	if err := utils.BindValidate(c, payload); err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErr {
				errorMessage := utils.GenerateErrorMessage(fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
				return utils.Response(nil, errorMessage, http.StatusBadRequest, c)
			}
		}
	}
	result := h.queryUsecase.Search(c.Request().Context(), *payload)
	if result.Error != nil {
		utils.ResponseError(result.Error, c)
	}
	return utils.Response(result.Data, "Success", http.StatusOK, c)
}
