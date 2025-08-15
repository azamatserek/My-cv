package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("DATABASE_URL is required")
    }
    db, err := sql.Open("postgres", dsn)
    if err != nil { log.Fatal(err) }
    if err := db.Ping(); err != nil { log.Fatal(err) }

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

    // Public content
    r.GET("/api/profile", GetProfileHandler(db))
    r.GET("/api/experience", ListExperienceHandler(db))
    r.GET("/api/education", ListEducationHandler(db))
    r.GET("/api/publications", ListPublicationsHandler(db))
    r.GET("/api/achievements", ListAchievementsHandler(db))

    // Optional: simple basic-auth for admin
    admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
        os.Getenv("ADMIN_USER"): os.Getenv("ADMIN_PASS"),
    }))
    admin.POST("/rebuild-cache", func(c *gin.Context){ c.JSON(200, gin.H{"status":"ok"}) })

    addr := ":8080"
    if p := os.Getenv("PORT"); p != "" { addr = ":" + p }
    log.Printf("listening on %s", addr)
    r.Run(addr)
}