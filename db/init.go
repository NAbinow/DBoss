package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

var DB *pgx.Conn

type Dummy []any

func Init_DB() {
	var err error
	DB, err = pgx.Connect(context.Background(), os.Getenv("PSQL_URL"))
	fmt.Println(os.Getenv("PSQL_URL"))
	if err != nil {
		fmt.Print(err)
		return
	}
}
