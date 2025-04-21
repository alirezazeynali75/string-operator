package api

import (
    "net/http"
    "sync"

    "github.com/gin-gonic/gin"
)

func (h *Handlers) ProcessAll(c *gin.Context) {
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
    errChan := make(chan error, 4)
    results := struct {
        UppercasedText string
        LowercasedText string
        ReversedText   string
    }{}

    // Perform uppercase operation
    wg.Add(1)
    go func() {
        defer wg.Done()
        var err error
        results.UppercasedText, err = h.uppercaseSvc.Do(c.Request.Context(), input.Text)
        if err != nil {
            errChan <- err
        }
    }()

    // Perform lowercase operation
    wg.Add(1)
    go func() {
        defer wg.Done()
        var err error
        results.LowercasedText, err = h.lowercaseSvc.Do(c.Request.Context(), input.Text)
        if err != nil {
            errChan <- err
        }
    }()

    // Perform reverse operation
    wg.Add(1)
    go func() {
        defer wg.Done()
        var err error
        results.ReversedText, err = h.reverseSvc.Do(c.Request.Context(), input.Text)
        if err != nil {
            errChan <- err
        }
    }()

    // Perform increase operation
    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := h.increaseSvc.Inc(c.Request.Context()); err != nil {
            errChan <- err
        }
    }()

    // Wait for all goroutines to finish
    go func() {
        wg.Wait()
        close(errChan)
    }()

    // Handle errors if any
    for err := range errChan {
        h.logger.Error("error in concurrent execution", "error", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process input"})
        return
    }

    // Return the results
    c.JSON(http.StatusOK, gin.H{
        "uppercased_text": results.UppercasedText,
        "lowercased_text": results.LowercasedText,
        "reversed_text":   results.ReversedText,
    })
}