package tools

import "html/template"

// `templates` is a global variable that holds parsed HTML templates
// These templates are parsed using the `template.ParseFiles` function, which loads
// the template files into memory. The `Must` function ensures that if there is an error
// during template parsing, the program will panic, making it easier to detect errors at startup.

// The `ParseFiles` method loads the following templates:
// - "templates/index.html" (The homepage or main form for generating ASCII art)
// - "templates/result.html" (The page that displays the generated ASCII art)
// - "templates/errors.html" (The error page displayed if any issues occur during processing)

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/result.html", "templates/errors.html"))
