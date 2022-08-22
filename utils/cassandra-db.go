// connect go to cassandra db
package utils

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "restfulapi"
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Connected to Cassandra")
}
