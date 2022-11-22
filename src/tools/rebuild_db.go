package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"boilerplate/database"
	"boilerplate/utils"
)

func main() {
	ctx := context.Background()

	dbc, err := database.Connect(ctx)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	log.Println("connected to db")

	err = database.DropSchemas(ctx, dbc.Master)
	if err != nil {
		log.Fatalf("error rebuilding schema: %v", err)
	}
	log.Println("rebuilt schemas")

	err = filepath.Walk(utils.ProjectPath(),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(path) != ".sql" {
				return nil
			}

			log.Printf("executing queries in %s", path)

			err = database.ExecuteSQLFile(ctx, dbc.Master, path)
			if err != nil {
				err = fmt.Errorf("error executing "+
					"queries from sql file: %v", err)
			}

			return err
		})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db rebuilt successfully")
}
