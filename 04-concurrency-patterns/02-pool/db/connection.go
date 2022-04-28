package db

import "fmt"

//Resource
type DBConnection struct {
	ID int
}

func (dbConnection *DBConnection) Close() error {
	fmt.Printf("Closing and discarding the resource (ID) : %d\n", dbConnection.ID)
	return nil
}
