package models

type University struct {
    Name    string `json:"name"`
    Address string `json:"address"`
    Country string `json:"country"`
}

type Major struct {
    ID          int         `json:"id"`
    Name        string      `json:"name"`
    Description string      `json:"description"`
    University  University  `json:"university"`
	SubjectDescription []string `json:"subjects"`
}