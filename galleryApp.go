package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {
	

	// creating tabs for each image
	tabs := container.NewAppTabs()


	// path of the root directory
	root_src := "C:\\Users\\91999\\Pictures\\Video Projects"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}

	// appending the files of png and jpg extension.
	for _, file := range files {
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" {
				tabs.Append(container.NewTabItem(file.Name(), canvas.NewImageFromFile(root_src+"\\"+file.Name())))
			}

		}
	}


	// can be either of trailing, leading, bottom or top as default
	tabs.SetTabLocation(container.TabLocationLeading)

	// w.Resize(fyne.NewSize(900, 480))
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, tabs))

	w.Show()
}
