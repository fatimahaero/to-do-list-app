package main

import (
	"log"
	"net/http"
	conf "to-do-list-app/config"
	"to-do-list-app/handler"
	"to-do-list-app/repository"
	"to-do-list-app/service"
)

func main() {
	db := conf.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/register", userHandler.RegisterUser)
	http.HandleFunc("/login", userHandler.LoginUser)

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	http.HandleFunc("/tasks", taskHandler.GetAllTasks)
	http.HandleFunc("/tasks/create", taskHandler.CreateTask)
	http.HandleFunc("/tasks/delete/", taskHandler.DeleteTask)
	http.HandleFunc("/tasks/update/", taskHandler.UpdateTask)

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
