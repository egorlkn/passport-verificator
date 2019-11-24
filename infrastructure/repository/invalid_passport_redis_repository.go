package repository

import "github.com/egorlkn/passport-verificator/domain/passport"

type InvalidPassportRedisRepository struct {
}

func (repository *InvalidPassportRedisRepository) IsExist(passport *passport.Passport) (bool, error) {
	return false, nil
}
