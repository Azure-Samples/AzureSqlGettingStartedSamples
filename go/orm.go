package main

import (
	"log"
	"fmt"
	"database/sql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"github.com/microsoft/go-mssqldb/azuread"
)

var sqlDB *sql.DB

var server = "your_server.database.windows.net"
var database = "your_database"
var port = 1433

// Define a User model struct
type User struct {
	gorm.Model
	ID	  uint `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	FirstName string
	LastName  string
	Tasks	  []Task 
}

// Define a Task model struct
type Task struct {
	gorm.Model
	ID	   uint `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title      string
	DueDate    string
	IsComplete bool
	UserID     uint
}



// Read and print all the tasks
func ReadAllTasks(db *gorm.DB) {
	var users []User
	var tasks []Task
	db.Find(&users)

	for _, user := range users {
		db.Model(&user).Association("Tasks").Find(&tasks)
		fmt.Printf("%s %s's tasks:\n", user.FirstName, user.LastName)
		for _, task := range tasks {
			fmt.Printf("\tTitle: %s\n\tDueDate: %s\n\tIsComplete:%t\n\n",
				task.Title, task.DueDate, task.IsComplete)
		}
	}
}

// Update a task based on a user
func UpdateSomeonesTask(db *gorm.DB, userId int) {
	var task Task
	db.Where("user_id = ?", userId).First(&task).Update("Title", "Buy donuts for Luis")
	fmt.Printf("\tTitle: %s\n\tDueDate: %s\n\tIsComplete:%t\n\n",
		task.Title, task.DueDate, task.IsComplete)
}

// Delete all the tasks for a user
func DeleteSomeonesTasks(db *gorm.DB, userId int) {
	db.Where("user_id = ?", userId).Delete(&Task{})
	fmt.Printf("\nDeleted all tasks for user %d", userId)
}

func main()  {

    	var err error

 	// Build connection string
    	connString := fmt.Sprintf("server=%s;port=%d;database=%s;fedauth=ActiveDirectoryDefault;", server, port, database)

    	// Create SQL connection
        sqlDB, err = sql.Open(azuread.DriverName, connString)
    	if err != nil {
        	log.Fatal("Error connecting to database: ", err.Error())
    	}

	//Use the SQL connection to initialize *gorm.DB 
	db, err := gorm.Open(sqlserver.New(sqlserver.Config{Conn: sqlDB,}), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM failed to connect database: ", err.Error())
	}

	fmt.Println("Migrating models...")
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Task{})

	// Create awesome Users
	fmt.Println("Creating awesome users...")
	db.FirstOrCreate(&User{FirstName: "Andrea", LastName: "Lam"},&User{FirstName: "Andrea", LastName: "Lam"})   //UserID: 1
	db.FirstOrCreate(&User{FirstName: "Meet", LastName: "Bhagdev"},&User{FirstName: "Meet", LastName: "Bhagdev"}) //UserID: 2
	db.FirstOrCreate(&User{FirstName: "Luis", LastName: "Bosquez"},&User{FirstName: "Luis", LastName: "Bosquez"}) //UserID: 3

	// Create appropriate Tasks for each user
	fmt.Println("Creating new appropriate tasks...")
	db.FirstOrCreate(&Task{Title: "Do laundry", DueDate: "2021-03-30", IsComplete: false, UserID: 1},&Task{Title: "Do laundry", DueDate: "2021-03-30", IsComplete: false, UserID: 1}) 
	db.FirstOrCreate(&Task{Title: "Mow the lawn", DueDate: "2021-03-30", IsComplete: false, UserID: 2},&Task{Title: "Mow the lawn", DueDate: "2021-03-30", IsComplete: false, UserID: 2}) 
	db.FirstOrCreate(&Task{Title: "Do more laundry", DueDate: "2021-03-30", IsComplete: false, UserID: 3}, &Task{Title: "Do more laundry", DueDate: "2021-03-30", IsComplete: false, UserID: 3}) 
	db.FirstOrCreate(&Task{Title: "Watch TV", DueDate: "2021-03-30", IsComplete: false, UserID: 3}, &Task{Title: "Watch TV", DueDate: "2021-03-30", IsComplete: false, UserID: 3}) 

	// Read
	fmt.Println("\nReading all tasks...")
	ReadAllTasks(db)

	// Update - update Task title to something more appropriate
	fmt.Println("Updating Andrea's task...")
	UpdateSomeonesTask(db, 1)

	// Delete - delete Luis's task
	DeleteSomeonesTasks(db, 3)

	// Read
	fmt.Println("\nReading all tasks...")
	ReadAllTasks(db)
}
