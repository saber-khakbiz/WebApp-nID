package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPerson godoc
// @Summary Get City name  Information by nID
// @Description Get City Information by National Identification
// @Tags ID
// @ID get-city-by-id
// @Accept  json
// @Produce  html
// @Param id path string true "National ID"
// @Header 200 {string} Token "qwerty"
// @Router /nids/{id} [get]
func GetNid(c *gin.Context) {
	if nid := c.Param("id"); nid != "" {
		// Check if the person exists
		if person, err := getId(nid); err == nil && person != nil {
			c.JSON(http.StatusOK, person)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}


