package helpers

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/piyush7833/Chat-Api/config"
)

// get helpers
func ValidateColumns(selectedColumns []string, validColumns map[string]bool, joinColumns map[string]string, orderByColumns []string) ([]string, []string, error) {
	// Pre-allocate slices with a reasonable initial capacity
	columns := []string{}
	orderBy := []string{}

	// Process selectedColumns
	if joinColumns != nil {
		for _, col := range selectedColumns {
			if validColumns[col] {
				parts := strings.Split(joinColumns[col], ".")
				columns = append(columns, fmt.Sprintf(`%s."%s" as "%s"`, parts[0], parts[1], col))
			}
		}
	} else {
		for _, col := range selectedColumns {
			if validColumns[col] {
				columns = append(columns, fmt.Sprintf(`"%s"`, col))
			}
		}
	}

	// Process orderByColumns
	if joinColumns != nil {
		for _, col := range orderByColumns {
			if validColumns[col] {
				parts := strings.Split(joinColumns[col], ".")
				orderBy = append(orderBy, fmt.Sprintf(`%s."%s"`, parts[0], parts[1]))
			}
		}
	} else {
		for _, col := range orderByColumns {
			if validColumns[col] {
				orderBy = append(orderBy, fmt.Sprintf(`"%s"`, col))
			}
		}
	}

	// Add default orderBy if none are valid
	if len(orderBy) == 0 {
		if joinColumns != nil {
			orderBy = append(orderBy, `t."createdAt"`)
		} else {
			orderBy = append(orderBy, `"createdAt"`)
		}
	}

	// Return error if no valid columns selected
	if len(columns) == 0 {
		return nil, nil, fmt.Errorf("no valid columns selected")
	}

	return columns, orderBy, nil
}

func ConstructGetQuery(table string, columns []string, page int, where *string, joins *string, orderBy []string, isDesc bool) (string, int) {
	selectColumns := strings.Join(columns, ", ")
	orderByColumns := strings.Join(orderBy, ", ")
	desc := "ASC"
	if isDesc {
		desc = "DESC"
	}

	var query string
	if joins != nil && where != nil {
		query = fmt.Sprintf(`
			SELECT %s FROM "%s" t
			%s
			WHERE %s
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, selectColumns, table, *joins, *where, orderByColumns, desc)
	} else if joins != nil {
		query = fmt.Sprintf(`
			SELECT %s FROM "%s" t
			%s
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, selectColumns, table, *joins, orderByColumns, desc)
	} else if where != nil {
		query = fmt.Sprintf(`
			SELECT %s FROM "%s"
			WHERE %s
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, selectColumns, table, *where, orderByColumns, desc)
	} else {
		query = fmt.Sprintf(`
			SELECT %s FROM "%s"
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, selectColumns, table, orderByColumns, desc)
	}
	return query, page * config.RowsPerPageGenral
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
			value := *(columnPointers[i].(*interface{}))
			switch v := value.(type) {
			case []byte:
				columnMap[colName] = string(v) // Convert []byte to string
			default:
				columnMap[colName] = value
			}
		}

		results = append(results, columnMap)
	}
	return results, nil
}

// update helpers
func ValidateUpdateColumns(updateData map[string]interface{}, validColumns map[string]bool) ([]string, []interface{}, error) {
	type result struct {
		column string
		value  interface{}
		err    error
	}

	results := make(chan result, len(updateData))
	columns := []string{}
	values := []interface{}{}

	for col, val := range updateData {
		go func(col string, val interface{}) {
			if validColumns[col] && val != nil {
				results <- result{
					column: fmt.Sprintf(`"%s" = $%d`, col, len(values)+1),
					value:  val,
					err:    nil,
				}
			} else {
				results <- result{
					err: fmt.Errorf("invalid column or nil value: %s", col),
				}
			}
		}(col, val)
	}

	for range updateData {
		res := <-results
		if res.err == nil {
			columns = append(columns, res.column)
			values = append(values, res.value)
		}
	}
	close(results)

	if len(columns) == 0 {
		return nil, nil, fmt.Errorf("no valid columns selected")
	}

	return columns, values, nil
}

func ConstructUpdateQuery(table string, columns []string, where *string) string {
	setClause := strings.Join(columns, ", ")
	if where == nil {
		return fmt.Sprintf(`
			UPDATE "%s"
			SET %s
		`, table, setClause)
	}
	return fmt.Sprintf(`
        UPDATE "%s"
        SET %s
        WHERE %s
    `, table, setClause, *where)
}

// insert helpers
func ColumnList(columns []string) string {
	return join(columns, ", ")
}

// paramList returns a comma-separated list of parameter placeholders
func ParamList(count int) string {
	params := make([]string, count)
	for i := 0; i < count; i++ {
		params[i] = fmt.Sprintf("$%d", i+1)
	}
	return join(params, ", ")
}

// join concatenates strings with a separator
func join(string []string, sep string) string {
	return strings.Join(string, sep)
}
