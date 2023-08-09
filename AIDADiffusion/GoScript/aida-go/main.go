package main

import (
	"image/color"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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

var clock *widget.Label
var imageUrl = "http://localhost:9998/files/testfile00000000.jpg"

func main() {
	println("Starting AIDA UI & staticserver...")
	go startstaticserver()

	myapp := app.New()
	mywindow := myapp.NewWindow("AIDADIFFUSION-GO")
	mywindow.SetFixedSize(false)
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 200}
	//blank := color.NRGBA{R: 0, G: 0, B: 0, A: 0}
	blank2 := color.Alpha{A: 0}
	emptyVerticalSpace := canvas.NewRectangle(blank2)
	emptyVerticalSpace.SetMinSize(fyne.NewSize(100, 10))
	emptyHorizontalSpace := canvas.NewRectangle(blank2)
	emptyHorizontalSpace.SetMinSize(fyne.NewSize(10, 26))

	titleText := canvas.NewText("AIDA", white)
	titleText.TextSize = 30
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle = fyne.TextStyle{Bold: true}
	var image *canvas.Image
	loadedImages, err := fyne.LoadResourceFromURLString(imageUrl)
	if err != nil {
		log.Println(err)
		image = canvas.NewImageFromFile("./res/fallback_IMG.png")

	} else {
		image = canvas.NewImageFromResource(loadedImages)
	}

	image.SetMinSize(fyne.NewSize(450, 450))
	image.FillMode = canvas.ImageFillContain

	promptInput := widget.NewEntry()
	promptInput.SetPlaceHolder("Enter Your Prompts Here...")
	promptInput.Resize(fyne.NewSize(400, 60))
	os.Chdir("./diffusers")
	generateBtn := widget.NewButton("GENERATE", func() { go generateImage(image, promptInput.Text) })
	refreshBtn := widget.NewButton("REFRESH IMG", func() { go refreshIMG(image) })

	titleBorder := container.NewBorder(emptyHorizontalSpace, emptyHorizontalSpace, nil, nil, titleText)
	imageBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, image)
	promptBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, promptInput)
	buttonGrid := container.NewHBox(generateBtn, refreshBtn)
	buttonBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, buttonGrid)
	centerButtonBorder := container.NewCenter(buttonBorder)

	contentGrid := container.New(
		layout.NewVBoxLayout(), titleBorder,
		imageBorder, promptBorder, centerButtonBorder)

	mywindow.SetContent(contentGrid)
	mywindow.ShowAndRun()

	mywindow.ShowAndRun()

}

func getImage(url string) (fyne.Resource, error) {
	var getURL = url
	loadedImages, err := fyne.LoadResourceFromURLString(getURL)
	if err != nil {
		log.Fatal(err)
		return loadedImages, err
	} else {
		return loadedImages, err
	}
}

func generateImage(image *canvas.Image, prompt string) {
	println("generating")
	testGenerateImage(prompt)
	go refreshIMG(image)
}

func refreshIMG(image *canvas.Image) {

	newImage, err := fyne.LoadResourceFromURLString(imageUrl)
	if err != nil {
		log.Println(err)
		image = canvas.NewImageFromFile("./res/fallback_IMG.png")
		image.Refresh()
	} else {
		image.Resource = newImage
		image.Refresh()
	}
}

func startstaticserver() {
	exec.Command("./staticserver.exe").Start()
	println("AIDA UI & staticserver is up and running!🎉")

}

func testGenerateImage(prompt string) {
	var testPrompt = prompt
	println(testPrompt)
	var itNum = "300"
	var saveNames = "testfile00000000"
	var width = "512"
	var height = "512"

	GenerateImage(testPrompt, itNum, saveNames, width, height)
	println("finish test generation")
}

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
