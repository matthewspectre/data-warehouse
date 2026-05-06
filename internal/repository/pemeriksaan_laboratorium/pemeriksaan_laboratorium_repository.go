package pemeriksaan_laboratorium

type Repository interface {
	GetAll(idPasien *int, idDokter *int) ([]*PemeriksaanLaboratoriumWithNames, error)
	GetAllB(idPasien *int, idDokter *int) ([]*PemeriksaanLaboratoriumWithNamesB, error)
	GetAllWarehouse(nik *string) ([]*PemeriksaanLaboratoriumWarehouse, error)
	UpsertWarehouseA(rows []*PemeriksaanLaboratoriumWithNames) (int64, error)
	UpsertWarehouseB(rows []*PemeriksaanLaboratoriumWithNamesB) (int64, error)
}

type PemeriksaanLaboratoriumWithNames struct {
	ID           int
	IDPasien     int
	NamaPasien   string
	IDDokter     int
	NamaDokter   string
	Visible      int
	Hb           *float64
	Ht           *float64
	Leukosit     *int
	Trombosit    *int
	GulaPuasa    *float64
	GulaSewaktu  *float64
	HbA1c        *float64
	Kolesterol   *float64
	HDL          *float64
	LDL          *float64
	Trigliserida *float64
	SGOT         *float64
	SGPT         *float64
	Ureum        *float64
	Kreatinin    *float64
	AsamUrat     *float64
	Natrium      *float64
	Kalium       *float64
	Klorida      *float64
	Waktu        string
}

type PemeriksaanLaboratoriumWithNamesB struct {
	ID             int
	IDPasien       int
	NamaPasien     string
	IDDokter       int
	NamaDokter     string
	Visible        int
	Hb             *float64
	Ht             *float64
	Leukosit       *int
	Trombosit      *int
	GulaPuasa      *float64
	GulaSewaktu    *float64
	HbA1c          *float64
	Kolesterol     *float64
	HDL            *float64
	LDL            *float64
	Trigliserida   *float64
	SGOT           *float64
	SGPT           *float64
	Ureum          *float64
	Kreatinin      *float64
	AsamUrat       *float64
	Natrium        *float64
	Kalium         *float64
	Klorida        *float64
	Waktu          string
	LajuEndapDarah *float64
	Albumin        *float64
}

type PemeriksaanLaboratoriumWarehouse struct {
	Source                    string
	IDPemeriksaanLaboratorium int
	IDPasien                  int
	NamaPasien                *string
	Hb                        *float64
	Ht                        *float64
	Leukosit                  *int
	Trombosit                 *int
	GulaPuasa                 *float64
	GulaSewaktu               *float64
	HbA1c                     *float64
	Kolesterol                *float64
	HDL                       *float64
	LDL                       *float64
	Trigliserida              *float64
	SGOT                      *float64
	SGPT                      *float64
	Ureum                     *float64
	Kreatinin                 *float64
	AsamUrat                  *float64
	Natrium                   *float64
	Kalium                    *float64
	Klorida                   *float64
	Waktu                     string
	LajuEndapDarah            *float64
	Albumin                   *float64
	Visible                   int
}
