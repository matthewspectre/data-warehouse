package anamnesis

// Use case (business logic) for anamnesis.

import (
	entity "rme/internal/entity/anamnesis"
	repo "rme/internal/repository/anamnesis"
)

// Usecase mendefinisikan operasi bisnis untuk Anamnesis.
// Di atas repository (DB), di bawah handler (HTTP/API).
type Usecase interface {
	GetAll() ([]*entity.Anamnesis, error)
	GetAllB() ([]*entity.AnamnesisB, error)
	GetAllWarehouse(nik *string) ([]*entity.AnamnesisWarehouse, error)
	ETLToWarehouse() (ETLResult, error)
}

type ETLResult struct {
	InsertedRSA int64
	InsertedRSB int64
}

// usecase adalah implementasi konkret dari Usecase.
type usecase struct {
	repo repo.Repository
}

// NewUsecase membuat instance baru usecase Anamnesis.
func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

// GetAll mengambil semua data Anamnesis yang masih visible.
// Jika `idDokter` atau `idPasien` tidak nil, hasil akan difilter berdasarkan kolom terkait.
func (u *usecase) GetAll() ([]*entity.Anamnesis, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetAllB() ([]*entity.AnamnesisB, error) {
	return u.repo.GetAllB()
}

func (u *usecase) GetAllWarehouse(nik *string) ([]*entity.AnamnesisWarehouse, error) {
	return u.repo.GetAllWarehouse(nik)
}

func (u *usecase) ETLToWarehouse() (ETLResult, error) {
	rsA, err := u.repo.GetAll()
	if err != nil {
		return ETLResult{}, err
	}
	rsB, err := u.repo.GetAllB()
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
