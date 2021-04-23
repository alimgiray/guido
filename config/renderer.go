package config

import (
	"embed"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

func LocalRenderer() multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles(
		"login",
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/user/login.html")
	renderer.AddFromFiles(
		"register",
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/user/register.html")
	renderer.AddFromFiles(
		"topic",
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/topic/topic.html", "ui/topic/post.html", "ui/topic/add.html")
	renderer.AddFromFiles(
		"create",
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/topic/create.html")
	return renderer
}

// File order is important!
func EmbedRenderer(f embed.FS) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	loginTemplate, err := template.ParseFS(f,
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/user/login.html")
	if err != nil {
		panic(err)
	}
	renderer.Add("login", loginTemplate)

	registerTemplate, err := template.ParseFS(f,
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/user/register.html")
	if err != nil {
		panic(err)
	}
	renderer.Add("register", registerTemplate)

	topicTemplate, err := template.ParseFS(f,
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/topic/topic.html", "ui/topic/post.html", "ui/topic/add.html")
	if err != nil {
		panic(err)
	}
	renderer.Add("topic", topicTemplate)

	createTemplate, err := template.ParseFS(f,
		"ui/layout/index.html", "ui/layout/header.html", "ui/layout/nav.html", "ui/layout/meta.html",
		"ui/topic/create.html")
	if err != nil {
		panic(err)
	}
	renderer.Add("create", createTemplate)

	return renderer
}
