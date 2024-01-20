package v1

import (
	"errors"
	"net/http"

	"github.com/eduardolat/tinylink/internal/echoutil"
	"github.com/eduardolat/tinylink/internal/shortener"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *handlers) retrieveHandler(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return echoutil.JsonError(c, http.StatusBadRequest, err)
	}

	link, err := h.shortener.Get(id)
	if errors.Is(err, shortener.ErrLinkNotFound) {
		return echoutil.JsonError(c, http.StatusNotFound, err)
	}
	if err != nil {
		return echoutil.JsonError(c, http.StatusInternalServerError, err)
	}

	return c.JSON(
		http.StatusOK,
		h.linkToJSON(link),
	)
}
