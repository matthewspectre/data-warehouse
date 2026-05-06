package pemeriksaan_ekg

import (
	repo "rme/internal/repository/pemeriksaan_ekg"
)

type Usecase interface {
	GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error)
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
