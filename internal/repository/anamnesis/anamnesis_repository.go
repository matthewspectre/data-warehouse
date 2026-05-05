package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	GetAll() ([]*anamnesis.Anamnesis, error)
	GetAllB() ([]*anamnesis.AnamnesisB, error)
}
