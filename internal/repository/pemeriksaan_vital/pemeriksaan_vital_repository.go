package pemeriksaanvital

import (
	entity "rme/internal/entity/pemeriksaan_vital"
)

type Repository interface {
	GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error)
	GetAllB(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVitalB, error)
	GetAllWarehouse(nik *string) ([]*entity.PemeriksaanVitalWarehouse, error)
	UpsertWarehouseA(rows []*entity.PemeriksaanVital) (int64, error)
	UpsertWarehouseB(rows []*entity.PemeriksaanVitalB) (int64, error)
}
