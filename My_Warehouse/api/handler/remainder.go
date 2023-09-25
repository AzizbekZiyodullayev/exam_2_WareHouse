package handler

import (
	"net/http"
	"warehouse/api/models"
	"warehouse/pkg/helper"

	"github.com/gin-gonic/gin"
)

// Create Remainder godoc
// @ID create_remainder
// @Router /remainder [POST]
// @Summary Create Remainder
// @Description Create Remainder
// @Tags Remainder
// @Accept json
// @Procedure json
// @Param remainder body models.CreateRemainder true "CreateRemainderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateRemainder(c *gin.Context) {

	var createRemainder models.CreateRemainder
	err := c.ShouldBindJSON(&createRemainder)
	if err != nil {
		h.handlerResponse(c, "error remainder should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Remainder().Create(c.Request.Context(), &createRemainder)
	if err != nil {
		h.handlerResponse(c, "storage.remainder.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storage.Remainder().GetByID(c.Request.Context(), &models.RemainderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.remainder.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create remainder resposne", http.StatusCreated, resp)
}

// GetByID remainder godoc
// @ID get_by_id_remainder
// @Router /remainder/{id} [GET]
// @Summary Get By ID Remainder
// @Description Get By ID Remainder
// @Tags Remainder
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdRemainder(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.storage.Remainder().GetByID(c.Request.Context(), &models.RemainderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.remainder.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id remainder resposne", http.StatusOK, resp)
}

// GetList remainder godoc
// @ID get_list_remainder
// @Router /remainder [GET]
// @Summary Get List Remainder
// @Description Get List Remainder
// @Tags Remainder
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param SearchByFilial query string false "search_by_filial"
// @Param SearchByCategory query string false "search_by_category"
// @Param SearchByBarcode query string false "search_by_barcode"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListRemainder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list remainder offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list remainder limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storage.Remainder().GetList(c.Request.Context(), &models.RemainderGetListRequest{
		Offset:           offset,
		Limit:            limit,
		SearchByBranch:   c.Query("search_by_branch"),
		SearchByCategory: c.Query("search_by_category"),
		SearchByBarcode:  c.Query("search_by_barcode"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.remainder.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list remainder resposne", http.StatusOK, resp)
}

// Update remainder godoc
// @ID update_remainder
// @Router /remainder/{id} [PUT]
// @Summary Update Remainder
// @Description Update Remainder
// @Tags Remainder
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param remainder body models.UpdateRemainder true "UpdateRemainderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateRemainder(c *gin.Context) {

	var (
		id              string = c.Param("id")
		updateRemainder models.UpdateRemainder
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateRemainder)
	if err != nil {
		h.handlerResponse(c, "error remainder should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateRemainder.Id = id
	rowsAffected, err := h.storage.Remainder().Update(c.Request.Context(), &updateRemainder)
	if err != nil {
		h.handlerResponse(c, "storage.remainder.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.remainder.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storage.Remainder().GetByID(c.Request.Context(), &models.RemainderPrimaryKey{Id: updateRemainder.Id})
	if err != nil {
		h.handlerResponse(c, "storage.remainder.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create remainder resposne", http.StatusAccepted, resp)
}

// Delete remainder godoc
// @ID delete_remainder
// @Router /remainder/{id} [DELETE]
// @Summary Delete Remainder
// @Description Delete Remainder
// @Tags Remainder
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteRemainder(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.storage.Remainder().Delete(c.Request.Context(), &models.RemainderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.remainder.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete remainder response", http.StatusAccepted, nil)
}
