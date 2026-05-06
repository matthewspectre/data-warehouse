package pemeriksaan_ekg

import (
	repo "rme/internal/repository/pemeriksaan_ekg"
)

type Usecase interface {
	GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error)
	GetAllB(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNamesB, error)
	GetAllWarehouse(nik *string) ([]*repo.PemeriksaanEkgWarehouse, error)
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

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) GetAllB(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNamesB, error) {
	return u.repo.GetAllB(idPasien, idDokter)
}

func (u *usecase) GetAllWarehouse(nik *string) ([]*repo.PemeriksaanEkgWarehouse, error) {
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
