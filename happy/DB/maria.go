package db

import "database/sql"

func mariadb() {

	var variable string
	var setTime string

	db, err := sql.Open("mysql", "maiorem:pass@tcp(localhost:8089)/happysave")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.QueryRow("Select variable, set_time From happy_user").Scan(&variable, &setTime)

}
