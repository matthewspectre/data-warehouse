package pemeriksaan_fungsi_organ

import entity "rme/internal/entity/pemeriksaan_fungsi_organ"

type Repository interface {
	Create(d *entity.PemeriksaanFungsiOrgan) error
	GetByID(id int) (*entity.PemeriksaanFungsiOrgan, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error)
	GetAllB(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error)
	GetAllWarehouse(nik *string) ([]*entity.PemeriksaanFungsiOrganWarehouse, error)
	UpsertWarehouseA(rows []*entity.PemeriksaanFungsiOrgan) (int64, error)
	UpsertWarehouseB(rows []*entity.PemeriksaanFungsiOrgan) (int64, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
}
