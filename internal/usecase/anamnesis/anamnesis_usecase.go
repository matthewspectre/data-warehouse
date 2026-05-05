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
