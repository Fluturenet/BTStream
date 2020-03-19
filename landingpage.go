package main

import (
	"io"
	_ "log"
	"net/http"
)

func landingpage(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, landingpage_template)
}

var landingpage_template = `<html>
<head></head>
<body>
<h1>BTStream</h1>

<div>Icons made by <a href="https://www.flaticon.com/authors/freepik" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a></div>
</body>
</html>`
