package mysqlprovisioner

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hpcloud/cf-usb/driver/mysql/config"
	"github.com/pivotal-golang/lager"
)

type MysqlProvisioner struct {
	Conf       config.MysqlDriverConfig
	Connection *sql.DB
	logger     lager.Logger
}

func New(logger lager.Logger) MysqlProvisionerInterface {
	return &MysqlProvisioner{logger: logger}
}

func (e *MysqlProvisioner) Connect(conf config.MysqlDriverConfig) error {
	var err error
	e.Conf = conf

	e.Connection, err = e.openSqlConnection()

	return err
}

func (e *MysqlProvisioner) Close() error {
	err := e.Connection.Close()
	return err
}

func (e *MysqlProvisioner) IsDatabaseCreated(databaseName string) (bool, error) {
	rows, err := e.Query(fmt.Sprintf("SHOW DATABASES WHERE `database` = '%s'", databaseName))
	if err != nil {
		return false, err
	}

	var (
		result    [][]string
		container []string
		pointers  []interface{}
	)

	cols, err := rows.Columns()
	if err != nil {
		return false, err
	}

	length := len(cols)

	for rows.Next() {
		pointers = make([]interface{}, length)
		container = make([]string, length)

		for i := range pointers {
			pointers[i] = &container[i]
		}

		err = rows.Scan(pointers...)
		if err != nil {
			return false, err
		}

		result = append(result, container)
	}
	for _, cont := range result {
		if cont[0] == databaseName {
			return true, nil
		}
	}

	return false, nil
}

func (e *MysqlProvisioner) IsUserCreated(userName string) (bool, error) {
	rows, err := e.Query(fmt.Sprintf("SELECT user from mysql.user WHERE user = '%s'", userName))
	if err != nil {
		return false, err
	}

	var (
		result    [][]string
		container []string
		pointers  []interface{}
	)

	cols, err := rows.Columns()
	if err != nil {
		return false, err
	}

	length := len(cols)

	for rows.Next() {
		pointers = make([]interface{}, length)
		container = make([]string, length)

		for i := range pointers {
			pointers[i] = &container[i]
		}

		err = rows.Scan(pointers...)
		if err != nil {
			return false, err
		}

		result = append(result, container)
	}
	for _, cont := range result {
		if cont[0] == userName {
			return true, nil
		}
	}

	return false, nil
}

func (e *MysqlProvisioner) CreateDatabase(databaseName string) error {
	err := e.executeTransaction(e.Connection, fmt.Sprintf("CREATE DATABASE %s", databaseName))
	if err != nil {
		e.logger.Error("create database", err)
		return err
	}

	return nil
}

func (e *MysqlProvisioner) DeleteDatabase(databaseName string) error {

	err := e.executeTransaction(e.Connection, fmt.Sprintf("DROP DATABASE %s", databaseName))
	if err != nil {
		e.logger.Error("delete database", err)
		return err
	}
	return nil
}

func (e *MysqlProvisioner) Query(query string) (*sql.Rows, error) {

	result, err := e.Connection.Query(query)

	if err != nil {
		e.logger.Error("query", err)
		return nil, err
	}
	return result, nil
}

func (e *MysqlProvisioner) CreateUser(databaseName string, username string, password string) error {

	e.logger.Info("Connection open - executing transaction")
	err := e.executeTransaction(e.Connection,
		fmt.Sprintf("CREATE USER '%s' IDENTIFIED BY '%s';", username, password),
		fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';", databaseName, username),
		"FLUSH PRIVILEGES;")
	e.logger.Info("Transaction done")
	if err != nil {
		e.logger.Error("create user", err)
		return err
	}

	return nil
}

func (e *MysqlProvisioner) DeleteUser(username string) error {

	err := e.executeTransaction(e.Connection, fmt.Sprintf("DROP USER '%s'", username))

	if err != nil {
		e.logger.Error("delete user", err)
		return err
	}

	return nil
}

func (e *MysqlProvisioner) openSqlConnection() (*sql.DB, error) {
	con, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/mysql", e.Conf.User, e.Conf.Pass, e.Conf.Host, e.Conf.Port))
	if err != nil {
		return nil, err
	}
	return con, nil
}

func (e *MysqlProvisioner) executeTransaction(con *sql.DB, querys ...string) error {
	tx, err := con.Begin()
	if err != nil {
		e.logger.Error("execute transaction", err)
		return err
	} else {
		for _, query := range querys {
			e.logger.Info(query)
			_, err = tx.Exec(query)
			if err != nil {
				e.logger.Error("execute transaction query", err)
				tx.Rollback()
				break
			}
		}
		tx.Commit()
	}
	if err != nil {
		return err
	}
	return nil
}
