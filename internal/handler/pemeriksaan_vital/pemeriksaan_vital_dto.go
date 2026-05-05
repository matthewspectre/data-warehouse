package pemeriksaanvital

import "time"

type PemeriksaanVitalResponse struct {
	ID             int       `json:"id_pemeriksaan_vital"`
	IDPasien       int       `json:"id_pasien"`
	IDDokter       int       `json:"id_dokter"`
	DateMake       time.Time `json:"date_make"`
	DateUpdate     time.Time `json:"date_update"`
	TekananDarah   string    `json:"tekanan_darah"`
	DenyutNadi     int       `json:"denyut_nadi"`
	SuhuTubuh      float64   `json:"suhu_tubuh"`
	FrekuensiNapas int       `json:"frekuensi_napas"`
	BeratBadan     float64   `json:"berat_badan"`
	TinggiBadan    float64   `json:"tinggi_badan"`
	Visible        int       `json:"visible"`
	NamaPasien     string    `json:"nama_pasien"`
	NamaDokter     string    `json:"nama_dokter"`
}

type PemeriksaanVitalBResponse struct {
	Source           string    `json:"source"`
	ID               int       `json:"id_pemeriksaan_vital"`
	IDPasien         int       `json:"id_pasien"`
	IDDokter         int       `json:"id_dokter"`
	DateMake         time.Time `json:"date_make"`
	DateUpdate       time.Time `json:"date_update"`
	TekananDarah     string    `json:"tekanan_darah"`
	DenyutNadi       int       `json:"denyut_nadi"`
	SuhuTubuh        float64   `json:"suhu_tubuh"`
	FrekuensiNapas   int       `json:"frekuensi_napas"`
	BeratBadan       float64   `json:"berat_badan"`
	TinggiBadan      float64   `json:"tinggi_badan"`
	SaturasiOksigen  int       `json:"saturasi_oksigen"`
	TingkatKesadaran string    `json:"tingkat_kesadaran"`
	IndeksMasaTubuh  float64   `json:"indeks_masa_tubuh"`
	Visible          int       `json:"visible"`
	NamaPasien       string    `json:"nama_pasien"`
	NamaDokter       string    `json:"nama_dokter"`
}

type PemeriksaanVitalWarehouseResponse struct {
	Source             string    `json:"source"`
	IDPemeriksaanVital int       `json:"id_pemeriksaan_vital"`
	IDPasien           int       `json:"id_pasien"`
	NamaPasien         *string   `json:"nama_pasien"`
	DateMake           time.Time `json:"date_make"`
	DateUpdate         time.Time `json:"date_update"`
	TekananDarah       string    `json:"tekanan_darah"`
	DenyutNadi         int       `json:"denyut_nadi"`
	SuhuTubuh          float64   `json:"suhu_tubuh"`
	FrekuensiNapas     int       `json:"frekuensi_napas"`
	BeratBadan         float64   `json:"berat_badan"`
	TinggiBadan        float64   `json:"tinggi_badan"`
	SaturasiOksigen    *int      `json:"saturasi_oksigen"`
	TingkatKesadaran   *string   `json:"tingkat_kesadaran"`
	IndeksMasaTubuh    *float64  `json:"indeks_masa_tubuh"`
	Visible            int       `json:"visible"`
}

type GetWarehouseResponse struct {
	Data []PemeriksaanVitalWarehouseResponse `json:"data"`
}
