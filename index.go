package main

import (
	"database/sql"
	"fmt"

	// "fmt"
	// "golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

type PalletTracking struct {
	PalletTrackingId   int
	PalletId           int
	PalletLocation     string
	PalletType         string
	Remarks            sql.NullString
	IsDeleted          int
	CreatedOn          time.Time
	CreatedByUserGuid  string
	CreatedByUserName  string
	ModifiedOn         sql.NullTime
	ModifiedByUserGuid sql.NullString
	ModifiedByUserName sql.NullString
}

var db *sql.DB

func main() {
	connString := "localhost;user id=hamza;password=6761795h;port=63122;database=todo"
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	// -------------------------------------------------------------------------------------------------------------
	// *********************************  Pallet Tracking Related Stuff  ******************************************************
	// -------------------------------------------------------------------------------------------------------------

	r.POST("/pallet-trackings", createPalletTracking)
	r.GET("/pallet-trackings/:id", getPalletTrackingByID)
	r.GET("/pallet-trackings", getAllPalletTrackings)

	// -------------------------------------------------------------------------------------------------------------
	// *********************************  Pallet Related Stuff  ******************************************************
	// -------------------------------------------------------------------------------------------------------------

	type Pallet struct {
		PalletId           int
		PalletNo           string
		Area               string
		IsActive           int
		IsLocked           int
		Status             int
		IsDeleted          int
		CreatedOn          time.Time
		CreatedByUserGuid  string
		CreatedByUserName  string
		ModifiedOn         sql.NullTime
		ModifiedByUserGuid sql.NullString
		ModifiedByUserName sql.NullString
	}

	r.POST("/pallets", func(c *gin.Context) {
		var pallet Pallet
		if err := c.ShouldBindJSON(&pallet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invali Request Payload"})
			return
		}
		// Insert pallet logic here
		_, err := db.Exec("INSERT INTO Pallet (PalletId,PalletNo, Area, IsActive, IsLocked, Status, IsDeleted, CreatedOn, CreatedByUserGuid, CreatedByUserName) VALUES (?,?, ?, ?, ?, ?, ?, GETDATE(), ?, ?)",
			pallet.PalletId, pallet.PalletNo, pallet.Area, pallet.IsActive, pallet.IsLocked, pallet.Status, pallet.IsDeleted, pallet.CreatedByUserGuid, pallet.CreatedByUserName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create pallet: %s", err.Error())})
			return
		}
		c.JSON(http.StatusOK, gin.H{"sucess": "You Created Pallet Sucessfully", "pallet": pallet})
	})

	//Get Api

	r.GET("/pallets", func(c *gin.Context) {
		var pallets []Pallet
		rows, err := db.Query("SELECT * FROM Pallet")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pallets"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var pallet Pallet
			err := rows.Scan(
				&pallet.PalletId, &pallet.PalletNo, &pallet.Area, &pallet.IsActive, &pallet.IsLocked,
				&pallet.Status, &pallet.IsDeleted, &pallet.CreatedOn, &pallet.CreatedByUserGuid,
				&pallet.CreatedByUserName, &pallet.ModifiedOn, &pallet.ModifiedByUserGuid,
				&pallet.ModifiedByUserName,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pallets"})
				return
			}

			pallets = append(pallets, pallet)
		}
		c.JSON(http.StatusOK, pallets)

	})

	//
	r.PUT("/pallets/:id", func(c *gin.Context) {
		id := c.Param("id")
		var pallet Pallet
		if err := c.ShouldBindJSON(&pallet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Update pallet logic here
		_, err := db.Exec("UPDATE Pallet SET PalletNo=?, Area=?, IsActive=?, IsLocked=?, Status=?, IsDeleted=?, ModifiedOn=GETDATE(), ModifiedByUserGuid=?, ModifiedByUserName=? WHERE PalletId=?",
			pallet.PalletNo, pallet.Area, pallet.IsActive, pallet.IsLocked, pallet.Status, pallet.IsDeleted, pallet.ModifiedByUserGuid, pallet.ModifiedByUserName, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pallet"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Pallet updated successfully"})
	})

	//delete option
	r.DELETE("/pallets/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Delete pallet logic here
		_, err := db.Exec("DELETE FROM Pallet WHERE PalletId=?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pallet"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Pallet deleted successfully"})
	})

	// -------------------------------------------------------------------------------------------------------------
	// *********************************  Users Related Stuff  ******************************************************
	// -------------------------------------------------------------------------------------------------------------

	type User struct {
		LoginName          string
		DisplayName        string
		Password           string
		UserType           int
		IsLocked           int
		Status             int
		IsDeleted          int
		CreatedOn          string
		CreatedByUserName  string
		CreatedByUserGuid  string
		ModifiedOn         sql.NullString
		ModifiedByUserName sql.NullString
		ModifiedByUserGuid sql.NullString
	}

	// Implement the GET route to fetch all users
	r.GET("/users", func(c *gin.Context) {

		users := []User{} // Define the slice to hold user records

		rows, err := db.Query("SELECT * FROM [User]")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			err := rows.Scan(
				&user.LoginName, &user.DisplayName, &user.Password, &user.UserType, &user.IsLocked,
				&user.Status, &user.IsDeleted, &user.CreatedOn, &user.CreatedByUserName,
				&user.CreatedByUserGuid, &user.ModifiedOn, &user.ModifiedByUserName,
				&user.ModifiedByUserGuid,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
				return
			}
			users = append(users, user) // Add user to the slice
		}

		c.JSON(http.StatusOK, users) // Return the users as JSON array
	})

	// POST USER
	r.POST("/users", func(c *gin.Context) {
		var user struct {
			LoginName          string
			DisplayName        string
			Password           string
			UserType           int
			IsLocked           int
			Status             int
			IsDeleted          int
			CreatedOn          string
			CreatedByUserName  string
			CreatedByUserGuid  string
			ModifiedOn         sql.NullString
			ModifiedByUserName sql.NullString
			ModifiedByUserGuid sql.NullString
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert user logic here
		_, err := db.Exec("INSERT INTO [User] (LoginName, DisplayName, Password, UserType, IsLocked, Status, IsDeleted, CreatedOn, CreatedByUserName, CreatedByUserGuid, ModifiedOn, ModifiedByUserName, ModifiedByUserGuid) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			user.LoginName,
			user.DisplayName,
			user.Password,
			user.UserType,
			user.IsLocked,
			user.Status,
			user.IsDeleted,
			user.CreatedOn,
			user.CreatedByUserName,
			user.CreatedByUserGuid,
			user.ModifiedOn.String,
			user.ModifiedByUserName.String,
			user.ModifiedByUserGuid.String,
		)

		if err != nil {
			log.Printf("Error inserting user: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	r.GET("/users/:loginName", func(c *gin.Context) {
		loginName := c.Param("loginName")

		var user struct {
			DisplayName string
		}

		err := db.QueryRow("SELECT DisplayName FROM [User] WHERE LoginName = ?", loginName).Scan(&user.DisplayName)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:loginName", func(c *gin.Context) {
		loginName := c.Param("loginName")

		var userUpdate struct {
			DisplayName string
		}

		if err := c.ShouldBindJSON(&userUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE [User] SET DisplayName = ? WHERE LoginName = ?", userUpdate.DisplayName, loginName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	})

	r.DELETE("/users/:loginName", func(c *gin.Context) {
		loginName := c.Param("loginName")
		_, err := db.Exec("DELETE FROM [User] WHERE LoginName = ?", loginName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	r.Run(":8080") // Run the server on port 8080
}

// -------------------------------------------------------------------------------------------------------------
// *********************************  Tracking Route functions  ******************************************************
// -------------------------------------------------------------------------------------------------------------

func createPalletTracking(c *gin.Context) {
	var palletTracking PalletTracking
	if err := c.ShouldBindJSON(&palletTracking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := db.Exec("INSERT INTO PalletTracking (...) VALUES (...)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create pallet tracking: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pallet tracking created successfully", "palletTracking": palletTracking})
}

// func getPalletTrackingByID(c *gin.Context) {
// 	id := c.Param("id")
// 	var palletTracking PalletTracking
// 	err := db.QueryRow("SELECT * FROM PalletTracking WHERE PalletTrackingId = ?", id).Scan(
// 		&palletTracking.PalletTrackingId, &palletTracking.PalletId, &palletTracking.PalletLocation,
// 		&palletTracking.PalletType, &palletTracking.Remarks, &palletTracking.IsDeleted,
// 		&palletTracking.CreatedOn, &palletTracking.CreatedByUserGuid,
// 		&palletTracking.CreatedByUserName, &palletTracking.ModifiedOn,
// 		&palletTracking.ModifiedByUserGuid, &palletTracking.ModifiedByUserName,
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Pallet tracking not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, palletTracking)
// }

func getAllPalletTrackings(c *gin.Context) {
	var palletTrackings []PalletTracking

	rows, err := db.Query("SELECT * FROM PalletTracking")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pallet trackings"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var palletTracking PalletTracking
		err := rows.Scan(
			&palletTracking.PalletTrackingId, &palletTracking.PalletId, &palletTracking.PalletLocation,
			&palletTracking.PalletType, &palletTracking.Remarks, &palletTracking.IsDeleted,
			&palletTracking.CreatedOn, &palletTracking.CreatedByUserGuid,
			&palletTracking.CreatedByUserName, &palletTracking.ModifiedOn,
			&palletTracking.ModifiedByUserGuid, &palletTracking.ModifiedByUserName,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pallet trackings"})
			return
		}
		palletTrackings = append(palletTrackings, palletTracking)
	}

	c.JSON(http.StatusOK, palletTrackings)
}
