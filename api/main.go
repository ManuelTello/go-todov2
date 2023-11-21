package main

import (
	"fmt"
	"net/http"

	"todo.proj/src/handlers"
	"todo.proj/src/repo"
)

// Global
var workspaces *repo.Repository = &repo.Repository{}

func main() {
	/***************	Testing    ***************/
	/********************************************/

	http.Handle("/fetch", handlers.WorkspaceFetch(workspaces))
	http.Handle("/add", handlers.WorkspacesAdd(workspaces))
	http.Handle("/remove", handlers.WorkspaceRemove(workspaces))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
