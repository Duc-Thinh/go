package main

import (
	// "fmt"
	"log"
	// "os"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

type Note struct {
	Id int  `json:"id,omitempty" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
	Has_fineshed bool `json:"has_finished" gorm:"column:has_finished;"`
	Status int `json:"status" gorm:"column:status;"`
}

type NoteUpdate struct {
	Status *int `json:"status" gorm:"column:status;"`
}

func (Note) TableName() string {
	return "NOTES"
}

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":3000")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// dsn := os.Getenv("DBConnectionStr")
  	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// create := make(chan *Note)

	// for i := 0; i < 10; i++ {
	// 	go func(j int){
	// 		j++
	// 		h := j
	// 		h++
	// 		data := &Note{Id: j, Title: "Demo note", Content: "This is content", Has_fineshed: true, Status: j}
	// 		create <- data
	// 	}(i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if err := db.Create(<-create); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }



	// newNote := Note{Id: 3, Title: "Demo note", Content: "This is content", Has_fineshed: true, Status: 1}
	// if err := db.Create(&newNote); err != nil {
	// 	fmt.Println(err)
	// }

	// var notes []Note
	// db.Where("Title = ?", "Demo note").Find(&notes)
	// fmt.Println(notes)
	// var note Note
	// if err := db.Where("status = 2").First(&note); err != nil {
	// 	log.Println(err)
	// }

	// newUpdate := 1
	// db.Table(Note{}.TableName()).Where("status = 1").Updates(&NoteUpdate{Status: &newUpdate})

	// db.Table(Note{}.TableName()).Where("status = 1").Updates(map[string]interface{}{
	// 	"title": "Demo 2",
	// })
	// fmt.Println(note)
}