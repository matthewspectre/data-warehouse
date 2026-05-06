package lokalis_bedah

import (
	entity "rme/internal/entity/lokalis_bedah"
	repo "rme/internal/repository/lokalis_bedah"
)

type Usecase interface {
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	GetAllB(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	GetAllWarehouse(nik *string) ([]*entity.LokalisBedahWarehouse, error)
	ETLToWarehouse() (ETLResult, error)
}

type ETLResult struct {
	InsertedRSA int64
	InsertedRSB int64
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	return u.repo.GetAll(idDokter, idPasien)
}

func (u *usecase) GetAllB(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	return u.repo.GetAllB(idDokter, idPasien)
}

func (u *usecase) GetAllWarehouse(nik *string) ([]*entity.LokalisBedahWarehouse, error) {
	return u.repo.GetAllWarehouse(nik)
}

func (u *usecase) ETLToWarehouse() (ETLResult, error) {
	rsA, err := u.repo.GetAll(nil, nil)
	if err != nil {
		return ETLResult{}, err
	}
	rsB, err := u.repo.GetAllB(nil, nil)
	if err != nil {
		return ETLResult{}, err
	}

	insertedA, err := u.repo.UpsertWarehouseA(rsA)
	if err != nil {
		return ETLResult{}, err
	}
	insertedB, err := u.repo.UpsertWarehouseB(rsB)
	if err != nil {
		return ETLResult{}, err
	}

	return ETLResult{InsertedRSA: insertedA, InsertedRSB: insertedB}, nil
}
