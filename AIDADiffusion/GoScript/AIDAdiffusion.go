package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type PromtData struct {
	PromtString     string
	PromtIterations string
	PromtSaveName   string
}

func main() {
	//start browser
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "http://localhost:9998/files/build/").Start()
	println("Starting AIDA UI & staticserver...")
	//start static file server
	startstaticserver()
	//start AIDAdiffuser server
	println("Starting AIDA server..")
	//change working directory to examples/inference, required to run the stable diffusion model
	os.Chdir("./diffusers/examples/inference")
	startServer()
}

func startstaticserver() {
	exec.Command("staticserver.exe").Start()
	println("AIDA UI & staticserver is up and running!🎉")

}

func startServer() {
	//API for Inferencing
	http.HandleFunc("/test", PromptAPI)
	println("AIDA server up and running!🎉")
	//listen and serve
	http.ListenAndServe(":9999", nil)
}

func PromptAPI(w http.ResponseWriter, r *http.Request) {
	var pD PromtData
	err := json.NewDecoder(r.Body).Decode(&pD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("PromptData: %+v", pD)
	GenerateImage(pD.PromtString, pD.PromtIterations, pD.PromtSaveName)

}

// func APIHandleInference(w http.ResponseWriter, r *http.Request) {
// 	//start Inferencing model
// 	InferenceModel()
// }

// func InferenceModel() {
// 	//run the ML script
// 	println("Inferencing.... Please wait.")
// 	exec.Command("python", "scripttest.py", "a dog driving a super fast cyberpunk flying car while scrolling twitter", "1", "dogdrivingcar").Run()
// }

func GenerateImage(promtInput string, iterationNums string, saveName string) {
	//execute InferenceModel() when api triggered
	println("Inferencing.... Please wait.")
	exec.Command("python", "scripttest.py", promtInput, iterationNums, saveName).Run()
	println("Inference completed!🎉")
}
