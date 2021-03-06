package db

import (
	"fmt"
	"io"
)

var idCounter int

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	fmt.Printf("DBConnectionFactory : Creating resource (ID) : %d\n", idCounter)
	return &DBConnection{ID: idCounter}, nil
}
