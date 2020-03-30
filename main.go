package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Eccomi '%s'\n", r.URL.Path[1:])
	numbers := []int{0,1,2,3,4,5,6,7,8}
	for c := range numbers {
		fmt.Fprintf(w, "     %d\n", c)
	}
}

func handlerHome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Titolo", "Body")
	fmt.Fprintf(w, "<a href='%s'>%s</a>", "/link", "Link")
}

func main() {
	go open("http://localhost:8080/")
	http.HandleFunc("/", handlerHome)
	http.HandleFunc("/link", handler)
	panic(http.ListenAndServe(":8080", nil))
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

