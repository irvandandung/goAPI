package data

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface {}
}

type Users struct {
	Id        float64 `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Role  string `form:"role" json:"role"`
}

type Buku struct {
	Id int `form:"id" json:"id"`
	Judul string `form:"judul" json:"judul"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Pencipta string `form:"pencipta" json:"pencipta"`
	Tahun string `form:"tahun" json:"tahun"`
}