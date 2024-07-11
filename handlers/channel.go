package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateChannel(c *gin.Context, db *sql.DB) {
	var channel Channel

	err := c.ShouldBindJSON(&channel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO channels (name) VALUES (?)", channel.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func ListChannels(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM channels")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var channels []Channel

	for rows.Next() {
		var channel Channel
		err := rows.Scan(&channel.ID, &channel.Name)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		channels = append(channels, channel)
	}

	c.JSON(http.StatusOK, channels)
}
