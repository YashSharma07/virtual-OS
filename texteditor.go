package main

import (
	"io/ioutil"
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor() {


	// creating new window for text editor
	w:= myApp.NewWindow("TEXT EDITOR")
	
	count := 1



	// multiline input
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter some text here!")

	content:= container.NewVBox(
		container.NewHBox(
			widget.NewLabel("THIS IS A TEXT EDITOR"),
		),
	)


	content.Add(widget.NewButton("Add New File", func ()  {
		
		content.Add(widget.NewLabel("New File "+ strconv.Itoa(count)))
		count++
	}),

)


	// save button functionality
	saveBtn:= widget.NewButton("Save", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)

				uc.Write(textData)
			},w)
		saveFileDialog.SetFileName("New File "+ strconv.Itoa(count) + ".txt")

		w.Resize(fyne.NewSize(500, 500))

		saveFileDialog.Show()
	})
	


	// open file functionality
	openBtn:= widget.NewButton("Open", func(){
		openFileDialog := dialog.NewFileOpen(
			func( r fyne.URIReadCloser, _ error){

				ReadData , _ := ioutil.ReadAll(r)
				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName),
				)

				w.SetContent(container.NewScroll(viewData))

				w.Resize(fyne.NewSize(500, 500))

				w.Show()


			},w)

			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

			openFileDialog.Show()
		})


		// showing the main content on the GUI
		editorContainer := container.NewVBox(
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(1,
					content,
					input,
				),
				container.NewGridWithColumns(1,
					saveBtn,
					openBtn,
				),
				 
			),
		)
	
	w.Resize(fyne.NewSize(400,400))
	
	w.SetContent(container.NewBorder(desktopBtn, nil , nil , editorContainer))

	w.Show()
}