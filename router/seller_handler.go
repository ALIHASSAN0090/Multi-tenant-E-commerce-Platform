package router

import (
	"ecommerce-platform/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetStore(c *gin.Context) {

	seller_id, exists := c.Get("Id")
	if !exists {
		c.JSON(400, gin.H{"error": "Seller ID not found"})
		return
	}

	storeData, err := r.SellerController.GetStore(seller_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in getting Store Data",
				Detail:     err.Error(),
			},
		})
	} else {

		c.JSON(http.StatusOK, models.SuccessResponse{
			Data:       storeData,
			Message:    "Store Data fetched Successfully",
			SubMessage: "Store Data fetched Successfully",
			StatusCode: http.StatusOK,
		})

	}
}
