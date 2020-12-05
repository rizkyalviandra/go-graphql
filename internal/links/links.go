package links

import (
	"log"
	database "go-graphql/internal/pkg/db/mysql"
	"go-graphql/internal/users"
)

// Link Struct
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

// Save Function
func (link Link) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Row Inserted!")
	return id
}

// GetAll from Link Function
func GetAll() []Link {
	stmt, err := database.Db.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err != nil {
		log.Fatal(err)
	}
	return links
}