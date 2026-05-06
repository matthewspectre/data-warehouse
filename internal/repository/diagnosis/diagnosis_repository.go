package diagnosis

import "rme/internal/entity/diagnosis"

type Repository interface {
	Create(d *diagnosis.Diagnosis) error
	GetByID(id int) (*diagnosis.Diagnosis, error)
	GetAll(idDokter *int, idPasien *int) ([]*diagnosis.Diagnosis, error)
	GetAllB(idDokter *int, idPasien *int) ([]*diagnosis.Diagnosis, error)
	ETLUpsertWarehouseA(rows []*diagnosis.Diagnosis) (int64, error)
	ETLUpsertWarehouseB(rows []*diagnosis.Diagnosis) (int64, error)
	GetAllWarehouse(nik *string) ([]*diagnosis.DiagnosisWarehouse, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
	GetIcdNamesForCodes(codes []string) (map[string]string, error)
}
