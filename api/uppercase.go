package api

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func (h *Handlers) Uppercase(c *gin.Context) {
    // Parse input from the request
    var input struct {
        Text string `json:"text" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        h.logger.Error("invalid input", "error", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    // Call uppercase service
    result, err := h.uppercaseSvc.Do(c.Request.Context(), input.Text)
    if err != nil {
        h.logger.Error("failed to uppercase text", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, gin.H{"uppercase_text": result})
}