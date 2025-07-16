package controllers

import (
    "net/http"
    "strconv"
    "time"
    "task_management/models"
    "task_management/services"
    "github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task.CreatedAt = time.Now()
    if err := services.CreateTask(&task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    status := c.Query("status")

    tasks, total, err := services.GetTasks(status, page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  tasks,
        "page":  page,
        "limit": limit,
        "total": total,
    })
}

func GetTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    task, err := services.GetTaskByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    task, err := services.GetTaskByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    var input models.Task
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task.Title = input.Title
    task.Description = input.Description
    task.Status = input.Status

    if err := services.UpdateTask(task); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    err := services.DeleteTask(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
