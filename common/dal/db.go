package dal

import (
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var database *pg.DB

// Database returns a singleton pg.DB instance. If the application cannot
// connect to the database, it panics
func Database() *pg.DB {
	if database == nil {
		if db, err := initDB(); err != nil {
			panic(err)
		} else {
			if err = createSchema(db); err != nil {
				panic(err)
			}

			database = db
		}
	}

	return database
}

// Shutdown closes the connection to the database which should be called
// during application shutdown
func Shutdown() error {
	if database != nil {
		return database.Close()
	}

	return nil
}

func initDB() (*pg.DB, error) {
	options, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return pg.Connect(options), nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Project)(nil), (*Build)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}

		if val, isType := model.(Indexed); isType {
			for _, index := range val.Indexes() {
				if _, err = db.Model(model).Exec(index); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
