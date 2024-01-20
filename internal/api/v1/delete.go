package v1

import (
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handlers) deleteHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	err = h.shortener.Delete(id)
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted",
	})
}
