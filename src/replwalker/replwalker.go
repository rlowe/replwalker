package replwalker

import (
	"database/sql"
  "fmt"
	//"github.com/go-sql-driver/mysql"
)

type MySQLServer struct {
	Cxn      *sql.DB
	Hostname string
	Master   *MySQLServer
	Slaves   []*MySQLServer
	Errors   []error
}

type MySQLCluster struct {
	Members []MySQLServer
}

func getMaster(host *MySQLServer) {
	defer host.Cxn.Close()
	err := host.Cxn.Ping()
	if err != nil {
		host.Errors = append(host.Errors, err)
		return
	}

	rows, err := host.Cxn.Query("SHOW SLAVE STATUS")
	if err != nil {
		host.Errors = append(host.Errors, err)
		return
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		host.Errors = append(host.Errors, err)
		return
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
		  host.Errors = append(host.Errors, err)
		  return
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
      if columns[i] == "Master_Host" {
        host.Master.Hostname = value
      }
		}
	}
	if err = rows.Err(); err != nil {
	  host.Errors = append(host.Errors, err)
	  return
	}
  return
}

func getSlaves(host *MySQLServer) {
  return
}

func GetLimitedClusterInfo(host MySQLServer, levels int) MySQLCluster {
  fmt.Println("heyo")
  return
}

func GetClusterInfo(host MySQLServer) MySQLCluster {
  fmt.Println("heyo")
  return
}
