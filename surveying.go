package main

import "gopkg.in/AlecAivazis/survey.v1"

var MainMenu = []*survey.Question{
	{
		Name: "tool-select",
		Prompt: &survey.Select{
			Message: "Select a tool.",
			Options: []string{"Sign & Encrypt", "Forecast Locations Parser", "Quit"},
			Default: "Sign & Encrypt",
		},
	},
}
