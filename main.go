// Package main is a collections of tips for Ruby on Rails developer
// how to solve day to day problems using `Go`
package main

import (
	"github.com/SphereSoftware/go4rails/data"
	"github.com/SphereSoftware/go4rails/views"
)

func main() {
	project := data.New()
	project.LoadData()
	readme := views.New(&project)
	readme.Build()
}
