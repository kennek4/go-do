package util

import (
	"fmt"
)

// Custom Psuedo Enum
type DeleteType int

const (
	title = iota + 1
	id
)

func getQueryString(queryType DeleteType, dbTable string, thingToDelete *interface{}) string {

	query := fmt.Sprintf("DELETE FROM %s", dbTable)

	switch queryType {
	case title:
		query = fmt.Sprintf("%s WHERE title='%s'", query, *thingToDelete)
	default:
		query = fmt.Sprintf("%s WHERE id=%d", query, *thingToDelete)
	}

	return query
}

func DeleteTaskInDB(queryType DeleteType, dbTable string, thingToDelete *interface{}, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := getQueryString(queryType, dbTable, thingToDelete)
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	statement.Exec() // Exec query

	return nil // Task successfully deleted
}
