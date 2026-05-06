package pemeriksaan_laboratorium

type Response struct {
	ID           int      `json:"id"`
	IDPasien     int      `json:"idPasien"`
	NamaPasien   string   `json:"namaPasien"`
	IDDokter     int      `json:"idDokter"`
	NamaDokter   string   `json:"namaDokter"`
	Visible      int      `json:"visible"`
	Hb           *float64 `json:"hb"`
	Ht           *float64 `json:"ht"`
	Leukosit     *int     `json:"leukosit"`
	Trombosit    *int     `json:"trombosit"`
	GulaPuasa    *float64 `json:"gulaPuasa"`
	GulaSewaktu  *float64 `json:"gulaSewaktu"`
	HbA1c        *float64 `json:"hba1c"`
	Kolesterol   *float64 `json:"kolesterolTotal"`
	HDL          *float64 `json:"hdl"`
	LDL          *float64 `json:"ldl"`
	Trigliserida *float64 `json:"trigliserida"`
	SGOT         *float64 `json:"sgot"`
	SGPT         *float64 `json:"sgpt"`
	Ureum        *float64 `json:"ureum"`
	Kreatinin    *float64 `json:"kreatinin"`
	AsamUrat     *float64 `json:"asamUrat"`
	Natrium      *float64 `json:"natrium"`
	Kalium       *float64 `json:"kalium"`
	Klorida      *float64 `json:"klorida"`
	Waktu        string   `json:"waktu"`
}

type ResponseB struct {
	Source         string   `json:"source"`
	ID             int      `json:"id"`
	IDPasien       int      `json:"idPasien"`
	NamaPasien     string   `json:"namaPasien"`
	IDDokter       int      `json:"idDokter"`
	NamaDokter     string   `json:"namaDokter"`
	Visible        int      `json:"visible"`
	Hb             *float64 `json:"hb"`
	Ht             *float64 `json:"ht"`
	Leukosit       *int     `json:"leukosit"`
	Trombosit      *int     `json:"trombosit"`
	GulaPuasa      *float64 `json:"gulaPuasa"`
	GulaSewaktu    *float64 `json:"gulaSewaktu"`
	HbA1c          *float64 `json:"hba1c"`
	Kolesterol     *float64 `json:"kolesterolTotal"`
	HDL            *float64 `json:"hdl"`
	LDL            *float64 `json:"ldl"`
	Trigliserida   *float64 `json:"trigliserida"`
	SGOT           *float64 `json:"sgot"`
	SGPT           *float64 `json:"sgpt"`
	Ureum          *float64 `json:"ureum"`
	Kreatinin      *float64 `json:"kreatinin"`
	AsamUrat       *float64 `json:"asamUrat"`
	Natrium        *float64 `json:"natrium"`
	Kalium         *float64 `json:"kalium"`
	Klorida        *float64 `json:"klorida"`
	Waktu          string   `json:"waktu"`
	LajuEndapDarah *float64 `json:"lajuEndapDarah"`
	Albumin        *float64 `json:"albumin"`
}

type WarehouseResponse struct {
	Source         string   `json:"source"`
	ID             int      `json:"id"`
	IDPasien       int      `json:"idPasien"`
	NamaPasien     *string  `json:"namaPasien"`
	Hb             *float64 `json:"hb"`
	Ht             *float64 `json:"ht"`
	Leukosit       *int     `json:"leukosit"`
	Trombosit      *int     `json:"trombosit"`
	GulaPuasa      *float64 `json:"gulaPuasa"`
	GulaSewaktu    *float64 `json:"gulaSewaktu"`
	HbA1c          *float64 `json:"hba1c"`
	Kolesterol     *float64 `json:"kolesterolTotal"`
	HDL            *float64 `json:"hdl"`
	LDL            *float64 `json:"ldl"`
	Trigliserida   *float64 `json:"trigliserida"`
	SGOT           *float64 `json:"sgot"`
	SGPT           *float64 `json:"sgpt"`
	Ureum          *float64 `json:"ureum"`
	Kreatinin      *float64 `json:"kreatinin"`
	AsamUrat       *float64 `json:"asamUrat"`
	Natrium        *float64 `json:"natrium"`
	Kalium         *float64 `json:"kalium"`
	Klorida        *float64 `json:"klorida"`
	Waktu          string   `json:"waktu"`
	LajuEndapDarah *float64 `json:"lajuEndapDarah"`
	Albumin        *float64 `json:"albumin"`
	Visible        int      `json:"visible"`
}

type GetWarehouseResponse struct {
	Data []WarehouseResponse `json:"data"`
}
