package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Fancy Adventure Go App",
	})
}

func Adventure(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		name = "Stranger"
	}

	adventures := []string{
		"🐉 fought a dragon",
		"💰 found treasure",
		"⛵ sailed the seven seas",
		"🕳️ explored a mysterious cave",
		"🧙 befriended a wizard",
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	adventure := adventures[rng.Intn(len(adventures))]

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Fancy Adventure Go App",
		"name":  name,
		"story": adventure,
	})
}
