package handler

type TaskRequest struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type TaskIDRequest struct {
	Id string `json:"Id"`
}
