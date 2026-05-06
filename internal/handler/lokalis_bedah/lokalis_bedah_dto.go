package lokalis_bedah

type LokalisBedahResponse struct {
	ID             int    `json:"id"`
	IDPasien       int    `json:"id_pasien"`
	IDDokter       int    `json:"id_dokter"`
	Tanggal        string `json:"tanggal"`
	LokasiKelainan string `json:"lokasiKelainan"`
	JenisKelainan  string `json:"jenisKelainan"`
	Ukuran         string `json:"ukuran"`
	Warna          string `json:"warna"`
	NyeriTekan     bool   `json:"nyeriTekan"`
	Konsistensi    string `json:"konsistensi"`
	Mobilitas      string `json:"mobilitas"`
	TandaRadang    bool   `json:"tandaRadang"`
	Fluktuasi      bool   `json:"fluktuasi"`
	Catatan        string `json:"catatan"`
}

type LokalisBedahBResponse struct {
	Source         string `json:"source"`
	ID             int    `json:"id"`
	IDPasien       int    `json:"id_pasien"`
	IDDokter       int    `json:"id_dokter"`
	Tanggal        string `json:"tanggal"`
	LokasiKelainan string `json:"lokasiKelainan"`
	JenisKelainan  string `json:"jenisKelainan"`
	Ukuran         string `json:"ukuran"`
	Warna          string `json:"warna"`
	NyeriTekan     bool   `json:"nyeriTekan"`
	Konsistensi    string `json:"konsistensi"`
	Mobilitas      string `json:"mobilitas"`
	TandaRadang    bool   `json:"tandaRadang"`
	Fluktuasi      bool   `json:"fluktuasi"`
	Catatan        string `json:"catatan"`
}

type LokalisBedahWarehouseResponse struct {
	Source         string  `json:"source"`
	NamaPasien     *string `json:"nama_pasien"`
	ID             int     `json:"id"`
	IDPasien       int     `json:"id_pasien"`
	IDDokter       int     `json:"id_dokter"`
	Tanggal        string  `json:"tanggal"`
	LokasiKelainan string  `json:"lokasiKelainan"`
	JenisKelainan  string  `json:"jenisKelainan"`
	Ukuran         string  `json:"ukuran"`
	Warna          string  `json:"warna"`
	NyeriTekan     bool    `json:"nyeriTekan"`
	Konsistensi    string  `json:"konsistensi"`
	Mobilitas      string  `json:"mobilitas"`
	TandaRadang    bool    `json:"tandaRadang"`
	Fluktuasi      bool    `json:"fluktuasi"`
	Catatan        string  `json:"catatan"`
}

type GetWarehouseResponse struct {
	Data []LokalisBedahWarehouseResponse `json:"data"`
}
