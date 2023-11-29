package repo

type TaskNode struct {
	Id   int       `json:"id"`
	Task string    `json:"task"`
	Next *TaskNode `json:"next"`
}

type BoardNode struct {
	Id       int        
	Name     string     
	TaskHead *TaskNode  
	Next     *BoardNode 
}

type WorkspaceNode struct {
	Id        int            
	Name      string         
	BoardHead *BoardNode     
	Next      *WorkspaceNode 
}

type WorkspaceDAO struct {
	Id        int            `json:"workspaceid"`
	Name      string         `json:"workspacename"`
	Next      *WorkspaceDAO `json:"next"`
}

type BoardDAO struct {
	Id       int        `json:"boardid"`
	Name     string     `json:"boardname"`
	TaskHead []string  `json:"tasks"`
	Next     *BoardDAO `json:"next"`
}

type Repository struct {
	Workspaces *WorkspaceNode
}

func (r Repository) AddTask(workspaceid int, boardid int, task string) {
	var aux_workspaces *WorkspaceNode = r.Workspaces

	for aux_workspaces.Id != workspaceid {
		aux_workspaces = aux_workspaces.Next
	}

	var aux_board *BoardNode = aux_workspaces.BoardHead

	for aux_board.Id != boardid {
		aux_board = aux_board.Next
	}

	if aux_board.TaskHead == nil {
		aux_board.TaskHead = &TaskNode{
			Id:   0,
			Task: task,
			Next: nil,
		}
	} else {
		var aux_task *TaskNode = aux_board.TaskHead

		for aux_task.Next != nil {
			aux_task = aux_task.Next
		}

		aux_task.Next = &TaskNode{
			Id:   aux_board.Id + 1,
			Task: task,
			Next: nil,
		}
	}
}

func (r Repository) RemoveTask(workspaceid int, boardid int, tasks []string) {
}

func (r *Repository) AddBoard(title string, workspaceid int) {
	var aux_workspace *WorkspaceNode = r.Workspaces

	for aux_workspace.Id != workspaceid {
		aux_workspace = aux_workspace.Next
	}

	if aux_workspace.BoardHead == nil {
		aux_workspace.BoardHead = &BoardNode{
			Id:       0,
			TaskHead: nil,
			Name:     title,
			Next:     nil,
		}
	} else {
		var aux_board *BoardNode = aux_workspace.BoardHead

		for aux_board.Next != nil {
			aux_board = aux_board.Next
		}

		aux_board.Next = &BoardNode{
			Id:       aux_board.Id + 1,
			TaskHead: nil,
			Name:     title,
			Next:     nil,
		}
	}
}

func (r Repository) RemoveBoard(workspaceid int, boardid int) {
	if r.Workspaces.Id == workspaceid {
		if r.Workspaces.BoardHead.Id == boardid {
			r.Workspaces.BoardHead = r.Workspaces.BoardHead.Next
		} else {
			var aux_board *BoardNode = r.Workspaces.BoardHead

			for aux_board.Next.Id != boardid {
				aux_board = aux_board.Next
			}

			aux_board.Next = aux_board.Next.Next
		}
	} else {
		var aux_workspace *WorkspaceNode = r.Workspaces

		for aux_workspace.Next != nil {
			aux_workspace = aux_workspace.Next
		}

		if aux_workspace.BoardHead.Id == boardid {
			aux_workspace.BoardHead = aux_workspace.BoardHead.Next
		} else {
			var aux_board *BoardNode = aux_workspace.BoardHead

			for aux_board.Next.Id != boardid {
				aux_board = aux_board.Next
			}

			aux_board.Next = aux_board.Next.Next
		}
	}
}

func (r *Repository) AddWorkspace(name string) {
	if r.Workspaces == nil {
		r.Workspaces = &WorkspaceNode{
			Id:        0,
			Name:      name,
			BoardHead: nil,
			Next:      nil,
		}
	} else {
		var aux_workspace *WorkspaceNode = r.Workspaces

		for aux_workspace.Next != nil {
			aux_workspace = aux_workspace.Next
		}

		aux_workspace.Next = &WorkspaceNode{
			Id:        aux_workspace.Id + 1,
			BoardHead: nil,
			Name:      name,
			Next:      nil,
		}
	}
}

func (r *Repository) RemoveWorkspace(workspaceid int) {
	if r.Workspaces.Id == workspaceid {
		r.Workspaces = r.Workspaces.Next
	} else {
		var aux_workpsace *WorkspaceNode = r.Workspaces

		for aux_workpsace.Next.Id == workspaceid {
			aux_workpsace = aux_workpsace.Next
		}

		aux_workpsace.Next = aux_workpsace.Next.Next
	}
}

func(r *Repository) FetchWorkspaces()*WorkspaceDAO{
	var workspaces_head *WorkspaceDAO = nil
	var aux_workpsaces *WorkspaceNode = r.Workspaces

	for aux_workpsaces != nil{
		if workspaces_head == nil {
			workspaces_head = &WorkspaceDAO{
				Id: aux_workpsaces.Id,
				Name: aux_workpsaces.Name,
				Next: nil,
			}
		}else{
			var aux_return *WorkspaceDAO = workspaces_head

			for aux_return.Next != nil {
				aux_return = aux_return.Next
			}

			aux_return.Next = &WorkspaceDAO{
				Id: aux_workpsaces.Id,
				Name: aux_return.Name,
				Next: nil,
			}
		}

		aux_workpsaces = aux_workpsaces.Next
	}

	return workspaces_head
}

func(r *Repository) FetchBoard(boardid int)*BoardNode{
	var board_head *BoardNode = nil

	return board_head
}