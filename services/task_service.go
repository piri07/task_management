package services

import (
    "task_management/config"
    "task_management/models"
)

func CreateTask(task *models.Task) error {
    return config.DB.Create(task).Error
}

func GetTasks(status string, page int, limit int) ([]models.Task, int64, error) {
    var tasks []models.Task
    var total int64

    query := config.DB.Model(&models.Task{})
    if status != "" {
        query = query.Where("status = ?", status)
    }

    query.Count(&total)
    err := query.Offset((page - 1) * limit).Limit(limit).Find(&tasks).Error
    return tasks, total, err
}

func GetTaskByID(id uint) (*models.Task, error) {
    var task models.Task
    result := config.DB.First(&task, id)
    return &task, result.Error
}

func UpdateTask(task *models.Task) error {
    return config.DB.Save(task).Error
}

func DeleteTask(id uint) error {
    return config.DB.Delete(&models.Task{}, id).Error
}
