package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	GetAll() ([]*anamnesis.Anamnesis, error)
	GetAllB() ([]*anamnesis.AnamnesisB, error)
	GetAllWarehouse(nik *string) ([]*anamnesis.AnamnesisWarehouse, error)
	UpsertWarehouseA(rows []*anamnesis.Anamnesis) (int64, error)
	UpsertWarehouseB(rows []*anamnesis.AnamnesisB) (int64, error)
}
