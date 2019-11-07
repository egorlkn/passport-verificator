package passport

type Passport struct {
	Serial string
	Number string
}

type ValidateUseCase interface {
	IsValid(passport *Passport) (bool, error)
}
