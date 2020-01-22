package rest

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ErrorRest struct {
	Msg string `json:"msg"`
}

type TasksGetAll struct {
	Tasks []TaskResponse `json:"tasks"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}
