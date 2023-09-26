package structs

type Mahasiswa struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
