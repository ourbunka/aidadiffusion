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
	//change working directory to examples/inference, required to run the stable diffusion model
	os.Chdir("./diffusers/examples/inference")
	println("For the app to run properly, please run staticserver.exe in the same directory, then go to ", "http://localhost:9998/files/build/", "in your webbrowser to start using the AIDAdiffusion app.", "Optionally, use https://ourbunka.github.io/aidadiffusion for the GUI.")
	startServer()

}

func startServer() {

	//create endpoint for serving react frontend
	http.HandleFunc("/home", APIHandleHome)
	//API for Inferencing
	//http.HandleFunc("/inference", APIHandleInference)
	http.HandleFunc("/test", testPromt)
	//listen and serve
	http.ListenAndServe(":9999", nil)
}

func APIHandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, welcome to the Homepage, Inferencing API is Up and running!🎉,make sure you start staticserver.exe too.")
	//add server for serving react front end

}

func testPromt(w http.ResponseWriter, r *http.Request) {
	var pD PromtData
	err := json.NewDecoder(r.Body).Decode(&pD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("PromtData: %+v", pD)
	GenerateImage(pD.PromtString, pD.PromtIterations, pD.PromtSaveName)

}

func APIHandleInference(w http.ResponseWriter, r *http.Request) {
	//start Inferencing model
	InferenceModel()
}

func InferenceModel() {
	//run the ML script
	println("Inferencing.... Please wait.")
	exec.Command("python", "scripttest.py", "a dog driving a super fast cyberpunk flying car while scrolling twitter", "1", "dogdrivingcar").Run()
}

func GenerateImage(promtInput string, iterationNums string, saveName string) {
	//execute InferenceModel() when api triggered
	println("Inferencing.... Please wait.")
	exec.Command("python", "scripttest.py", promtInput, iterationNums, saveName).Run()
	println("Inference completed!🎉")
}
