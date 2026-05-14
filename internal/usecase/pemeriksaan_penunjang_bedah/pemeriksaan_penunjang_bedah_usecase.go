package pemeriksaan_penunjang_bedah

import (
	"time"

	entity "rme/internal/entity/pemeriksaan_penunjang_bedah"
	repo "rme/internal/repository/pemeriksaan_penunjang_bedah"
)

type Usecase interface {
	Create(d *entity.PemeriksaanPenunjangBedah) error
	GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	GetAllWarehouse(nik *string) ([]*entity.PemeriksaanPenunjangBedahWarehouse, error)
	ETLToWarehouse() (ETLResult, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
	HideByPatient(idPasien int) error
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

func (u *usecase) Create(d *entity.PemeriksaanPenunjangBedah) error {
	if d.DateMake.IsZero() {
		d.DateMake = time.Now()
	}
	if d.DateUpdate.IsZero() {
		d.DateUpdate = d.DateMake
	}
	return u.repo.Create(d)
}

func (u *usecase) GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	return u.repo.GetAll(idDokter, idPasien)
}

func (u *usecase) GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	return u.repo.GetAllB(idDokter, idPasien)
}

func (u *usecase) GetAllWarehouse(nik *string) ([]*entity.PemeriksaanPenunjangBedahWarehouse, error) {
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

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}

func (u *usecase) Hide(id int) error {
	return u.repo.Hide(id)
}

func (u *usecase) HideByPatient(idPasien int) error {
	return u.repo.HideByPatient(idPasien)
}
