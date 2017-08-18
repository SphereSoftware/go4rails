package main

import (
	"fmt"

	"github.com/theplant/blackfriday"
)

func main() {
	// For basic usage, it is as simple as getting your input into a byte slice
	input := []byte(`
# Test H1
and some text here

## Test
Other text

- First
  - Second
`)

	// and calling MarkdownBasic
	output := blackfriday.MarkdownBasic(input)
	fmt.Println(string(output))
	// This renders it with no extensions enabled.
	// As result you will have following html
	// -------------------------------------
	// <h1>Test H1</h1>
	// <p>and some text here</p>
	// <h2>Test</h2>
	// <p>Other text</p>
	// <ul>
	// 	<li>First
	// 		<ul>
	// 			<li>Second</li>
	// 		</ul>
	// 	</li>
	// </ul>
	// -------------------------------------
}
