package handler

import (
	"net/http"
	"warehouse/api/models"

	"github.com/gin-gonic/gin"
)

// Income Warehouse godoc
// @ID Income_Warehouse_
// @Router /do-income/{coming_id} [POST]
// @Summary Send Coming
// @Description Send Coming
// @Tags Send Coming
// @Accept json
// @Procedure json
// @Param coming_id path string true "coming_id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Income_Warehouse(c *gin.Context) {

	var (
		id      = c.Param("coming_id")
		respPrd models.RespProduct
	)

	check, err := h.storage.Coming().GetByID(c.Request.Context(), &models.ComingPrimaryKey{
		Id: id,
	})

	if check == nil {
		h.handlerResponse(c, "Not Found This Coming!", http.StatusInternalServerError, nil)
		return
	}

	if err != nil {
		h.handlerResponse(c, "Not Found This Coming", http.StatusInternalServerError, err.Error())
		return
	}
	if check.Status == "finished" {
		h.handlerResponse(c, "This Coming has already FINISHED!!!", http.StatusOK, nil)
		return
	}

	_, err = h.storage.Coming().Update(c, &models.UpdateComing{
		Id:       check.Id,
		ComingID: check.ComingID,
		BranchID: check.BranchID,
		DateTime: check.DateTime,
		Status:   "finished",
	})

	product_data, err := h.storage.ComingProducts().GetByComingID(c.Request.Context(), &models.ComingPrimaryKey{
		Id: id,
	})
	if product_data == nil {
		h.handlerResponse(c, "storage.comingProducts.Not_Found", http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		h.handlerResponse(c, "storage.comingProducts.get", http.StatusInternalServerError, err.Error())
		return
	}

	respPrd.BranchID = check.BranchID
	respPrd.ComingID = id
	respPrd.CategoryID = product_data.CategoryID
	respPrd.Name = product_data.Name
	respPrd.Barcode = product_data.Barcode
	respPrd.Price = product_data.Price
	respPrd.Amount = product_data.Amount
	respPrd.TotalPrice = product_data.TotalPrice

	resp2, err := h.storage.Remainder().DoIncomeWareHouse(c, &respPrd)

	if err != nil {
		h.handlerResponse(c, "Error Coming Products in WhareHouse", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "Doing Coming Products in WhareHouse", http.StatusAccepted, resp2)

}
