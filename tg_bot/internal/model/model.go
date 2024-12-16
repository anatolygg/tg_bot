package model

type MLRequest struct {
	Question string `json:"question"`
}

type MLResponse struct {
	Answer string `json:"answer"`
}
