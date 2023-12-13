package blogposts_test

import (
	"errors"
	"io/fs"
	blogposts "main/reading_files/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

type StrubFailingFS struct {
}

func (s *StrubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("failed to open the file")
}

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	if err != nil {
		t.Fatal(err)
	}

	assertPosts(t, got, want)

	//test to check the error
	// t.Run("trying to open a failing file system", func(t *testing.T) {

	// 	failingFS := &StrubFailingFS{}

	// 	_, err := blogposts.NewPostsFromFS(failingFS)

	// 	if err == nil {
	// 		t.Fatal("expected an error but got none")
	// 	}
	// })
}

func assertPosts(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
