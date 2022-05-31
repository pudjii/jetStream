package config

import (
	"context"
	"os"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var (
	// DBURL database url
	DBURL string = os.Getenv("DIACONDBLNK")
	// DBUSER database user
	DBUSER string = os.Getenv("DIACONDBAS")
	// DBPASS database password
	DBPASS string = os.Getenv("DIACONDBASP")
	// DBNAME database name
	DBNAME   string = os.Getenv("DIACONDBUSE")
	TIMEZONE        = os.Getenv("TZ")
	Ctx             = context.Background()
)

type DBConn struct {
	DBLive driver.Client
}

func GetConnection() *DBConn {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBURL},
	})

	if err != nil {
		panic(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUSER, DBPASS),
	})

	if err != nil {
		panic(err)
	}

	return &DBConn{
		DBLive: client,
	}
}

func NewArangoDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}
