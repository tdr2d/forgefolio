package admin

type HeaderData struct {
	title  string
	styles []string
}

var DEFAULT_TITLE string = "My default title"
var DEFAULT_STYLE string = "/style.css"

func GetHeaderData() *HeaderData {
	h := HeaderData{title: DEFAULT_TITLE, styles: []string{DEFAULT_STYLE}}
	return &h
}
