package articles

// Contents - content reference table
type Contents interface {
	Name() string
	DataID() string
	Children() []Contents
}

// Article - article data
type Article interface {
	Data() string
}

// Articles set of articles stored by ID
type Articles map[string]Article

type contents struct {
	N  string     `json:"name"`
	D  string     `json:"dataId"`
	Ch []contents `json:"contents"`
}

func (cnts contents) Name() string {
	return cnts.N
}
func (cnts contents) DataID() string {
	return cnts.D
}
func (cnts contents) Children() []Contents {
	children := make([]Contents, len(cnts.Ch))
	for i, e := range cnts.Ch {
		children[i] = e
	}
	return children
}

type article struct {
	ID string `json:"id"`
	D  string `json:"data"`
}

func (art article) Data() string {
	return art.D
}
