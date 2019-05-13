package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getTeams() ([]team, error) {
	ctx := context.TODO()
	un := os.Getenv("LOCAL_MYSQL_UN")
	pw := os.Getenv	("LOCAL_MYSQL_PW")
	conn := fmt.Sprintf("%v:%v@(127.0.0.1:3306)/",un,pw)

	db,err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	rows,err := db.QueryContext(ctx,"SELECT name from test.team")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []team{}

	for rows.Next() {
		var t team
		err := rows.Scan(&t.name)
		if err != nil {
			return nil, err
		}

		teams = append(teams,t)
	}

	return teams, nil
}