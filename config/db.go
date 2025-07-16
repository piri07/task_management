package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "task_management/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&models.Task{})
    DB = db
}
