package admin

// DataDir represents struct of variables used for storing data
type dataDir struct {
	Blog  string
	Page  string
	Theme string
}

var DataDir dataDir = dataDir{
	Blog:  "assets/blogdata",
	Page:  "assets/pagedata",
	Theme: "assets/themedata",
}

type nav struct {
	name  string
	link  string
	logo  string // feather logo
	class string // css class
}

// Constant represents struct of variables shared in jet templates
type Constant struct {
	MediaDir          string
	MediaThumbnailDir string
	BodyLimit         int
	Navigation        []nav
}

var Constants Constant = Constant{
	MediaDir:          "assets/media",
	MediaThumbnailDir: "assets/thumbnail",
	Navigation: []nav{
		{name: "Home", link: "/admin", logo: "home"},
		{name: "Media", link: "/admin/medias", logo: "image"},
		{name: "Blog Posts", link: "/admin/blog-posts", logo: "file-text"},
		{name: "Settings", link: "/admin/settings", logo: "settings"},
		{name: ""},
		{name: "Theme", link: "/admin/theme", logo: "layout", class: "multi-color"},
		{name: "Pages", link: "/admin/pages", logo: "layers", class: "desert-color"},
		{name: ""},
	},
}
