package info

//go:generate gofp -destination fp.go -pkg info -type "Info"
type Info struct {
	Title string
	Info  string
	Link  []string
}

func GetAllInfo() []Info {
	var allInfo []Info
	allInfo = append(allInfo, getMiscInfo()...)
	return allInfo
}
