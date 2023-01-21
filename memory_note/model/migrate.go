package model

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).AutoMigrate(&Task{})
	//DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")//mainkey
	if err != nil {
		return
	}
}
