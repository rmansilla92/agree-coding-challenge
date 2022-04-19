package services

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	ServicesTestSuite struct {
		suite.Suite
		services Services
		sql      SQLServices
	}
)

func TestServicesSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}

func (suite *ServicesTestSuite) SetupSuite() {
	fmt.Printf("Starting testing server")
	suite.sql = getTestingSQLService(suite.T())
}

func (suite *ServicesTestSuite) TearDownSuite() {
	fmt.Printf("Shutting down testing server")
}

func getTestingSQLService(t *testing.T) SQLServices {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	err = loadSQLFile(sqlDB, "../../../test/db.sql")
	if err != nil {
		t.Fatal(err)
	}
	return NewSQLServices(db)
}

func loadSQLFile(db *sql.DB, sqlFile string) error {
	file, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback()
	}()
	for _, q := range strings.Split(string(file), ";") {
		q := strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := tx.Exec(q); err != nil {
			return err
		}
	}
	return tx.Commit()
}