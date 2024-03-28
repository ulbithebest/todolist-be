package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/todolist-be/model"
	repo "github.com/ulbithebest/todolist-be/repository"
	"gorm.io/gorm"
	"net/http"
)

func GetAllTask(c *fiber.Ctx) error {
	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk mendapatkan semua task
	tasks, err := repo.GetAllTask(db)
	if err != nil {
		// Jika terjadi kesalahan saat mengambil data task, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada task yang ditemukan, mengembalikan pesan kesalahan
	if len(tasks) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"code": http.StatusNotFound, "success": false, "status": "error", "message": "Data task tidak ditemukan", "data": nil})
	}

	// Jika tidak ada kesalahan, mengembalikan data task sebagai respons JSON
	response := fiber.Map{
		"code":    http.StatusOK,
		"success": true,
		"status":  "success",
		"data":    tasks,
	}

	return c.Status(http.StatusOK).JSON(response)
}

func GetTaskById(c *fiber.Ctx) error {
	// Mendapatkan parameter ID task dari URL
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk mendapatkan task berdasarkan ID
	task, err := repo.GetTaskById(db, id)
	if err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Jika tidak ada kesalahan, mengembalikan data task sebagai respons JSON
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "data": task})
}

func InsertTask(c *fiber.Ctx) error {
	// Mendeklarasikan variabel untuk menyimpan data task dari body request
	var task model.Task

	// Mem-parsing body request ke dalam variabel task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk menyisipkan data task ke dalam database
	if err := repo.InsertTask(db, task); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan task"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.Status(http.StatusCreated).JSON(fiber.Map{"code": http.StatusCreated, "success": true, "status": "success", "message": "Buku berhasil disimpan", "data": task})
}

func UpdateTask(c *fiber.Ctx) error {
	// Mendapatkan parameter ID
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	// Mendeklarasikan variabel untuk menyimpan data task yang diperbarui dari body request
	var updatedTask model.Task

	// Mem-parsing body request ke dalam variabel updatedTask
	if err := c.BodyParser(&updatedTask); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Gagal memproses request"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk memperbarui data task di dalam database
	if err := repo.UpdateTask(db, id, updatedTask); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memperbarui task"})
	}

	// Mengembalikan respons sukses dengan pesan
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Task berhasil diperbarui"})
}

func DeleteTask(c *fiber.Ctx) error {
	// Mendapatkan parameter ID task dari URL
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ID task tidak ditemukan"})
	}

	// Mendapatkan koneksi database dari context Fiber
	db := c.Locals("db").(*gorm.DB)

	// Memanggil fungsi repo untuk menghapus data task dari database berdasarkan ID
	if err := repo.DeleteTask(db, id); err != nil {
		// Jika terjadi kesalahan, mengembalikan respons error
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus task"})
	}

	// Mengembalikan respons sukses
	return c.JSON(fiber.Map{"code": http.StatusOK, "success": true, "status": "success", "message": "Task berhasil dihapus", "deleted_id": id})
}