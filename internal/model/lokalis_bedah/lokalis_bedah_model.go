package lokalis_bedah

import "time"

type LokalisBedahModel struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	Tanggal        time.Time `gorm:"column:tanggal"`
	LokasiKelainan string    `gorm:"column:lokasi_kelainan"`
	JenisKelainan  string    `gorm:"column:jenis_kelainan"`
	Ukuran         string    `gorm:"column:ukuran"`
	Warna          string    `gorm:"column:warna"`
	NyeriTekan     bool      `gorm:"column:nyeri_tekan"`
	Konsistensi    string    `gorm:"column:konsistensi"`
	Mobilitas      string    `gorm:"column:mobilitas"`
	TandaRadang    bool      `gorm:"column:tanda_radang"`
	Fluktuasi      bool      `gorm:"column:fluktuasi"`
	Catatan        string    `gorm:"column:catatan"`
	Visible        int       `gorm:"column:visible"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
}

func (LokalisBedahModel) TableName() string {
	return "status_lokalis_bedah"
}

// LokalisBedahBModel maps to the rme-b `status_lokalis_bedah` table structure.
// Catatan: struktur rsB tidak memiliki kolom warna/konsistensi/mobilitas.
type LokalisBedahBModel struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	Tanggal        time.Time `gorm:"column:tanggal"`
	LokasiKelainan string    `gorm:"column:lokasi_kelainan"`
	JenisKelainan  string    `gorm:"column:jenis_kelainan"`
	Ukuran         string    `gorm:"column:ukuran"`
	NyeriTekan     bool      `gorm:"column:nyeri_tekan"`
	TandaRadang    bool      `gorm:"column:tanda_radang"`
	Fluktuasi      bool      `gorm:"column:fluktuasi"`
	Catatan        string    `gorm:"column:catatan"`
	Visible        int       `gorm:"column:visible"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
}

func (LokalisBedahBModel) TableName() string {
	return "status_lokalis_bedah"
}

// LokalisBedahWarehouseModel maps to the `data_warehouse`.status_lokalis_bedah table.
// It stores data from both rsA and rsB with a `source` marker.
type LokalisBedahWarehouseModel struct {
	Source               string    `gorm:"column:source"`
	IDStatusLokalisBedah int       `gorm:"column:id_status_lokalis_bedah"`
	IDPasien             int       `gorm:"column:id_pasien"`
	NamaPasien           *string   `gorm:"column:nama_pasien"`
	IDDokter             int       `gorm:"column:id_dokter"`
	Tanggal              time.Time `gorm:"column:tanggal"`
	LokasiKelainan       string    `gorm:"column:lokasi_kelainan"`
	JenisKelainan        string    `gorm:"column:jenis_kelainan"`
	Ukuran               string    `gorm:"column:ukuran"`
	Warna                string    `gorm:"column:warna"`
	NyeriTekan           bool      `gorm:"column:nyeri_tekan"`
	Konsistensi          string    `gorm:"column:konsistensi"`
	Mobilitas            string    `gorm:"column:mobilitas"`
	TandaRadang          bool      `gorm:"column:tanda_radang"`
	Fluktuasi            bool      `gorm:"column:fluktuasi"`
	Catatan              string    `gorm:"column:catatan"`
	Visible              int       `gorm:"column:visible"`
	DateMake             time.Time `gorm:"column:date_make"`
	DateUpdate           time.Time `gorm:"column:date_update"`
}

func (LokalisBedahWarehouseModel) TableName() string {
	return "`data_warehouse`.status_lokalis_bedah"
}
