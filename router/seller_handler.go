package router

import (
	"ecommerce-platform/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetStoreItems(c *gin.Context) {

	sellerID, err := GetContextID(c)
	if !err {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid Seller ID",
				Detail:     "Seller ID not found or invalid",
			},
		})
		return
	}
	storeData, err1 := r.SellerController.GetStoreItems(sellerID)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in getting Store Data",
				Detail:     err1.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       storeData,
		Message:    "Store Data fetched Successfully",
		SubMessage: "Store Data fetched Successfully",
		StatusCode: http.StatusOK,
	})
}
