package services

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
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
		log.Fatal(err)
	}

	// Ping the database to check if the connection is successful
	err = Db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to the database")

}

// Now you can perform database operations using this 'db' connection

func ValidateColumns(selectedColumns []string, validColumns map[string]bool) ([]string, error) {
	columns := []string{}
	for _, col := range selectedColumns {
		if validColumns[col] {
			columns = append(columns, fmt.Sprintf(`"%s"`, col))
		}
	}
	if len(columns) == 0 {
		return nil, fmt.Errorf("no valid columns selected")
	}
	return columns, nil
}

func ConstructGetQuery(table string, columns []string, page int, where *string) (string, int) {
	selectColumns := strings.Join(columns, ", ")
	var query string
	if where != nil {
		query = fmt.Sprintf(`
			SELECT %s FROM "%s"
			WHERE %s
			ORDER BY "createdAt" DESC
			LIMIT $1 OFFSET $2
		`, selectColumns, table, *where)
	} else {
		query = fmt.Sprintf(`
				SELECT %s FROM "%s"
				ORDER BY "createdAt" DESC
				LIMIT $1 OFFSET $2
			`, selectColumns, table)
	}
	return query, page
}

func ScanRows(rows *sql.Rows, columns []string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		columnPointers := make([]interface{}, len(colNames))
		columnMap := make(map[string]interface{})
		for i := range columnPointers {
			columnPointers[i] = new(interface{})
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		for i, colName := range colNames {
			columnMap[colName] = *(columnPointers[i].(*interface{}))
		}

		results = append(results, columnMap)
	}
	return results, nil
}

// GetRows fetches rows from a specified table with pagination and selected columns
func GetRows(table string, page int, selectedColumns []string, validColumns map[string]bool, where *string) ([]map[string]interface{}, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout*time.Second)
	defer cancel()

	columns, err := ValidateColumns(selectedColumns, validColumns)
	if err != nil {
		return nil, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	query, offset := ConstructGetQuery(table, columns, page, where)
	// fmt.Println(query)
	rows, err := Db.QueryContext(ctx, query, config.RowsPerPageGenral, offset)
	if err != nil {
		return nil, types.ErrorType{
			Message:    "q" + err.Error(),
			StatusCode: 500,
		}
	}
	defer rows.Close()

	results, err := ScanRows(rows, columns)
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

//update

func ConstructUpdateQuery(table string, columns []string, where *string) string {
	setClause := strings.Join(columns, ", ")
	query := fmt.Sprintf(`
        UPDATE "%s"
        SET %s
        WHERE %s
    `, table, setClause, *where)
	return query
}

func ValidateUpdateColumns(updateData map[string]interface{}, validColumns map[string]bool) ([]string, []interface{}, error) {
	columns := []string{}
	values := []interface{}{}
	for col, val := range updateData {
		if validColumns[col] && val != nil {
			columns = append(columns, fmt.Sprintf(`"%s" = $%d`, col, len(values)+1))
			values = append(values, val)
		}
	}
	if len(columns) == 0 {
		return nil, nil, fmt.Errorf("no valid columns selected")
	}
	return columns, values, nil
}

func UpdateRows(table string, updateData interface{}, where *string, validColumns map[string]bool) (int64, types.ErrorType) {
	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout*time.Second)
	defer cancel()
	updatedMap := helpers.StructToMap(updateData)
	columns, values, err := ValidateUpdateColumns(updatedMap, validColumns)
	if err != nil {
		return 0, types.ErrorType{
			Message:    err.Error(),
			StatusCode: 400,
		}
	}

	fmt.Println(columns, values)
	query := ConstructUpdateQuery(table, columns, where)
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
    `, table, columnList(columns), paramList(len(params)))

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

// columnList returns a comma-separated list of columns
func columnList(columns []string) string {
	return join(columns, ", ")
}

// paramList returns a comma-separated list of parameter placeholders
func paramList(count int) string {
	return join(make([]string, count), ", ")
}

// join concatenates strings with a separator
func join(string []string, sep string) string {
	return strings.Join(string, sep)
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
