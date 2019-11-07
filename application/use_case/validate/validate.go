package validate

import "github.com/egorlkn/passport-verificator/domain/passport"

type interactor struct {
	repository InvalidPassportRepository
}

func (interactor *interactor) IsValid(passport *passport.Passport) (bool, error) {
	isExist, err := interactor.repository.IsExist(passport)
	if err != nil {
		return false, err
	}

	return !isExist, nil
}

func NewPassportValidateInteractor(repository InvalidPassportRepository) *interactor {
	return &interactor{repository: repository}
}

type InvalidPassportRepository interface {
	IsExist(passport *passport.Passport) (bool, error)
}
