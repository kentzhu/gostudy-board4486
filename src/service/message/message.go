package message

import (
	"board4486/service/database"
	"log"
)

type Message struct {
	Id       string `json:"id" db:"id"`
	Datetime string `json:"datetime" db:"ctime"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	Author   string `json:"author" db:"author"`
}

func List() *[]Message {
	rows, err := database.GetDB().Query("SELECT id,ctime,title,content,author FROM message ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}

	var messageList = make([]Message, 0)

	for rows.Next() {
		var message Message
		err = rows.Scan(&message.Id, &message.Datetime, &message.Title, &message.Content, &message.Author)
		if err != nil {
			log.Fatal(err)
		}
		messageList = append(messageList, message)
	}
	defer rows.Close()

	return &messageList
}

func Insert(newMessage *Message) error {
	_, err := database.GetDB().Exec(
		"INSERT INTO message(title,content,author) values (?,?,?)",
		newMessage.Title, newMessage.Content, newMessage.Author,
	)
	return err
}

func Delete(messageId *int) error {
	_, err := database.GetDB().Exec("DELETE FROM message WHERE id=?", messageId)
	return err
}
