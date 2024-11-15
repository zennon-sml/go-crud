package db

const (
  host = "localhost"
  port = 5432
  user = zennon
  password = "asdf"
  dbname = "gocrud"
)

func ConnectDB() (*sql.DB, error) {

dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

}
