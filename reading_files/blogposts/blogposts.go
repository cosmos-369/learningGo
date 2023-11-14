package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {

	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := filesystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	textLine := readMetaLine(titleSeperator)
	descriptionLine := readMetaLine(descriptionSeperator)
	tagsLine := readMetaLine(tagsSeperator)

	scanner.Scan() //ignore --- line
	// scanner.Scan()
	// scanner.Scan()
	body := readBody(scanner)

	return Post{Title: textLine,
		Description: descriptionLine,
		Tags:        strings.Split(tagsLine, ", "),
		Body:        body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
