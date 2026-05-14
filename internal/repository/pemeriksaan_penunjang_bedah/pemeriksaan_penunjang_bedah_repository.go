package pemeriksaan_penunjang_bedah

import entity "rme/internal/entity/pemeriksaan_penunjang_bedah"

type Repository interface {
	Create(d *entity.PemeriksaanPenunjangBedah) error
	GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	GetAllWarehouse(nik *string) ([]*entity.PemeriksaanPenunjangBedahWarehouse, error)
	UpsertWarehouseA(rows []*entity.PemeriksaanPenunjangBedah) (int64, error)
	UpsertWarehouseB(rows []*entity.PemeriksaanPenunjangBedah) (int64, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
	HideByPatient(idPasien int) error
}
