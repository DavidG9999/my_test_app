package repository

type Authorization interface {
}

type AutoRepository interface {
}

type ContragentRepository interface {
}

type DispetcherRepository interface {
}

type DriverRepository interface {
}

type MechanicRepository interface {
}

type OrganizationRepository interface {
}

type PutlistRepository interface {
}

type Repository struct {
	Authorization
	AutoRepository
	ContragentRepository
	DispetcherRepository
	DriverRepository
	MechanicRepository
	OrganizationRepository
	PutlistRepository
}

func NewRepository() *Repository {
	return &Repository{}
}
