package main

import (
	"bytes"
	"github.com/xyproto/splash"
	"io/ioutil"
	"time"
)

var styles = []string{
	"abap",
	"algol",
	"algol_nu",
	"api",
	"arduino",
	"autumn",
	"borland",
	"bw",
	"colorful",
	"dracula",
	"emacs",
	"friendly",
	"fruity",
	"github",
	"igor",
	"lovelace",
	"manni",
	"monokai",
	"monokailight",
	"murphy",
	"native",
	"paraiso-dark",
	"paraiso-light",
	"pastie",
	"perldoc",
	"pygments",
	"rainbow_dash",
	"rrt",
	"solarized-dark",
	"solarized-dark256",
	"solarized-light",
	"swapoff",
	"tango",
	"trac",
	"vim",
	"vs",
	"xcode",
}

const (
	title     = "Chroma Style Gallery"
	simpleCSS = "body { font-family: sans-serif; margin: 4em; } .chroma { padding: 1em; }"
)

func main() {

	for i, styleName := range styles {

		// Generate a HTML document for the current style name
		var inputBuffer bytes.Buffer
		inputBuffer.WriteString("<!doctype html><html><head><title>")
		inputBuffer.WriteString(styleName)
		inputBuffer.WriteString("</title><style>")
		inputBuffer.WriteString(simpleCSS)
		inputBuffer.WriteString("</style></head><body><h1>")
		inputBuffer.WriteString(styleName)
		inputBuffer.WriteString("</h1><code><pre>")
		inputBuffer.WriteString(`
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`)
		inputBuffer.WriteString("</pre></code>")

		// Button to the previous style, if possible
		if i > 0 {
			prevName := styles[i-1]
			inputBuffer.WriteString("<button onClick=\"location.href='" + prevName + ".html'\">Prev</button>")
		} else {
			inputBuffer.WriteString("<button disabled='true'>Prev</button>")
		}

		// Button to the next style, if possible
		if i < (len(styles) - 1) {
			nextName := styles[i+1]
			inputBuffer.WriteString("<button onClick=\"location.href='" + nextName + ".html'\">Next</button>")
		} else {
			inputBuffer.WriteString("<button disabled='true'>Next</button>")
		}

		// Button to the overview
		inputBuffer.WriteString("<button onClick=\"location.href='index.html'\">Up</button>")

		inputBuffer.WriteString("</body></html>")

		// Highlight the source code in the HTML with the current style
		htmlBytes, err := splash.Splash(inputBuffer.Bytes(), styleName)
		if err != nil {
			panic(err)
		}

		// Write the HTML sample for this style name
		err = ioutil.WriteFile(styleName+".html", htmlBytes, 0644)
		if err != nil {
			panic(err)
		}

	}

	// Generate an Index file for viewing the different styles
	var buf bytes.Buffer
	buf.WriteString("<!doctype html><html><head><title>")
	buf.WriteString(title)
	buf.WriteString("</title><style>")
	buf.WriteString(simpleCSS)
	buf.WriteString("</style></head><body><h1>")
	buf.WriteString(title)
	buf.WriteString("</h1><ul>")
	for _, styleName := range styles {
		buf.WriteString("<li><a href=\"" + styleName + ".html\">" + styleName + "</a></li>")
	}
	buf.WriteString("</ul><small>Generated ")
	buf.WriteString(time.Now().UTC().Format(time.RFC3339)[:10])
	buf.WriteString(" by <a href=\"https://github.com/xyproto\">xyproto</a> using <a href=\"https://github.com/xyproto/splash\">splash</a>.")
	buf.WriteString("</small></body></html>")
	err := ioutil.WriteFile("index.html", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

}
