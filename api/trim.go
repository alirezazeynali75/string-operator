package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Trim(c *gin.Context) {
    // Parse input from the request
    var input struct {
        Text string `json:"text" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        h.logger.Error("invalid input", "error", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }
		data, err := h.trimSvc.Do(c.Request.Context(), input.Text)

    // Execute the Trim workflow step
    if err != nil {
        h.logger.Error("failed to execute trim workflow", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, gin.H{"trimmed_text": data})
}