package pemeriksaan_ekg

type Repository interface {
	GetAll(idPasien *int, idDokter *int) ([]*PemeriksaanEkgWithNames, error)
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
