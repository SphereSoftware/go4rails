package data

// ProjectData is root object collects whole information about this project and
// used for TOC generation as well as HTML version build.
type ProjectData struct {
	Name string
}

// Preload walk through data dir and load all examples, parse Ruby & Go files
// Add Sections and Articles into ProjectData object for future use.
func Preload() ProjectData {
	project := ProjectData{
		Name: "Go4Rails",
	}

	return project
}
