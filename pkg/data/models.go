package data

type ResponseBasic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type ResponseBuku struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Buku
}

type Users struct {
	Id        int `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
}

type Buku struct {
	Id int `form:"id" json:"id"`
	Judul string `form:"judul" json:"judul"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Pencipta string `form:"pencipta" json:"pencipta"`
	Tahun string `form:"tahun" json:"tahun"`
}