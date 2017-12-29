package actions

import (
	"github.com/gobuffalo/buffalo"
	"net/http"
)

func I18NHandler(c buffalo.Context) error {
	key := c.Param("key")
	val, err := T.Translate(c, key)
	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(val))
}
