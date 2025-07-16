# 🧩 Task Management Microservice (Go + Gin + GORM)

A simple Task Management System built as a microservice using Go, Gin, and GORM (with SQLite). It supports full CRUD operations, pagination, and filtering by status.

---

## 🧠 Problem Breakdown & Design Decisions

### 🎯 Objective
Build a lightweight service to manage tasks with:
- CRUD operations
- Filtering by `status`
- Pagination for list endpoint

### ⚒️ Tech Choices
- **Go (Golang)**: Chosen for its performance and simplicity for microservices.
- **Gin**: Lightweight HTTP web framework for routing and middleware support.
- **GORM**: ORM to manage data access without writing raw SQL.
- **SQLite**: Easy local persistence for development.

### 📦 Microservices Principles Applied
- **Single Responsibility**: This service handles only task-related operations.
- **Stateless**: No session or in-memory state; all data is persisted in SQLite.
- **Well-defined API**: RESTful endpoints with clear contracts.
- **Separation of Concerns**: Clean architecture with controller, service, and model layers.
- **Scalable Foundation**: Easy to extract into a container and deploy as a service in a larger architecture.

---

## 🚀 Getting Started

### ✅ Prerequisites

- Go 1.18 or later
- Git

### 📦 Install Dependencies

```bash
go mod tidy
go run main.go
