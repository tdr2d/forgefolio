package admin

// Nav nav
type Nav struct {
	name string
	link string
	logo string // feather logo
}

// Constant represent constant variables used in jet templates
type Constant struct {
	MediaDir          string
	MediaThumbnailDir string
	BlogDataDir       string
	PageDataDir       string
	BodyLimit         int
	Navigation        []Nav
}

var Constants Constant = Constant{
	MediaDir:          "assets/media",
	MediaThumbnailDir: "assets/thumbnail",
	BlogDataDir:       "assets/blogdata",
	PageDataDir:       "assets/pagedata",
	BodyLimit:         4 * 1024 * 1024,
	Navigation: []Nav{
		{name: "Home", link: "/admin", logo: "home"},
		{name: "Media", link: "/admin/medias", logo: "image"},
		{name: "Settings", link: "/admin/settings", logo: "settings"},
		{name: ""},
		{name: "Blog Posts", link: "/admin/blog-posts", logo: "file-text"},
		{name: "Pages", link: "/admin/pages", logo: "layout"},
	},
}
