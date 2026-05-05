package pemeriksaanvital

import (
	entity "rme/internal/entity/pemeriksaan_vital"
	repo "rme/internal/repository/pemeriksaan_vital"
)

type Usecase interface {
	GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error)
	GetAllB(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVitalB, error)
	GetAllWarehouse(nik *string) ([]*entity.PemeriksaanVitalWarehouse, error)
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

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) GetAllB(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVitalB, error) {
	return u.repo.GetAllB(idPasien, idDokter)
}

func (u *usecase) GetAllWarehouse(nik *string) ([]*entity.PemeriksaanVitalWarehouse, error) {
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
