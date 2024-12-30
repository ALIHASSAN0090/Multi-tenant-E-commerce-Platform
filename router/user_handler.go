package router

import (
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) CreateOrder(c *gin.Context) {
	var orderData models.CreateOrder

	id, err := utils.GetContextId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in getting context id",
				Detail:     err.Error(),
			},
		})
		return
	}
	if err := c.ShouldBindJSON(&orderData); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to bind JSON",
				Detail:     err.Error(),
			},
		})
		return
	}

	orderData.Order.UserID, orderData.Order.CreatedBy = id, id

	data, err := r.UserController.CreateOrder(c, orderData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling create Order controller",
				Detail:     err.Error(),
			},
		})
	} else {

		c.JSON(http.StatusOK, models.SuccessResponse{
			Data:       data,
			Message:    "Order Created Successfully",
			SubMessage: "Order Created Successfully",
			StatusCode: http.StatusOK,
		})
	}

}
func (r *Router) CreateSeller(c *gin.Context) {

	var data models.SellerStore

	idInterface, exists := c.Get("Id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found in context"})
		return
	}

	idUint, ok := idInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID is not a uint"})
		return
	}

	id := int64(idUint)

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to bind JSON",
				Detail:     err.Error(),
			},
		})
		return
	}

	data.UserID = id

	sellerData, storeData, err := r.UserController.CreateSellerStore(c, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling create seller controller",
				Detail:     err.Error(),
			},
		})
	} else {

		c.JSON(http.StatusOK, models.SuccessResponse{
			Data: gin.H{
				"seller": sellerData,
				"store":  storeData,
			},
			Message:    "Seller and Store Created Successfully",
			SubMessage: "Seller and Store Created Successfully",
			StatusCode: http.StatusOK,
		})
	}
}

func (r *Router) GetStores(c *gin.Context) {

	stores, err := r.UserController.GetStores(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling create seller controller",
				Detail:     err.Error(),
			},
		})
	} else {

		c.JSON(http.StatusOK, models.SuccessResponse{
			Data:       stores,
			Message:    "Store Successfully",
			SubMessage: "Seller and Store Created Successfully",
			StatusCode: http.StatusOK,
		})
	}

}

func (r *Router) GetStoreAndItems(c *gin.Context) {

	store_id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid store ID",
				Detail:     err.Error(),
			},
		})
		return
	}

	store, err := r.UserController.GetStoreItems(c, store_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling create seller controller",
				Detail:     err.Error(),
			},
		})
	} else {

		c.JSON(http.StatusOK, models.SuccessResponse{
			Data:       store,
			Message:    "Store Successfully",
			SubMessage: "Seller and Store Created Successfully",
			StatusCode: http.StatusOK,
		})
	}
}
