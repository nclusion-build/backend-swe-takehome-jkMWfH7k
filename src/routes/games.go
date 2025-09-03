package routes

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterGameRoutes(r *gin.Engine) {
    g := r.Group("/games")
    // TODO: Implement games endpoints: create, get, status, join, moves, stats, list, delete
    g.GET(":id/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "TODO"})
    })
}


// TODO: Add validation for game creation and move inputs [ttt.todo.validation.game-inputs]