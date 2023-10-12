package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type Contact struct {
	Id      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	PhoneNo string `json:"phone_no" xml:"phone_no"`
}

type Contacts struct {
	Contacts []Contact `json:"contacts" xml:"contacts"`
}

type Response struct {
	Status  string `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
	Data    any    `json:"data" xml:"data"`
}

// constructor function
func (std *Response) fill_defaults() {

	// setting default values
	// if no values present
	if std.Status == "" {
		std.Status = ""
	}

	if std.Message == "" {
		std.Message = ""
	}

	if std.Data == "" {
		std.Data = ""
	}
}

func ContactGetList() echo.HandlerFunc {

	return func(c echo.Context) error {

		// Mengganti dengan informasi koneksi MySQL Anda
		dsn := "root:@tcp(127.0.0.1:3306)/crud"

		// Membuka koneksi ke database MySQL
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		defer db.Close()

		// Menguji koneksi ke database
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		fmt.Println("Koneksi sukses!")

		// Membuat tabel jika belum ada
		createTable := `
		CREATE TABLE IF NOT EXISTS contacts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50),
			phone_no VARCHAR(15)
		);
	`
		_, err = db.Exec(createTable)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}

		// Mengambil data dari tabel
		rows, err := db.Query("SELECT id, name, phone_no FROM contacts ORDER BY name ASC")
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		defer rows.Close()

		var contacts Contacts

		for rows.Next() {
			var id int
			var name string
			var phoneNo string
			err := rows.Scan(&id, &name, &phoneNo)
			if err != nil {
				log.Fatal(err)
				return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
			}
			fmt.Printf("ID: %d, Name: %s, Phone Number: %s\n", id, name, phoneNo)

			contacts.Contacts = append(contacts.Contacts, Contact{id, name, phoneNo})
		}

		return c.JSON(http.StatusOK, Response{"success", "Success", contacts})
	}
}

func ContactAdd() echo.HandlerFunc {
	return func(c echo.Context) error {
		var contact Contact
		contact.Name = c.FormValue("name")
		contact.PhoneNo = c.FormValue("phone_no")

		if contact.Name == "" {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Name is required"})
		}

		if contact.PhoneNo == "" {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Phone is required"})
		}

		// Mengganti dengan informasi koneksi MySQL Anda
		dsn := "root:@tcp(127.0.0.1:3306)/crud"

		// Membuka koneksi ke database MySQL
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		defer db.Close()

		// Menguji koneksi ke database
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		fmt.Println("Koneksi sukses!")

		// Membuat tabel jika belum ada
		createTable := `
		CREATE TABLE IF NOT EXISTS contacts (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50),
			phone_no VARCHAR(15)
		);
	`
		_, err = db.Exec(createTable)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}

		// Menyisipkan data ke dalam tabel
		insertContact := "INSERT INTO contacts (name, phone_no) VALUES (?, ?)"
		result, err := db.Exec(insertContact, contact.Name, contact.PhoneNo)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}

		// Mengambil ID data yang baru saja disisipkan
		lastID, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID data baru: %d\n", lastID)

		return c.JSON(http.StatusOK, Response{"error", "Data added!", ""})
	}
}

func ContactUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		i, err := strconv.Atoi(c.FormValue("id"))

		if err != nil {
			// ... handle error
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Invalid ID"})
		}

		var contact Contact
		contact.Id = i
		contact.Name = c.FormValue("name")
		contact.PhoneNo = c.FormValue("phone_no")

		if contact.Id == 0 {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Id is required"})
		}

		if contact.Name == "" {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Name is required"})
		}

		if contact.PhoneNo == "" {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Phone is required"})
		}

		// Mengganti dengan informasi koneksi MySQL Anda
		dsn := "root:@tcp(127.0.0.1:3306)/crud"

		// Membuka koneksi ke database MySQL
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		defer db.Close()

		// Menguji koneksi ke database
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		fmt.Println("Koneksi sukses!")

		// Menyisipkan data ke dalam tabel
		insertContact := "UPDATE contacts SET name=?, phone_no=? WHERE id=?"
		result, err := db.Exec(insertContact, contact.Name, contact.PhoneNo, contact.Id)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}

		return c.JSON(http.StatusOK, Response{"error", "Data updated!", result})
	}
}

func ContactDelete() echo.HandlerFunc {
	return func(c echo.Context) error {

		i, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			// ... handle error
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Invalid ID"})
		}

		var contact Contact
		contact.Id = i

		if contact.Id == 0 {
			return c.JSON(http.StatusBadRequest, Response{Status: "warning", Message: "Id is required"})
		}

		// Mengganti dengan informasi koneksi MySQL Anda
		dsn := "root:@tcp(127.0.0.1:3306)/crud"

		// Membuka koneksi ke database MySQL
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		defer db.Close()

		// Menguji koneksi ke database
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}
		fmt.Println("Koneksi sukses!")

		// Menyisipkan data ke dalam tabel
		insertContact := "DELETE FROM contacts WHERE id=?"
		result, err := db.Exec(insertContact, contact.Id)
		if err != nil {
			log.Fatal(err)
			return c.JSON(http.StatusInternalServerError, Response{Status: "warning", Message: string(err.Error())})
		}

		return c.JSON(http.StatusOK, Response{"error", "Data deleted!", result})
	}
}
