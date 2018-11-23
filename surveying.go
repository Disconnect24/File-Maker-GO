package main

import "gopkg.in/AlecAivazis/survey.v1"

var MainMenu = []*survey.Question{
	{
		Name: "tool-select",
		Prompt: &survey.Select{
			Message: "Select a tool.",
			Options: []string{"Forecast Locations Parser", "Background Mode", "Quit"},
			Default: "Background Mode",
		},
	},
}
