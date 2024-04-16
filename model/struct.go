package model

type Task struct {
	IdTask    int    `gorm:"primaryKey;column:id_task" json:"id_task"`
	IdUser    int    `gorm:"column:id_user" json:"is_user"`
	Judul     string `gorm:"column:judul" json:"judul"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	DueDate   string `gorm:"column:due_date" json:"due_date"`
	Completed bool   `gorm:"column:completed" json:"completed"`
}
