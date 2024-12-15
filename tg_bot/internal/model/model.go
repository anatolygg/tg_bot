package model

type MLRequest struct {
	Question string `yaml:"question"`
}

type MLResponse struct {
	Answer string `yaml:"answer"`
}
