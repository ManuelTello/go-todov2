package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"todo.proj/src/helpers"
	"todo.proj/src/lib"
	"todo.proj/src/repo"
)

func WorkspacesAdd(repository *repo.Repository) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var body helpers.Body = lib.MapBody(r.Body)

			switch body.Context["target"] {
			case "task":
				var todo string = body.Context["todo"]
				workspaceid, err_work := strconv.Atoi(body.Context["workspaceid"])
				boardid, err_board := strconv.Atoi(body.Context["boardid"])

				if err_work != nil || err_board != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					repository.AddTask(workspaceid, boardid, todo)
					w.WriteHeader(http.StatusOK)
				}
			case "board":
				var board_name string = body.Context["name"]
				workspace_id, err := strconv.Atoi(body.Context["workspaceid"])

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					repository.AddBoard(board_name, workspace_id)
					w.WriteHeader(http.StatusOK)
				}
			case "workspace":
				var workspace_name string = body.Context["name"]

				if repository.Workspaces == nil {
					repository.Workspaces = &repo.WorkspaceNode{
						Id:   0,
						Name: workspace_name,
					}
				} else {
					repository.AddWorkspace(workspace_name)
					w.WriteHeader(http.StatusOK)
				}
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
		}
	}

	return http.HandlerFunc(handler)
}

func WorkspaceFetch(repository *repo.Repository) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data, err := json.Marshal(repository.Workspaces)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Header().Add("Content-type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(data)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
		}
	}

	return http.HandlerFunc(handler)
}

func WorkspaceRemove(repository *repo.Repository) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			var body helpers.Body = lib.MapBody(r.Body)

			switch body.Context["target"] {
			case "task":
				workspaceid, err_work := strconv.Atoi(body.Context["workspaceid"])
				boardid, err_board := strconv.Atoi(body.Context["boardid"])

				if err_board != nil || err_work != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					var tasks []string = strings.Split(body.Context["tasks"], "&_&")
					repository.RemoveTask(workspaceid, boardid, tasks)
				}
			case "board":
				workspaceid, errwork := strconv.Atoi(body.Context["workspaceid"])
				boardid, errboard := strconv.Atoi(body.Context["boardid"])

				if errboard != nil || errwork != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					repository.RemoveBoard(workspaceid, boardid)
					w.WriteHeader(http.StatusOK)
				}
			case "workspace":
				workspaceid, err := strconv.Atoi(body.Context["workspaceid"])

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					repository.RemoveWorkspace(workspaceid)
					w.WriteHeader(http.StatusOK)
				}
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 not found"))
		}
	}

	return http.HandlerFunc(handler)
}
