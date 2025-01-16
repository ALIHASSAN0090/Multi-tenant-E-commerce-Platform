package router

import (
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Router) GetAllOrders(c *gin.Context) {

	id, err := utils.GetContextId(c)
	utils.HandleError(err)

	data, err := r.SellerController.GetAllOrders(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in getting orders Data",
				Detail:     err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       data,
		Message:    "Orders Data fetched Successfully",
		SubMessage: "Orders Data fetched Successfully",
		StatusCode: http.StatusOK,
	})

}
func (r *Router) GetStoreItems(c *gin.Context) {

	sellerID, valid := GetContextID(c)
	if !valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid Seller ID",
				Detail:     "Seller ID not found or invalid",
			},
		})
		return
	}

	storeData, err := r.SellerController.GetStoreItems(sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in getting Store Data",
				Detail:     err.Error(),
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

func (r *Router) GetStoreItem(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64) //items table id pk
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid ID format",
				Detail:     "ID should be a valid integer",
			},
		})
		return
	}

	item, err := r.SellerController.GetStoreItem(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling create seller controller",
				Detail:     err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       item,
		Message:    "Store Item fetched Successfully",
		SubMessage: "Store Item fetched Successfully",
		StatusCode: http.StatusOK,
	})
}

func (r *Router) UpdateItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) //items table id pk
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid ID format",
				Detail:     "ID should be a valid integer",
			},
		})
		return
	}

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid JSON format",
				Detail:     err.Error(),
			},
		})
		return
	}

	updated_item, err := r.SellerController.UpdateStoreItem(c, id, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling UpdateStoreItem controller",
				Detail:     err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       updated_item,
		Message:    "Store Item updated and fetched Successfully",
		SubMessage: "Store Item updated and fetched Successfully",
		StatusCode: http.StatusOK,
	})
}

func (r *Router) CreateItem(c *gin.Context) {

	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid JSON format",
				Detail:     err.Error(),
			},
		})
		return
	}

	sellerID, valid := GetContextID(c)
	if !valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid Seller ID",
				Detail:     "Seller ID not found or invalid",
			},
		})
		return
	}

	created_item, err := r.SellerController.CreateItem(c, sellerID, item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling CreateItem controller",
				Detail:     err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       created_item,
		Message:    "Store Item created and fetched Successfully",
		SubMessage: "Store Item created and fetched Successfully",
		StatusCode: http.StatusOK,
	})
}

func (r *Router) GetStore(c *gin.Context) {

	sellerID, valid := GetContextID(c)
	if !valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid Seller ID",
				Detail:     "Seller ID not found or invalid",
			},
		})
		return
	}

	store, err := r.SellerController.GetStore(c, sellerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: models.Error{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error in calling CreateItem controller",
				Detail:     err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Data:       store,
		Message:    "Store fetched Successfully",
		SubMessage: "Store fetched Successfully",
		StatusCode: http.StatusOK,
	})
}
