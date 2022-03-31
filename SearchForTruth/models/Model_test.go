package models

import (
	"errors"
	"fmt"
	"testing"
)

var projectCollection = make([]Project, 0)

func TestModelProgressionLiner(t *testing.T) {

	var linearTestProject = "TestLinearProject"
	var TestModel3 = Model{
		Id:        "3",
		NextModel: []*Model{},
		Input:     map[string]string{"TestModel2Results": ""},
		// Mocked out data
		Output: map[string]string{"TestModel3Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(linearTestProject, "3", map[string]string{"TestModel3Results": "RESULTS_FROM_TEST_MODEL_3"})
		},
	}

	var TestModel2 = Model{
		Id:        "2",
		NextModel: []*Model{&TestModel3},
		Input:     map[string]string{"TestModel1Results": ""},
		Output:    map[string]string{"TestModel2Results": ""},
		// Mocked out data
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(linearTestProject, "2", map[string]string{"TestModel2Results": "RESULTS_FROM_TEST_MODEL_2"})
		},
	}

	var TestModel1 = Model{
		Id:        "1",
		NextModel: []*Model{&TestModel2},
		Input:     map[string]string{"startingData": "BOB"},
		// Mocked out data
		Output: map[string]string{"TestModel1Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(linearTestProject, "1", map[string]string{"TestModel1Results": "RESULTS_FROM_TEST_MODEL_1"})
		},
	}

	var testProject = Project{
		Id:           linearTestProject,
		Values:       []Value{},
		Requirements: []Requirement{},
		Models:       []Model{TestModel1},
	}
	projectCollection = append(projectCollection, testProject)
	testProject.RunModels(map[string]string{"startingData": "BOB"})

	if TestModel1.Input["startingData"] != "BOB" ||
		TestModel1.Output["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel2.Input["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel2.Output["TestModel2Results"] != "RESULTS_FROM_TEST_MODEL_2" ||
		TestModel3.Input["TestModel2Results"] != "RESULTS_FROM_TEST_MODEL_2" ||
		TestModel3.Output["TestModel3Results"] != "RESULTS_FROM_TEST_MODEL_3" {
		t.Fail()
	}

}

func TestBranchingModels(t *testing.T) {

	var branchTestProject = "TestBranchedProject"
	var TestModel3 = Model{
		Id:        "3",
		NextModel: []*Model{},
		Input:     map[string]string{"TestModel1Results": ""},
		// Mocked out data
		Output: map[string]string{"TestModel3Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(branchTestProject, "3", map[string]string{"TestModel3Results": "RESULTS_FROM_TEST_MODEL_3"})
		},
	}

	var TestModel2 = Model{
		Id:        "2",
		NextModel: []*Model{},
		Input:     map[string]string{"TestModel1Results": ""},
		// Mocked out data
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(branchTestProject, "2", map[string]string{"TestModel2Results": "RESULTS_FROM_TEST_MODEL_2"})
		},
		Output: map[string]string{"TestModel2Results": ""},
	}

	var TestModel1 = Model{
		Id:        "1",
		NextModel: []*Model{&TestModel2, &TestModel3},
		Input:     map[string]string{"startingData": ""},
		// Mocked out data
		Output: map[string]string{"TestModel1Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(branchTestProject, "1", map[string]string{"TestModel1Results": "RESULTS_FROM_TEST_MODEL_1"})
		},
	}

	var testProject = Project{
		Id:           branchTestProject,
		Values:       []Value{},
		Requirements: []Requirement{},
		Models:       []Model{TestModel1},
	}
	projectCollection = append(projectCollection, testProject)
	testProject.RunModels(map[string]string{"startingData": "BOB"})

	if TestModel1.Input["startingData"] != "BOB" ||
		TestModel1.Output["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel2.Input["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel2.Output["TestModel2Results"] != "RESULTS_FROM_TEST_MODEL_2" ||
		TestModel3.Input["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel3.Output["TestModel3Results"] != "RESULTS_FROM_TEST_MODEL_3" {
		t.Fail()
	}
}

func TestMultipleStartPoints(t *testing.T) {
	var multiStartTestProject = "TestMultipleStartsProject"
	var TestModel3 = Model{
		Id:        "3",
		NextModel: []*Model{},
		Input:     map[string]string{"TestModel2Results": ""},
		// Mocked out data
		Output: map[string]string{"TestModel3Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(multiStartTestProject, "3", map[string]string{"TestModel3Results": "RESULTS_FROM_TEST_MODEL_3"})
		},
	}

	var TestModel2 = Model{
		Id:        "2",
		NextModel: []*Model{&TestModel3},
		Input:     map[string]string{"startingData": ""},
		// Mocked out data
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(multiStartTestProject, "2", map[string]string{"TestModel2Results": "RESULTS_FROM_TEST_MODEL_2"})
		},
		Output: map[string]string{"TestModel2Results": ""},
	}

	var TestModel1 = Model{
		Id:        "1",
		NextModel: []*Model{},
		Input:     map[string]string{"startingData": ""},
		// Mocked out data
		Output: map[string]string{"TestModel1Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(multiStartTestProject, "1", map[string]string{"TestModel1Results": "RESULTS_FROM_TEST_MODEL_1"})
		},
	}

	var testProject = Project{
		Id:           multiStartTestProject,
		Values:       []Value{},
		Requirements: []Requirement{},
		Models:       []Model{TestModel1, TestModel2},
	}
	projectCollection = append(projectCollection, testProject)
	testProject.RunModels(map[string]string{"startingData": "BOB"})

	if TestModel1.Input["startingData"] != "BOB" ||
		TestModel1.Output["TestModel1Results"] != "RESULTS_FROM_TEST_MODEL_1" ||
		TestModel2.Input["startingData"] != "BOB" ||
		TestModel2.Output["TestModel2Results"] != "RESULTS_FROM_TEST_MODEL_2" ||
		TestModel3.Input["TestModel2Results"] != "RESULTS_FROM_TEST_MODEL_2" ||
		TestModel3.Output["TestModel3Results"] != "RESULTS_FROM_TEST_MODEL_3" {
		t.Fail()
	}

}

