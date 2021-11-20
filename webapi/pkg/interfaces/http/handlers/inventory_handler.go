package handlers

import (
	"encoding/json"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IInventoryHandler interface {
	CreateProduct(httpRequest http.HttpRequest) http.HttpResponse
	GetById(httpRequest http.HttpRequest) http.HttpResponse
}

type inventoryHandler struct {
	logger               interfaces.ILogger
	validator            interfaces.IValidator
	getByIdUseCase       usecases.IGetProductByIdUseCase
	createProductUseCase usecases.ICreateProductUseCase
}

func (pst inventoryHandler) GetById(httpRequest http.HttpRequest) http.HttpResponse {
	id, ok := httpRequest.Params["id"]
	if !ok {
		return http.BadRequest(models.StringToErrorResponse("id is required"), nil)
	}

	result, err := pst.getByIdUseCase.Perform(httpRequest.Ctx, id)
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(models.ToProductResponse(result), nil)
}

func (pst inventoryHandler) CreateProduct(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.CreateProductModel{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		pst.logger.Error(err.Error())
		return http.BadRequest(models.StringToErrorResponse("body is required"), nil)
	}

	if validationErrs := pst.validator.ValidateStruct(model); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return http.BadRequest(models.StringToErrorResponse(validationErrs[0].Message), nil)
	}

	result, err := pst.createProductUseCase.Perform(httpRequest.Ctx, model.ToProductDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Created(models.ToProductResponse(result), nil)
}

func NewInventoryHandler(
	logger interfaces.ILogger,
	validator interfaces.IValidator,
	getByIdUseCase usecases.IGetProductByIdUseCase,
	crateProductUseCase usecases.ICreateProductUseCase,
) IInventoryHandler {
	return inventoryHandler{
		logger,
		validator,
		getByIdUseCase,
		crateProductUseCase,
	}
}
