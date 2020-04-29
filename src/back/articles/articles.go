package articles

import (
)


type Content struct {
	Data
}

type  Reference interface {
	Name() string
}

// Contents type represents articles tree
type Contents struct {
	Name string `json:"Name"`
	Articles []Article `json:"Articles"`
	Children []Contents `json:"Children"`
}

// Article type represent single article data
type Article struct {
	Id int			`json:"Id"`
	Name string 	`json:"Name"`
	Data *string 	`json:"Data"`
}

func ScanFiles(path string) {

}
