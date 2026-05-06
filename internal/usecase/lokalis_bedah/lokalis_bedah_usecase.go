package lokalis_bedah

import (
	entity "rme/internal/entity/lokalis_bedah"
	repo "rme/internal/repository/lokalis_bedah"
)

type Usecase interface {
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
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
