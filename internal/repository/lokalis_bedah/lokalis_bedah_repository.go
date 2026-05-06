package lokalis_bedah

import entity "rme/internal/entity/lokalis_bedah"

type Repository interface {
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	GetAllB(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	GetAllWarehouse(nik *string) ([]*entity.LokalisBedahWarehouse, error)
	UpsertWarehouseA(rows []*entity.LokalisBedah) (int64, error)
	UpsertWarehouseB(rows []*entity.LokalisBedah) (int64, error)
}
