package main

import (
	"fmt"
	"sort"
)

const (
	FONT_ROBOTO = "ROBOTO"
	FONT_TIMES = "TIMES"
	FONT_COMIC = "COMIC"

	FONT_MAX =  1000000000
)

func font_metrics(font_family string, text string, font_size int) (float64,float64) {
	width_factor, height_factor := 1.0,1.0
	switch font_family {
	case FONT_ROBOTO:
		width_factor, height_factor = 0.45, 0.9
	case FONT_TIMES:
		width_factor, height_factor = 0.6, 0.95
	case FONT_COMIC:
		width_factor, height_factor = 0.5, 0.85
	}
	width := width_factor * float64(len(text)) * float64(font_size)
	height := height_factor * float64(len(text)) *float64(font_size)
	return width, height
}

func MaximizeFont(params *MaximizeFontRequest) (MaximizeFontResponse, error){
	maxFont, maxStr := maxFontWithSplitting(params.fontFamily, params.text, float64(params.boxWidth), float64(params.boxHeight))
	return MaximizeFontResponse{
		text:     maxStr,
		fontSize: maxFont,
	}, nil
	//maxFont := maxFontWithoutSplitting(params.fontFamily, params.text, float64(params.boxWidth), float64(params.boxHeight))
	//return MaximizeFontResponse{
	//	text:     params.text,
	//	fontSize: maxFont,
	//}, nil
}

var splits []int

func min(a,b int) int{
	if a >= b {
		return b
	} else {
		return a
	}
}

func maxFontWithSplitting(fontFamily string, text string, maxWidth float64, maxHeight float64) (int,string) {
	maxFont := maxFontWithoutSplitting(fontFamily, text, maxWidth, maxHeight)
	maxStr := text
	for i,c := range text {
		if string(c) == " " {
			leftMax := maxFontWithoutSplitting(fontFamily, text[:i], maxWidth, maxHeight/2)
			rightMax, rightStr  := maxFontWithSplitting(fontFamily, text[i+1:], maxWidth, maxHeight/2)

			if min(leftMax, rightMax) > maxFont {
				maxFont = min(leftMax, rightMax)
				maxStr = text[:i] + "\n" + rightStr
			}
		}
	}
	return maxFont, maxStr
}

func maxFontWithoutSplitting(fontFamily string, text string, maxWidth float64, maxHeight float64) int {
	maxFontByWidth := sort.Search( FONT_MAX, func(fontSize int) bool {
		w,_ := font_metrics(fontFamily, text, fontSize)
		if maxWidth < w {
			return true
		} else {
			return false
		}
	})
	maxFontByHeight := sort.Search( FONT_MAX, func(fontSize int) bool {
		_,h := font_metrics(fontFamily, text, fontSize)
		if maxHeight < h {
			return true
		} else {
			return false
		}
	})
	fontMax := min(maxFontByWidth, maxFontByHeight)
	wMax, hMax := font_metrics(fontFamily, text, fontMax)
	fmt.Printf("wMax: %d, hMax: %d \n", int(wMax), int(hMax) )
	return fontMax
}