package main

import (
	"log"
	"net/http"
	"os/exec"
)

func focusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("     Focus request recived.")
	// change workspace
	exec.Command("i3-msg", "workspace", "1").Run()
	exec.Command("xdotool", "mousemove", "10", "500").Run()

	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/focus", focusHandler)

	log.Println("     Agent on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
