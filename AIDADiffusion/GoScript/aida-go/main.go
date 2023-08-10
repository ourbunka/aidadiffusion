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
var imageUrl string = "http://localhost:9998/files/NewFileName_001.jpg"
var widthVal string
var heightVal string
var saveName string
var itNum string

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
	emptyShortHorizontalSpace := canvas.NewRectangle(blank2)
	emptyShortHorizontalSpace.SetMinSize(fyne.NewSize(10, 12))
	emptyWideVerticalSpace := canvas.NewRectangle(blank2)
	emptyWideVerticalSpace.SetMinSize(fyne.NewSize(200, 10))

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

	image.SetMinSize(fyne.NewSize(350, 350))
	image.FillMode = canvas.ImageFillContain

	promptInput := widget.NewEntry()
	promptInput.SetPlaceHolder("Enter Your Prompts Here...")

	itNumText := canvas.NewText("Iteration :  ", white)
	itNumInput := widget.NewEntry()
	itNumInput.SetPlaceHolder("Enter Iteration here...")
	itNumInput.SetText("50")

	fileNameText := canvas.NewText("File Name :  ", white)
	fileNameInput := widget.NewEntry()
	fileNameInput.SetPlaceHolder("Enter File Name here...")
	fileNameInput.SetText("NewFileName_001")

	centerItNumText := container.NewCenter(itNumText)
	centerFileNameText := container.NewCenter(fileNameText)
	itNumInputVBox := container.NewVBox(itNumInput)
	fileNameInputVBox := container.NewVBox(fileNameInput)
	itNumBorder := container.NewBorder(nil, nil, emptyWideVerticalSpace, emptyWideVerticalSpace, itNumInputVBox)
	fileNameBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyWideVerticalSpace, emptyWideVerticalSpace, fileNameInputVBox)

	widthText := canvas.NewText("Width      :  ", white)
	widthRadio := widget.NewRadioGroup([]string{"256px", "512px", "768px", "1024px"}, func(value string) {
		//log.Println("Radio set to", value)
		if value == "256px" {
			widthVal = "256"
		}
		if value == "512px" {
			widthVal = "512"
		}
		if value == "768px" {
			widthVal = "768"
		}
		if value == "1024px" {
			widthVal = "1024"
		}
		println("width :" + widthVal)
	})
	widthRadio.Horizontal = true
	widthRadio.SetSelected("512px")
	widthSelectorBox := container.NewHBox(widthText, widthRadio)

	heightText := canvas.NewText("Height     :  ", white)
	heightRadio := widget.NewRadioGroup([]string{"256px", "512px", "768px", "1024px"}, func(value string) {
		//log.Println("Radio set to", value)
		if value == "256px" {
			heightVal = "256"
		}
		if value == "512px" {
			heightVal = "512"
		}
		if value == "768px" {
			heightVal = "768"
		}
		if value == "1024px" {
			heightVal = "1024"
		}
		println("height :" + heightVal)
	})
	heightRadio.Horizontal = true
	heightRadio.SetSelected("512px")
	heightSelectorBox := container.NewHBox(heightText, heightRadio)

	os.Chdir("./diffusers")
	generateBtn := widget.NewButton("GENERATE", func() {
		go generateImage(image, promptInput.Text, itNumInput.Text, fileNameInput.Text, widthVal, heightVal)
	})
	refreshBtn := widget.NewButton("REFRESH IMG", func() { go refreshIMG(image, fileNameInput.Text) })
	promptBox := container.NewVBox(promptInput)
	titleBorder := container.NewBorder(emptyHorizontalSpace, emptyHorizontalSpace, nil, nil, titleText)
	imageBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, image)
	promptBorder := container.NewBorder(nil, emptyShortHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, promptBox)
	buttonGrid := container.NewHBox(generateBtn, refreshBtn)
	buttonBorder := container.NewBorder(nil, emptyHorizontalSpace, emptyVerticalSpace, emptyVerticalSpace, buttonGrid)
	centerButtonBorder := container.NewCenter(buttonBorder)

	radioVBox := container.NewVBox(widthSelectorBox, heightSelectorBox)
	radioBorder := container.NewBorder(nil, nil, emptyVerticalSpace, emptyVerticalSpace, radioVBox)
	centerRadioBorder := container.NewCenter(radioBorder)

	contentGrid := container.New(
		layout.NewVBoxLayout(), titleBorder,
		imageBorder, promptBorder, centerButtonBorder,
		centerRadioBorder, centerItNumText, itNumBorder,
		centerFileNameText, fileNameBorder)

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

func generateImage(image *canvas.Image, prompt string, itNum string, fileName string, widthVal string, heightVal string) {
	println("generating")
	//testGenerateImage(prompt, itNum, fileName, widthVal, heightVal)
	GenerateImage(prompt, itNum, fileName, widthVal, heightVal)
	go refreshIMG(image, fileName)
}

func refreshIMG(image *canvas.Image, fileName string) {
	imageUrl = "http://localhost:9998/files/" + fileName + ".jpg"
	println(imageUrl)
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

func testGenerateImage(prompt string, itNums string, fileName string, widthVal string, heightVal string) {
	var testPrompt = prompt
	println(testPrompt, itNums, fileName, widthVal, heightVal)
	var itNum = itNums
	var saveNames = fileName
	var width = widthVal
	var height = heightVal

	GenerateImage(testPrompt, itNum, saveNames, width, height)
	println("finish test generation")
}

func GenerateImage(promtInput string, iterationNums string, saveName string, inputWidth string, inputHeight string) {
	println("Inferencing.... Please wait.")
	println("input data : ", promtInput, iterationNums, saveName, inputWidth, inputHeight)
	exec.Command("python", "script_cuda.py", promtInput, iterationNums, saveName, inputWidth, inputHeight).Run()
	println("Inference completed!🎉")
}

func GenerateImageAdvance(promtInput string, iterationNums string, saveName string, userSeedInput string, PromtUserScheduler string, PromtUserGuidance string, inputWidth string, inputHeight string) {
	println("Inferencing.... Please wait.")
	exec.Command("python", "script_cudaadv.py", promtInput, iterationNums, saveName, userSeedInput, PromtUserScheduler, PromtUserGuidance, inputWidth, inputHeight).Run()
	println("Inference completed!🎉")
}
