package farm

import (
	"errors"
	"fmt"
)

// direcly created by exercism
type WeightFodder interface {
	FodderAmount() (float64, error)
}

type SillyNephewError struct {
	numberOfCows int
}

// directy created by exercism
var ErrScaleMalfunction = errors.New("sensor error")

func (s *SillyNephewError) Error() string {
	returnError := ""
	if s.numberOfCows < 0 {
		returnError = fmt.Sprintf("silly nephew, there cannot be %v cows", s.numberOfCows)
	}
	return returnError
}


func DivideFood(w WeightFodder, cows int) (float64, error) {

	fodderAmountResult, err := w.FodderAmount()

	if cows == 0 {
		cantDivideByZero := errors.New("division by zero")
		return 0.0, cantDivideByZero
	} 
	if err == ErrScaleMalfunction && fodderAmountResult > 0 {
		newFodderAmountResult := fodderAmountResult * 2
		divideFoodResult := newFodderAmountResult / float64(cows)
		return divideFoodResult, nil
	} 
	if err == ErrScaleMalfunction && fodderAmountResult < 0 {
		newError := errors.New("negative fodder")
		return 0.0, newError
	} 
	if fodderAmountResult < 0 && err == nil {
		newError := errors.New("negative fodder")
		return 0.0, newError
	}  
	if cows < 0 {
		sillyManagement := SillyNephewError{
			numberOfCows: cows,
		}
		newError := sillyManagement.Error()
		displayError := errors.New(newError)
		return 0.0, displayError
	} 
	if cows < 0 && fodderAmountResult < 0 && err == ErrScaleMalfunction || cows < 0  && fodderAmountResult < 0 && err == nil {
		sillyManagement := SillyNephewError{
			numberOfCows: cows,
		}
		newError := sillyManagement.Error()
		displayError := errors.New(newError)
		return 0.0, displayError
	}
	if err != nil  {
		return 0.0, err
	}  
	divideFoodResult := fodderAmountResult / float64(cows)
	return divideFoodResult, err
}