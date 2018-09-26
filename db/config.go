package db

type Conf struct {
	username     string
	password     string
	host         string
	port         string
	dbName       string
	errSleepTime int
}

var (
	psql = Conf{
		username:     "postgres",
		password:     "postgres",
		host:         "db",
		port:         "5432",
		dbName:       "tax_calc",
		errSleepTime: 10000,
	}
)
