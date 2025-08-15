package http

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetProfileHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // TODO: use generated sqlc code; mocked payload for starter
        c.JSON(http.StatusOK, gin.H{
            "full_name": "Azamat Serek",
            "title": "Researcher & Engineer",
            "email": "azamat@example.com",
            "about": "Short bio...",
        })
    }
}

func ListExperienceHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) { c.JSON(http.StatusOK, []any{}) }
}
func ListEducationHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) { c.JSON(http.StatusOK, []any{}) }
}
func ListPublicationsHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) { c.JSON(http.StatusOK, []any{}) }
}
func ListAchievementsHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) { c.JSON(http.StatusOK, []any{}) }
}