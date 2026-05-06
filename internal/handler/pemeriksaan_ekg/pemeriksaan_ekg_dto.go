package pemeriksaan_ekg

type Response struct {
	ID                 int      `json:"id"`
	IDPasien           int      `json:"idPasien"`
	NamaPasien         string   `json:"namaPasien"`
	IDDokter           int      `json:"idDokter"`
	NamaDokter         string   `json:"namaDokter"`
	Visible            int      `json:"visible"`
	DetakJantung       *int     `json:"detakJantung"`
	Irama              string   `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	AxisJantung        string   `json:"axisJantung"`
	STElevationDepress string   `json:"st_elevation_depression"`
	TWaveAbnormality   string   `json:"t_wave_abnormality"`
	InterpretasiDokter string   `json:"interpretasi_dokter"`
	DateMake           string   `json:"dateMake"`
}

type ResponseB struct {
	Source             string   `json:"source"`
	ID                 int      `json:"id"`
	IDPasien           int      `json:"idPasien"`
	NamaPasien         string   `json:"namaPasien"`
	IDDokter           int      `json:"idDokter"`
	NamaDokter         string   `json:"namaDokter"`
	Visible            int      `json:"visible"`
	DetakJantung       *int     `json:"detakJantung"`
	Irama              string   `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	PWave              string   `json:"p_wave"`
	AVBlock            string   `json:"av_block"`
	InterpretasiDokter string   `json:"interpretasi_dokter"`
	DateMake           string   `json:"dateMake"`
	DateUpdate         string   `json:"dateUpdate"`
}

type WarehouseResponse struct {
	Source             string   `json:"source"`
	ID                 int      `json:"id"`
	IDPasien           int      `json:"idPasien"`
	NamaPasien         *string  `json:"namaPasien"`
	Visible            int      `json:"visible"`
	DetakJantung       *int     `json:"detakJantung"`
	Irama              string   `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	AxisJantung        string   `json:"axisJantung"`
	STElevationDepress string   `json:"st_elevation_depression"`
	TWaveAbnormality   string   `json:"t_wave_abnormality"`
	PWave              string   `json:"p_wave"`
	AVBlock            string   `json:"av_block"`
	InterpretasiDokter string   `json:"interpretasi_dokter"`
	DateMake           string   `json:"dateMake"`
	DateUpdate         string   `json:"dateUpdate"`
}

type GetWarehouseResponse struct {
	Data []WarehouseResponse `json:"data"`
}
