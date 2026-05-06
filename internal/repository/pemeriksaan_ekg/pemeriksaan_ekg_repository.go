package pemeriksaan_ekg

type Repository interface {
	GetAll(idPasien *int, idDokter *int) ([]*PemeriksaanEkgWithNames, error)
	GetAllB(idPasien *int, idDokter *int) ([]*PemeriksaanEkgWithNamesB, error)
	GetAllWarehouse(nik *string) ([]*PemeriksaanEkgWarehouse, error)
	UpsertWarehouseA(rows []*PemeriksaanEkgWithNames) (int64, error)
	UpsertWarehouseB(rows []*PemeriksaanEkgWithNamesB) (int64, error)
}

type PemeriksaanEkgWithNames struct {
	ID                 int
	IDPasien           int
	NamaPasien         string
	IDDokter           int
	NamaDokter         string
	Visible            int
	DetakJantung       *int
	Irama              string
	PRInterval         *float64
	QRSDuration        *float64
	QTTcInterval       *float64
	AxisJantung        string
	STElevationDepress string
	TWaveAbnormality   string
	InterpretasiDokter string
	DateMake           string
}

type PemeriksaanEkgWithNamesB struct {
	ID                 int
	IDPasien           int
	NamaPasien         string
	IDDokter           int
	NamaDokter         string
	Visible            int
	DetakJantung       *int
	Irama              string
	PRInterval         *float64
	QRSDuration        *float64
	QTTcInterval       *float64
	PWave              string
	AVBlock            string
	InterpretasiDokter string
	DateMake           string
	DateUpdate         string
}

type PemeriksaanEkgWarehouse struct {
	Source           string
	IDPemeriksaanEkg int
	IDPasien         int
	NamaPasien       *string
	Visible          int

	DetakJantung *int
	Irama        string
	PRInterval   *float64
	QRSDuration  *float64
	QTTcInterval *float64

	AxisJantung        string
	STElevationDepress string
	TWaveAbnormality   string

	PWave   string
	AVBlock string

	InterpretasiDokter string
	DateMake           string
	DateUpdate         string
}
