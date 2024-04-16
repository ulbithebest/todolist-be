package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/ulbithebest/todolist-be/model"
	repo "github.com/ulbithebest/todolist-be/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var user model.Users

	db := c.Locals("db").(*gorm.DB)

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body Invalid",
		})
	}

	// Menyimpan user ke database menggunakan repository
	if err := repo.CreateUser(db, &user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Register",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Berhasil Register!",
	})
}

func LoginUser(c *fiber.Ctx) error {
	var user model.Users

	db := c.Locals("db").(*gorm.DB)

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Request Body Invalid",
		})
	}

	// Cari user di database berdasarkan username menggunakan repository
	userData, err := repo.GetUserByUsername(db, user.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau Password Salah",
		})
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Username atau Password Salah",
		})
	}

	// Buat JWT token menggunakan repository
	token, err := repo.GenerateToken(userData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal Mendapatkan Token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func GetMe(c *fiber.Ctx) error {
	// Mendapatkan data user yang sedang login melalui JWT token
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*model.JWTClaims)
	db := c.Locals("db").(*gorm.DB)

	// Cari user di database berdasarkan user ID menggunakan repository
	userData, err := repo.GetUserById(db, claims.IdUser)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"user": userData,
	})
}

func Authenticate(c *fiber.Ctx) error {
	// Mendapatkan token dari header Authorization
	authHeader := c.Get("Authorization")
	token := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Header Otorisasi Salah",
		})
	}

	// Verifikasi token
	claims := new(model.JWTClaims)
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret_key"), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token Invalid",
			})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Gagal Mengautentikasi Token",
		})
	}

	if !tkn.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token Salah",
		})
	}

	// Menyimpan data user ke local context
	c.Locals("user", tkn)

	return c.Next()
}

func LogoutUser(c *fiber.Ctx) error {
	// Hapus token dari Authorization header
	c.Set("Authorization", "")

	return c.JSON(fiber.Map{
		"message": "Logout berhasil",
	})
}
