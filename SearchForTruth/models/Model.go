package models

import (
	"errors"
)

type Model struct {
	Id            string
	NextModel     []*Model
	PreviousModel *Model
	Output        map[string]string
	Input         map[string]string
	// have to put this in for testing purposes
	Action func()
}

/*
	Validate the update against the requirements
	TODO
*/
func (model Model) Validate() bool {
	return true
}

/*
 Each Model will have an action, first pass implementation, highly likely to change.
*/
func (model *Model) DoAction() {
	// Supply the Model inputs, Project Id, Model Id and URL to send back the results
	model.Action()
}

/*
	Public FindModel, submits a default value for vistingModels
*/
func (model *Model) FindModel(modelId string) (*Model, error) {
	emptySlice := make([]string, 0)
	return model.findModel(modelId, &emptySlice)
}

/*
	Private findModel method as Golang does not support default values in the parameters
	Using Depth First recussive call to search for the model
	Note Using the  visitedModelId, because we cannot change a pointers' slice, then each searcher will have its own copy of where it has been
*/
func (model *Model) findModel(findingModelId string, visitedModelIds *[]string) (*Model, error) {

	// Happy path, found the model in question, return it
	if model.Id == findingModelId {
		return model, nil
	}
	// Have not found the model in question, record it as visted, as a mechanism for finding out cyclic dependency
	updatedVisted := append(*visitedModelIds, model.Id)

	// Check for presence of nextModel, if there is none then there is no model found
	if model.NextModel == nil || len(model.NextModel) == 0 {
		return nil, errors.New("No Model Found with ID " + findingModelId)
	}

	// Search for the model downstream
	for _, futureModel := range model.NextModel {
		if result, _ := futureModel.findModel(findingModelId, &updatedVisted); result != nil {
			return result, nil
		}
	}

	return nil, errors.New("No Model Found with ID " + findingModelId)

}
