package views

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/SphereSoftware/go4rails/data"
)

type View struct {
	project      *data.ProjectData
	viewFilePath string
	Toc          string
}

// New is factory method for views in a project
func New(project *data.ProjectData) *View {
	return &View{
		project:      project,
		viewFilePath: project.ProjectRootFolder + "/views/README.md",
	}
}

func walker(node *data.Node, level int) string {
	if !node.Info.IsDir {
		return ""
	}

	level++
	tmpPad := fmt.Sprintf(fmt.Sprintf("%%%ds", level*2), "- ")
	tocItem := tmpPad + fmt.Sprintf("[%s](%s)\n", node.Info.Name, node.LocalPath)

	for _, child := range node.Children {
		tocItem += walker(child, level)
	}

	return tocItem
}

func (v *View) BuildToc() {
	var toc string

	node := v.project.Root
	toc = walker(node, 0)

	v.Toc = toc
}

func (v *View) Build() {
	var err error
	rawTpl, _ := ioutil.ReadFile(v.viewFilePath)
	v.BuildToc()

	t := template.Must(template.New("readme").Parse(string(rawTpl)))
	// err = t.Execute(os.Stdout, v)
	// if err != nil {
	// 	log.Println("executing template:", err)
	// }

	f, err := os.OpenFile(v.project.ProjectRootFolder+"/README.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("can't open a file:", err)
	}

	err = t.Execute(f, v)
	if err != nil {
		log.Println("Generating README.md file", err)
	}
}
