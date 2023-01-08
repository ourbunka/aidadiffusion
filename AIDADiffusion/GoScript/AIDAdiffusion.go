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
	PromptWidth     string
	PromptHeight    string
}

type advancePromtData struct {
	PromtString        string
	PromtIterations    string
	PromtSaveName      string
	PromtSeedInput     string
	PromtUserScheduler string
	PromtUserGuidance  string
	PromptWidth        string
	PromptHeight       string
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
	os.Chdir("./diffusers")
	startServer()
}

func startstaticserver() {
	exec.Command("staticserver.exe").Start()
	println("AIDA UI & staticserver is up and running!🎉")

}

func startServer() {
	//API for Inferencing
	http.HandleFunc("/standard", PromptAPI)
	http.HandleFunc("/advance", PromptAPIAdvance)
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
	fmt.Println("PromptData: %+v", pD, err)
	GenerateImage(pD.PromtString, pD.PromtIterations, pD.PromtSaveName, pD.PromptWidth, pD.PromptHeight)

}

func PromptAPIAdvance(w http.ResponseWriter, r *http.Request) {
	var apD advancePromtData
	err := json.NewDecoder(r.Body).Decode(&apD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("PromptData: %+v", apD, err)
	GenerateImageAdvance(apD.PromtString, apD.PromtIterations, apD.PromtSaveName, apD.PromtSeedInput, apD.PromtUserScheduler, apD.PromtUserGuidance, apD.PromptWidth, apD.PromptHeight)

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

func GenerateImage(promtInput string, iterationNums string, saveName string, inputWidth string, inputHeight string) {
	//execute InferenceModel() when api triggered
	println("Inferencing.... Please wait.")
	exec.Command("python", "script_cuda.py", promtInput, iterationNums, saveName, inputWidth, inputHeight).Run()
	println("Inference completed!🎉")
}

func GenerateImageAdvance(promtInput string, iterationNums string, saveName string, userSeedInput string, PromtUserScheduler string, PromtUserGuidance string, inputWidth string, inputHeight string) {
	//execute InferenceModel() when api triggered
	println("Inferencing.... Please wait.")
	exec.Command("python", "script_cudaadv.py", promtInput, iterationNums, saveName, userSeedInput, PromtUserScheduler, PromtUserGuidance, inputWidth, inputHeight).Run()
	println("Inference completed!🎉")
}
