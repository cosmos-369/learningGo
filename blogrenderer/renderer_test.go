package blogrenderer_test

import (
	"bytes"
	"io"
	"main/blogrenderer"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	newPost := blogrenderer.NewPostVM(aPost, postRenderer)

	t.Run("it converts a single post into an HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, newPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("it renders a index of posts", func(t *testing.T) {
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		buf := bytes.Buffer{}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(t *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	newPost := blogrenderer.NewPostVM(aPost, postRenderer)

	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		postRenderer.Render(io.Discard, newPost)
	}
}
