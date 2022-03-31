package models

type Value struct {
}

type Requirement struct {
}

type Project struct {
	Id           string
	Values       []Value
	Requirements []Requirement
	// Top level model is a starting point, not the entirity
	Models []Model
}

/*
	Method for the project to start running the assosiated models
*/
func (project *Project) RunModels(input map[string]string) error {

	// Validate the starting models have what they need from the input
	for _, startingModel := range project.Models {
		for modelRequirement, _ := range startingModel.Input {
			if inputValueNeeded, isfound := input[modelRequirement]; isfound {
				// Update the Input in the starting model ready to be kicked off
				startingModel.Input[modelRequirement] = inputValueNeeded
			}
		}
	}
	// Start the initial models
	for _, startingModel := range project.Models {
		startingModel.Action()
	}

	return nil

}
