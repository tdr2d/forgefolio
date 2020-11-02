package admin

type HeadData struct {
	title  string
	styles []string
}

var DEFAULT_TITLE string = "My default title"
var DEFAULT_STYLE string = "/style.css"

func GetHeadData() *HeadData {
	h := HeadData{title: DEFAULT_TITLE, styles: []string{DEFAULT_STYLE}}
	return &h
}
