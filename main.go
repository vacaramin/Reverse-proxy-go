package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	website             = "https://beta.frase.io"
	initialName         = "beta.frase.io"
	finalName           = "spyfu.host"
	authorization_Token = "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJzZW1pbmFyc2VvQG91dGxvb2suY29tIiwiZXhwIjoxNzA4MzQ0Mjc5fQ.gfktiSmnouz_opQSmfH0mhSi1PFvidT6-VsKDL15TCwKcN-TlWDMPPw0h7kyTa5XrKrKfoxlCCmTW95UWNwq2w"
	TimeoutURL          = "https://spyfu.host/app/documents"
)

func main() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/login", loginhandler)

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
	}

}

func homehandler(w http.ResponseWriter, r *http.Request) {
	//	for Handling CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	agent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"
	uri := "/app/documents"

	if r.URL.Path == "/" {
		http.Redirect(w, r, uri, http.StatusFound)
		return
	}

	targetURL := website + r.URL.Path

	headers := make(http.Header)
	for name, values := range r.Header {
		if inArray(name, []string{"accept", "accept-language", "x-requested-with", "main-request", "x-newrelic-id", "x-xsrf-token", "cache-control", "content-type", "content-length", "authorization", "x-access-token", "x-human-token", "x-csrf-token", "x-requested-with"}) {
			headers[name] = values
		}
	}

	cookie := r.Header.Get("Cookie")
	headers.Add("Cookie", cookie)
	headers.Add("Origin", website)

	client := http.Client{}
	req, err := http.NewRequest(r.Method, targetURL, bytes.NewBuffer([]byte{}))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = headers
	req.Header.Set("User-Agent", agent)

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Modify the response body
	bodyString := string(body)
	bodyString = strings.ReplaceAll(bodyString, "</body>", "</body>replaceCode")
	bodyString = strings.ReplaceAll(bodyString, initialName, finalName)
	bodyString = strings.ReplaceAll(bodyString, "</head>", "<style>.pt-3{display:none!important;}</style></head>")

	// Set the Content-Type header
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// Write the modified response body to the client
	w.Write([]byte(bodyString))
}

func inArray(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	//	for Handling CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set the Content-Type header to text/html
	w.Header().Set("Content-Type", "text/html")

	// Write the HTML response directly to the client
	fmt.Fprintf(w, `<h1>Login Successful, <a href="#" onclick="window.open('/ryte','_blank');">click here</a> to continue</h1>
<script>
window.localStorage.setItem(
 "frase_token",
 '"%v"'
);

window.localStorage.setItem(
	"frase_refreshToken",
	'"61bcfa92-5d2d-489e-b12e-69d87cb5332f"'
);

window.localStorage.setItem(
	"lastVisitedPath",
	'/app/documents'
);

window.localStorage.setItem(
	"isDarkMode",
	'false'
);

sessionStorage.setItem(
	"sentryReplaySession",
	'{"id":"0ab4bc60fb2b46fe8a9173d734fe6892","started":1708260850132,"lastActivity":1708260853555,"segmentId":0,"sampled":"buffer"}'
);

setTimeout(() => location.href = '%s', 100);
</script>`, authorization_Token, TimeoutURL)
}
