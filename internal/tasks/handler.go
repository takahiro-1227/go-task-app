package tasks

import (
	"go-task-app/internal/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID        uint      `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(255); unique; not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func GetTasks(c *gin.Context) {
	var tasks []Task

	result := config.DB.Find(&tasks)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "タスクの取得に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func CreateTask(c *gin.Context) {
	var newTask Task

	err := c.ShouldBindJSON(&newTask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "入力値が不正です。",
		})
		return
	}

	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タスクのタイトルを入力してください。",
		})
		return
	}

	err = config.DB.Where("title = ?", newTask.Title).First(&Task{}).Error

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タスクのタイトルが重複しています。",
		})
		return
	}

	result := config.DB.Create(&newTask)

	if result.Error != nil {
		log.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タスクの作成に失敗しました。",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": newTask,
	})
}
