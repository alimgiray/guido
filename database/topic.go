package database

import (
	"database/sql"
	"log"
	"time"
)

const createTopicTableQuery = `
	CREATE TABLE topic (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"title" TEXT NOT NULL,
		"url" TEXT NOT NULL,
		"user" INTEGER NOT NULL,
		"createdAt" TEXT NOT NULL,
		"updatedAt" TEXT,
		FOREIGN KEY(user) REFERENCES user(id)
	);
`
const createPostTableQuery = `
	CREATE TABLE post (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"text" TEXT NOT NULL,
		"user" INTEGER NOT NULL,
		"topic" INTEGER NOT NULL,
		"createdAt" TEXT NOT NULL,
		"updatedAt" TEXT,
		FOREIGN KEY(user) REFERENCES user(id),
		FOREIGN KEY(topic) REFERENCES topic(id)
	);
`

func createDefaultTopic(db *sql.DB) {
	topic := "Guido"
	url := "guido"
	post := "Your first post"
	topicQuery := `INSERT INTO topic(title, url, user, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)`
	postQuery := `INSERT INTO post(text, user, topic, createdAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(topicQuery)
	_, err = statement.Exec(topic, url, 1, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
	}
	statement, err = db.Prepare(postQuery)
	_, err = statement.Exec(post, 1, 1, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
	}
}
