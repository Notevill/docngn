package fsarticles

import (
	"crypto/md5"
	"io/ioutil"
	"os"

	"github.com/Notevill/docngn/back/articles"
)

type contents struct {
	N  string
	D  string
	Ch []contents
}

func (cnts contents) Name() string {
	return cnts.N
}
func (cnts contents) DataID() string {
	return cnts.D
}
func (cnts contents) Children() []articles.Contents {
	children := make([]articles.Contents, len(cnts.Ch))
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

//ScanArticles scan articles on file system path
func ScanArticles(path string) (articles.Articles, articles.Contents, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	arts := make(articles.Articles)
	root := contents{
		N: "root",
		D: "",
	}
	for _, file := range files {

		if !file.IsDir() {
			// read regular file if it's md file
			// and save it to articles buffer
			if file.Name()[len(file.Name())-3:] == ".md" {
				data, err := ioutil.ReadFile(path + string(os.PathSeparator) + file.Name())
				if err != nil {
					return nil, nil, err
				}
				sum := md5.Sum(data)
				articleContent := article{
					ID: string(sum[:]),
					D:  string(data),
				}
				arts[articleContent.ID] = articleContent
				root.Ch = append(root.Ch, contents{
					N: file.Name(),
					D: articleContent.ID,
				})
			}
		} else {
			// recursively scan all dirrectories
			a, cnts, err := ScanArticles(path + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, nil, err
			}
			// append two maps
			for k, v := range a {
				arts[k] = v
			}
			root.Ch = append(root.Ch, cnts.(contents))
		}
	}
	return arts, root, nil
}
