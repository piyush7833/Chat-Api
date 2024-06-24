package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

var Db *sql.DB

var err error

func ConnectDb() {
	Db, err = sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err, "error connecting to db")
	}

	// Ping the database to check if the connection is successful
	err = Db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to the database")

}

// CloseDb closes the database connection
func DisconnectDb() {
	Db.Close()
}

// GetRows fetches rows from a specified table with pagination and selected columns
func GetRows(table string, page int, selectedColumns []string, validColumns map[string]bool, where *string, joins *string, joinColumns map[string]string, orderBy []string, isDesc bool) ([]map[string]interface{}, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout*time.Second)
	defer cancel()
	columns, orderBy, err := helpers.ValidateColumns(selectedColumns, validColumns, joinColumns, orderBy)
	if err != nil {
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	query, offset := helpers.ConstructGetQuery(table, columns, page, where, joins, orderBy, isDesc)
	// fmt.Println(query, "query")
	rows, err := Db.QueryContext(ctx, query, config.RowsPerPageGenral, offset)
	if err != nil {
		return nil, types.ErrorType{
			Message:    "q" + err.Error(),
			StatusCode: 500,
		}
	}
	defer rows.Close()

	results, err := helpers.ScanRows(rows, columns)
	if err != nil {
		return nil, types.ErrorType{
			Message:    "s" + err.Error(),
			StatusCode: 500,
		}
	}
	if len(results) == 0 {
		return nil, types.ErrorType{
			Message:    fmt.Sprintf("No %s found ", table),
			StatusCode: 404,
		}
	}
	return results, types.ErrorType{}
}

func UpdateRows(table string, updateData interface{}, where *string, validColumns map[string]bool) (int64, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout*time.Second)
	defer cancel()
	updatedMap := helpers.StructToMap(updateData)
	columns, values, err := helpers.ValidateUpdateColumns(updatedMap, validColumns)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	// fmt.Println(columns, values)
	query := helpers.ConstructUpdateQuery(table, columns, where)
	stmt, err := Db.PrepareContext(ctx, query)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, values...)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	if rowsAffected == 0 {
		return 0, types.ErrorType{
			Message:    "No rows updated",
			StatusCode: 404,
		}
	}

	return rowsAffected, types.ErrorType{}
}

// InsertRow inserts a new row into the specified table
func InsertRow(table string, valuesPtr interface{}) (int64, types.ErrorType) {
	values := helpers.StructToMap(valuesPtr)
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	columns := make([]string, 0, len(values))
	params := make([]interface{}, 0, len(values))
	for col, val := range values {
		columns = append(columns, fmt.Sprintf(`"%s"`, col))
		params = append(params, val)
	}

	query := fmt.Sprintf(`
        INSERT INTO "%s" (%s)
        VALUES (%s)
    `, table, helpers.ColumnList(columns), helpers.ParamList(len(params)))
	result, err := Db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	return rowsAffected, types.ErrorType{}
}

func DeleteRow(table, condition string) (int64, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	query := fmt.Sprintf(`
        DELETE FROM "%s"
        WHERE %s
    `, table, condition)

	result, err := Db.ExecContext(ctx, query)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	return rowsAffected, types.ErrorType{}
}
