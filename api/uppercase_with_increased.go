package api

import (
    "net/http"
    "sync"

    "github.com/gin-gonic/gin"
)

func (h *Handlers) UppercaseWithIncrease(c *gin.Context) {
    // Parse input from the request
    var input struct {
        Text string `json:"text" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        h.logger.Error("invalid input", "error", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    var wg sync.WaitGroup
    errChan := make(chan error, 2)
    var data string

    // Call uppercaseSvc.Do concurrently
    wg.Add(1)
    go func() {
        defer wg.Done()
        var err error
        data, err = h.uppercaseSvc.Do(c.Request.Context(), input.Text)
        if err != nil {
            errChan <- err
        }
    }()

    // Call increase concurrently
    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := h.increaseSvc.Inc(c.Request.Context()); err != nil {
            errChan <- err
        }
    }()

    // Wait for both goroutines to finish
    wg.Wait()
    close(errChan)

    // Handle errors if any
    for err := range errChan {
        h.logger.Error("error in concurrent execution", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the result
    c.JSON(http.StatusOK, gin.H{"trimmed_text": data})
}