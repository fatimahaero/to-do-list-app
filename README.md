# To-Do List App

This is a simple **To-Do List App** built using **Golang** with **Hexagonal Architecture** and **SQLite** as the database.

## Features

- User Registration & Login (with password encryption)
- Create, Update, Delete Checklists
- Create, Update, Delete To-Do Items inside a Checklist
- Update To-Do Item status (e.g., mark as completed)
- Soft delete functionality (tasks marked as deleted instead of actual deletion)

---

## 🔧 **System Design: Hexagonal Architecture**
This project follows **Hexagonal Architecture**, also known as **Ports and Adapters**. It consists of:

- `domain/` → Contains core business logic (entities & interfaces)
- `repository/` → Handles data persistence (SQLite)
- `service/` → Implements business rules & interacts with repositories
- `handler/` → Handles HTTP requests & responses
- `config/` → Database setup & configuration
- `main.go` → Entry point of the application

---

## 🚀 **How to Run the Project**

### **1. Clone the Repository**
```sh
git clone git@github.com:fatimahaero/to-do-list-app.git
cd to-do-list-app
```

### **2. Install Dependencies**
```sh
go mod tidy
```

### **3. Run the Application**
```sh
go run main.go
```

---
## 🛠 ** API Endpoints**

### **1. User Registration**
```sh
POST /register
```

```json
{
    "username": "testuser",
    "password": "securepassword"
}
```


### **2. Login**
```sh
POST /login
```

```json
{
    "username": "testuser",
    "password": "securepassword"
}
```