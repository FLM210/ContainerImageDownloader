package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var image, filename string

func main() {
	LISTENPORT := os.Getenv("LISTENPORT")
	http.HandleFunc("/", handleForm)
	http.HandleFunc("/download", handleDownload)
	if LISTENPORT != "" {
		log.Fatal(http.ListenAndServe(":"+LISTENPORT, nil))
	} else {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}

}

func handleForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `
			<html>
				<body>
					<form action="/" method="post">
						<label for="image">Enter imageName:</label>
						<input type="text" id="image" name="image">
						<input type="submit" value="Download">
					</form>
				</body>
			</html>
		`)
	} else if r.Method == "POST" {
		arr := strings.Split(r.FormValue("image"), "/")
		filename = strings.Replace(arr[len(arr)-1], ":", "-", 1) + ".tar"
		err := os.WriteFile(filename, []byte("000"), 0644)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		image = r.FormValue("image")
		http.Redirect(w, r, "/download?filename="+filename, http.StatusSeeOther)
	}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	// cmd := exec.Command("skopeo", "copy", "docker://"+image, "oci-archive:"+filename)
	// fmt.Println(cmd)
	Command("skopeo copy docker://" + image + " oci-archive:" + filename)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/tar")
	http.ServeFile(w, r, filename)
}

func Command(cmd string) error {
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			fmt.Print(readString)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}
