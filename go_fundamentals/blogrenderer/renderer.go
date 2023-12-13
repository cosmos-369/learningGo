package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ  *template.Template
	parser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	tmp, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return &PostRenderer{}, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: tmp, parser: p}, nil
}

func (r *PostRenderer) Render(w io.Writer, p PostViewModel) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {

	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type PostViewModel struct {
	Post
	HTMLBody template.HTML
}

func NewPostVM(p Post, r *PostRenderer) PostViewModel {
	doc := r.parser.Parse([]byte(p.Body))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	Bodyhtml := markdown.Render(doc, renderer)

	return PostViewModel{p, template.HTML(Bodyhtml)}
}
