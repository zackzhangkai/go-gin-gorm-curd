package main

import (
	initialenviroment "go-gin-web/initialEnviroment"
	"go-gin-web/models"
)

func init() {
	initialenviroment.EnvInit()
	initialenviroment.DbInit()
}

func main() {
	initialenviroment.DB.AutoMigrate(&models.Blog{})
	// Create table for `User`
	// initialEnviroment.DB.Migrator().CreateTable(&User{})

	// // Append "ENGINE=InnoDB" to the creating table SQL for `User`
	// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&User{})

	// // Check table for `User` exists or not
	// db.Migrator().HasTable(&User{})
	// db.Migrator().HasTable("users")

	// // Drop table if exists (will ignore or delete foreign key constraints when dropping)
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable("users")

	// // Rename old table to new table
	// db.Migrator().RenameTable(&User{}, &UserInfo{})
	// db.Migrator().RenameTable("users", "user_infos")
}
