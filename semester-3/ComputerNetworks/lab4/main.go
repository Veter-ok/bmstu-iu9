package main

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB

func connectToDB() *sql.DB {
	dsn := "iu9networkslabs:Je2dTYr6@tcp(students.yss.su:3306)/iu9networkslabs"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func collectEmails() ([]string, []string) {
	rows, err := Database.Query("SELECT * FROM veterokEmails")
	if err != nil {
		rows.Close()
		panic(err)
	}
	defer rows.Close()
	var emails, names []string
	for rows.Next() {
		var email, name string
		if err := rows.Scan(&name, &email); err != nil {
			panic(err)
		}
		emails = append(emails, email)
		names = append(names, name)
	}
	return emails, names
}

func logMessage(email, message string) {
	now := time.Now()
	now = now.Add(3 * time.Hour)
	_, err := Database.Exec("INSERT INTO veterokLogs (dest, message, date) VALUES (?, ?, ?)", email, message, now)
	if err != nil {
		panic(err)
	}
}

func main() {
	Database = connectToDB()

	from := "veterok@veterok.123aaa.ru"
	password := "EffSZmzA95PDDjXY6K8B"

	to, names := collectEmails()
	fmt.Println(to)

	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	subject := "Лавров Родион ИУ9-32Б (Лабораторная работа №4)"

	// for {
	for i := 0; i < len(to); i++ {
		html := fmt.Sprintf("Здравствуйте, %s.\nРезультат лабораторной работы №4", names[i])
		message := fmt.Sprintf(
			"To: %s\r\n"+
				"Subject: %s\r\n"+
				"\r\n"+
				"%s\r\n",
			to[i],
			subject,
			html,
		)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to[i]}, []byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Письмо ушло, %s!\n", to[i])
		logMessage(to[i], html)
	}
	// 	time.Sleep(1 * time.Second)
	// }
}
