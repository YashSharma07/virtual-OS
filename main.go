// this tells the compiler that this file should be executed as an executable program
package main


//import the required libraries
import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)


//creating new app and window
var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("PEP OS")


//creating various button variables of fyne.Widget type
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var desktopBtn fyne.Widget


// creating image variable of fyne.CanvasObject
var img fyne.CanvasObject 

var panelContent *fyne.Container


// main function 
func main() {

	// image assignment
	img = canvas.NewImageFromFile("a.jpg")


	// assigning the previously declared buttons
	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalculatorApp()
	})

	btn3 = widget.NewButtonWithIcon("Gallery App", theme.StorageIcon(), func() {
		showGalleryApp(myWindow)
	})

	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentCreateIcon(), func() {
		showTextEditor()
	})

	btn5 = widget.NewButtonWithIcon("Movie Adda!", theme.MediaPlayIcon(), func() {
		showMovieAdda(myWindow)
	})

	btn6 = widget.NewButtonWithIcon("Music Player", theme.MediaRecordIcon(), func() {
		showMusicPlayer(myWindow)
	})

	desktopBtn = widget.NewButtonWithIcon("DESKTOP", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	//defining our panel content (whatever comes as navigation bar at the top)
	panelContent = container.NewVBox((container.NewGridWithColumns(7, desktopBtn, btn1, btn2, btn3, btn4, btn5, btn6)))

	myWindow.Resize(fyne.NewSize(1380, 720))
	myWindow.CenterOnScreen()

	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	myWindow.ShowAndRun()

}
