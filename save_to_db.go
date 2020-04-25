package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func save_to_db(images []Img) {
	var db_path = "./image_processing.db"
	os.Remove(db_path)

	database, _ := sql.Open("sqlite3", db_path)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS objects (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, img_id INTEGER , object_id TEXT, x_min INTEGER , x_max INTEGER, y_min INTEGER, y_max INTEGER)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO objects (img_id, object_id, x_min, x_max, y_min, y_max) VALUES (?, ?, ?, ?, ?, ?)")
	for i := 0; i < len(images); i++ {
		var img = images[i]
		for k := 0; k < len(images[i].Objects); k++ {
			var object = img.Objects[k]
			statement.Exec(img.ID, object.ID, object.BndBox.Xmin, object.BndBox.Xmax, object.BndBox.Ymin, object.BndBox.Ymax)
		}
	}

	//Printing data to check
	/**
	rows, _ := database.Query("SELECT id, img_id, object_id FROM objects")
	var id int
	var object_id string
	var img_id int

	for rows.Next() {
		rows.Scan(&id, &img_id, &object_id)
		fmt.Println(strconv.Itoa(id) + ": " + strconv.Itoa(img_id) + ":" + object_id)
	}
	*/
}
