package validate

import (
	"errors"
	"github.com/egorlkn/passport-verificator/domain/passport"
	"testing"
)

func TestIsValid(t *testing.T) {
	repositoryMock := &invalidPassportRepositoryMock{}

	interactor := NewPassportValidateInteractor(repositoryMock)

	fakeValidPassport := &passport.Passport{Serial: "1234", Number: "654321"}
	trueResult, _ := interactor.IsValid(fakeValidPassport)
	if trueResult != true {
		t.Error("Expected true, got", trueResult)
	}

	fakeInvalidPassport := &passport.Passport{Serial: "4321", Number: "123456"}
	falseResult, _ := interactor.IsValid(fakeInvalidPassport)
	if falseResult != false {
		t.Error("Expected false, got", falseResult)
	}

	emptyPassport := &passport.Passport{Serial: "", Number: ""}
	_, err := interactor.IsValid(emptyPassport)
	if err == nil {
		t.Error("Expected error, got", nil)
	}
}

type invalidPassportRepositoryMock struct {
}

func (repository *invalidPassportRepositoryMock) IsExist(passport *passport.Passport) (bool, error) {
	if passport.Serial == "" && passport.Number == "" {
		return false, errors.New("")
	}

	result := passport.Serial == "4321" && passport.Number == "123456"

	return result, nil
}
