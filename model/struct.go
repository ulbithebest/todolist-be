package model

type Task struct {
	IdTask    int    `gorm:"primaryKey;column:id_task" json:"id_task"`
	Judul     string `gorm:"column:judul" json:"judul"`
	Deskripsi string `gorm:"column:deskripsi" json:"deskripsi"`
	DueDate   string `gorm:"column:due_date" json:"due_date"`
	Completed bool   `gorm:"column:completed" json:"completed"`
}
