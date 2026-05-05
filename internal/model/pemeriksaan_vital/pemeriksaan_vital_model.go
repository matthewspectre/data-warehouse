package pemeriksaanvital

import "time"

type PemeriksaanVitalModel struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id_pemeriksaan_vital"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
	TekananDarah   string    `gorm:"column:tekanan_darah"`
	DenyutNadi     int       `gorm:"column:denyut_nadi"`
	SuhuTubuh      float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas int       `gorm:"column:frekuensi_napas"`
	BeratBadan     float64   `gorm:"column:berat_badan"`
	TinggiBadan    float64   `gorm:"column:tinggi_badan"`
	Visible        int       `gorm:"column:visible"`
}

func (PemeriksaanVitalModel) TableName() string {
	return "pemeriksaan_vital"
}

// PemeriksaanVitalBModel maps to the rme-b `pemeriksaan_vital` table structure.
type PemeriksaanVitalBModel struct {
	ID               int       `gorm:"primaryKey;autoIncrement;column:id_pemeriksaan_vital"`
	IDPasien         int       `gorm:"column:id_pasien"`
	IDDokter         int       `gorm:"column:id_dokter"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
	TekananDarah     string    `gorm:"column:tekanan_darah"`
	DenyutNadi       int       `gorm:"column:denyut_nadi"`
	SuhuTubuh        float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas   int       `gorm:"column:frekuensi_napas"`
	BeratBadan       float64   `gorm:"column:berat_badan"`
	TinggiBadan      float64   `gorm:"column:tinggi_badan"`
	SaturasiOksigen  int       `gorm:"column:saturasi_oksigen"`
	TingkatKesadaran string    `gorm:"column:tingkat_kesadaran"`
	IndeksMasaTubuh  float64   `gorm:"column:indeks_masa_tubuh"`
	Visible          int       `gorm:"column:visible"`
}

func (PemeriksaanVitalBModel) TableName() string {
	return "pemeriksaan_vital"
}

// PemeriksaanVitalWarehouseModel maps to the `data_warehouse`.pemeriksaan_vital table.
// It stores data from both rsA and rsB with a `source` marker.
type PemeriksaanVitalWarehouseModel struct {
	Source             string    `gorm:"column:source"`
	IDPemeriksaanVital int       `gorm:"column:id_pemeriksaan_vital"`
	IDPasien           int       `gorm:"column:id_pasien"`
	NamaPasien         *string   `gorm:"column:nama_pasien"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	TekananDarah       string    `gorm:"column:tekanan_darah"`
	DenyutNadi         int       `gorm:"column:denyut_nadi"`
	SuhuTubuh          float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas     int       `gorm:"column:frekuensi_napas"`
	BeratBadan         float64   `gorm:"column:berat_badan"`
	TinggiBadan        float64   `gorm:"column:tinggi_badan"`
	SaturasiOksigen    *int      `gorm:"column:saturasi_oksigen"`
	TingkatKesadaran   *string   `gorm:"column:tingkat_kesadaran"`
	IndeksMasaTubuh    *float64  `gorm:"column:indeks_masa_tubuh"`
	Visible            int       `gorm:"column:visible"`
}

func (PemeriksaanVitalWarehouseModel) TableName() string {
	return "`data_warehouse`.pemeriksaan_vital"
}
