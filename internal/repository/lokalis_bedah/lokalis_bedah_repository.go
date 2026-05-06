package lokalis_bedah

import entity "rme/internal/entity/lokalis_bedah"

type Repository interface {
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
}
