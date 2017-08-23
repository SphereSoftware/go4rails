package data

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ProjectData is root object collects whole information about this project and
// used for TOC generation as well as HTML version build.
type ProjectData struct {
	Name              string // Project title
	Root              *Node  // Root fot parsed data
	ProjectRootFolder string // path to data folder
}

// FileInfo is a struct created from os.FileInfo interface for serialization.
type FileInfo struct {
	Name    string      `json:"name"`
	Size    int64       `json:"size"`
	Mode    os.FileMode `json:"mode"`
	ModTime time.Time   `json:"mod_time"`
	IsDir   bool        `json:"is_dir"`
}

// Node represents a node in a directory tree.
type Node struct {
	FullPath  string    `json:"path"`
	LocalPath string    `json:"local_path"`
	Info      *FileInfo `json:"info"`
	Children  []*Node   `json:"children"`
	IsDir     bool      `json:"is_dir"`
	Parent    *Node     `json:"-"`
}

// New is basic factory that build ProjectData object
func New() ProjectData {
	project := ProjectData{
		Name:              "Go4Rails",
		ProjectRootFolder: projectRrootDir(),
	}

	return project
}

func projectRrootDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// Helper function to create a local FileInfo struct from os.FileInfo interface.
func fileInfoFromInterface(v os.FileInfo) *FileInfo {
	return &FileInfo{v.Name(), v.Size(), v.Mode(), v.ModTime(), v.IsDir()}
}

// NewTree Create directory hierarchy tree.
func newTree(root string) *Node {
	parents := make(map[string]*Node)

	// walker function
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		parents[path] = &Node{
			FullPath:  path,
			LocalPath: strings.TrimLeft(path, root),
			Info:      fileInfoFromInterface(info),
			Children:  make([]*Node, 0),
			IsDir:     info.IsDir(),
		}

		return nil
	}

	filepath.Walk(root+"/data", walkFunc)

	var result *Node
	for path, node := range parents {

		parentPath := filepath.Dir(path)
		parent, exists := parents[parentPath]

		// If a parent does not exist, this is the root.
		if !exists {
			result = node
		} else {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
		}

	}

	return result
}

// LoadData will go through data dir and load all examples, parse Ruby & Go files
// Add Sections and Articles into ProjectData object for future use.
func (p *ProjectData) LoadData() {
	p.Root = newTree(p.ProjectRootFolder)
}
