package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

// creating the showCalculatorApp
func showCalculatorApp() {


	// various variable declarations
	output := ""

	isHistory := false
	historyText := ""


	// creating a history label
	history := widget.NewLabel(historyText)
	var historyArr []string


	// creating the input label
	input := widget.NewLabel(output)

	historyBtn := widget.NewButton("HISTORY", func() {

		if isHistory {
			historyText = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyText = historyText + historyArr[i]
				historyText += "\n"
			}
		}
		isHistory = !isHistory
		history.SetText(historyText)

	})
	backBtn := widget.NewButton("BACK", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn := widget.NewButton("CLEAR", func() {
		output = ""
		input.SetText(output)
	})
	openBracketBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBracketBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	divideBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	addBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {

			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				output = "INVALID EXPRESSION"
			}

		} else {
			output = "ERROR"
		}
		input.SetText(output)
	})


	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	fourBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})


	equalBtn.Importance = widget.HighImportance

	calcContainer := container.NewVBox(container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn),
			container.NewGridWithColumns(4,
				clearBtn,
				openBracketBtn,
				closeBracketBtn,
				divideBtn),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				addBtn),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				minusBtn),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn),
				container.NewGridWithColumns(1,
					equalBtn),
			),
		),
	),
	)

	w:= myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(400,480))

	// the new border := (top, bottom, left, right, fyne.CanvasObject)
	w.SetContent(container.NewBorder(desktopBtn, nil, nil,nil, calcContainer))

	w.Show()
}
