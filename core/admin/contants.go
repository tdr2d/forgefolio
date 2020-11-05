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

// MediaDir directory of medias
const MediaDir string = "assets/media"

// MediaThumbnailDir directory of media thumbnails
const MediaThumbnailDir string = "assets/thumbnail"

// BodyLimit max upload size
const BodyLimit int = 4 * 1024 * 1024
