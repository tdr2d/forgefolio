package admin

type Nav struct {
	name string
	link string
	logo string // feather logo
}

// TODO load from database
var Navigation []Nav = []Nav{
	{name: "Home", link: "/", logo: "home"},
	{name: "Media", link: "/medias", logo: "image"},
	{name: "Settings", link: "/settings", logo: "settings"},
	{name: "Blog Posts", link: "/blog-posts", logo: "file-text"},
}
