package main

type MaximizeFontRequest struct {
	fontFamily string `json:"fontFamily"`
	text string `json:"text"`
	boxWidth int `json:"boxWidth"`
	boxHeight int `json:"boxHeight"`
}

type MaximizeFontResponse struct {
	text string `json:"text"`
	fontSize int `json:"fontSize"`
}