func TestModelWontStartUnlessAllRequirementsAreFullfilled(t *testing.T) {

	var oneModelShouldRunTestProject = "oneModelShouldRunTestProject"
	var TestModel2 = Model{
		Id:        "2",
		NextModel: []*Model{},
		Input:     map[string]string{"TestModel1Results": "", "AnotherReq": ""},
		// Mocked out data
		Output: map[string]string{"TestModel1Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(oneModelShouldRunTestProject, "1", map[string]string{"TestModel1Results": "RESULTS_FROM_TEST_MODEL_1"})
		},
	}

	var TestModel1 = Model{
		Id:        "1",
		NextModel: []*Model{&TestModel2},
		Input:     map[string]string{"startingData": "BOB"},
		// Mocked out data
		Output: map[string]string{"TestModel1Results": ""},
		Action: func() {
			//Do some logic here, the input would be in the model
			// Output the results to the method call, in production this would be a web call
			UpdateModel(oneModelShouldRunTestProject, "1", map[string]string{"TestModel1Results": "RESULTS_FROM_TEST_MODEL_1"})
		},
	}

	var testProject = Project{
		Id:           oneModelShouldRunTestProject,
		Values:       []Value{},
		Requirements: []Requirement{},
		Models:       []Model{TestModel1},
	}
	projectCollection = append(projectCollection, testProject)
	testProject.RunModels(map[string]string{"startingData": "BOB"})

	fmt.Println("hi there world")
}
func assertCondition(projectId string, modelId string, pred func(Model) bool, t *testing.T) {
	foundModel, err := findModelFacade(projectId, modelId)
	if err != nil {
		t.Fail()
	}
	if !pred(*foundModel) {
		t.Fail()
	}
}

//////////////////////////////////////////////////////////////////////////////////////

/*
	What would be Web Logic, putting it here due to importing cycles
*/
func findProject(projectId string) *Project {
	for _, project := range projectCollection {
		if project.Id == projectId {
			return &project
		}
	}
	return nil
}

func findModelFacade(projectId string, modelId string) (*Model, error) {

	var project = findProject(projectId)
	if project == nil {
		return nil, errors.New("No Project found with ID " + projectId)
	}

	for _, initialModel := range project.Models {
		foundModel, err := initialModel.FindModel(modelId)
		if err != nil {
			continue
		}
		return foundModel, nil
	}
	return nil, errors.New("Cannot find Model")
}

func verifyInputsFullfilled(input map[string]string) bool {
	for _, value := range input {
		if len(value) == 0 {
			return false
		}
	}
	return true
}

/*
	Linnker method to transfer outputs of one model to the input of another
*/
func mapOutputToInput(currentModelOutput map[string]string, nextModelinput map[string]string) (map[string]string, error) {
	for nextModelReq, _ := range nextModelinput {
		foundValueForInput, found := currentModelOutput[nextModelReq]
		if !found {
			return nil, errors.New("Cannot find required input for next Model")
		}
		nextModelinput[nextModelReq] = foundValueForInput
	}
	return nextModelinput, nil
}

/*
	Create a method, this would be a webservice in prod
	UpdateModel, first finds the model in the hierachy
	Attempts to
	Update the output
	Validate the output
	TODO persist to disc
	Start downstream Models
*/
func UpdateModel(projectId string, modelId string, updatePacket map[string]string) error {

	// Try to find the model
	foundModel, err := findModelFacade(projectId, modelId)
	// Check that the model exsists, an err will be returned should there not be a model
	if err != nil {
		return err
	}

	// update the output model in memory
	for requirement, value := range updatePacket {
		if _, isFound := foundModel.Output[requirement]; isFound {
			foundModel.Output[requirement] = value
		}
	}

	// Validate if the updated model is valid against the requirement
	if !foundModel.Validate() {
		return errors.New("Invalid output from the updated data")
	}

	// TODO persist the model down to DB // in this test rig, because of in memory and the native structure, it is persisted

	// The current model is now satisfied, time to kickoff the down stream models
	// Need to pass from model to model
	// Check the next Model is not in the ones already visited

	// Launch the next Models
	if foundModel.NextModel != nil {
		for _, nextModel := range foundModel.NextModel {
			// update the next model with its required inputs
			nextActionInput, mappingError := mapOutputToInput(foundModel.Output, nextModel.Input)
			if mappingError != nil {
				return errors.New("There has been a mapping error")
			}
			// verify all of the inputs are present before starting the next model
			if verifyInputsFullfilled(nextActionInput) {
				// Passing in data to the next action
				nextModel.Input = nextActionInput
				nextModel.DoAction()
			}

		}
	} else {
		// We have finished the line, might want to do a call here?
	}
	return nil
}
