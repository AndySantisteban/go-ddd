package mssql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.con/AndyGo/go-ddd/internal/domain/entities"
)

// // Generics
// type objectType interface {
// }
type SQLQueryBuilder struct {
	db           *sql.DB
	selectClause string
	fromClause   string
	whereClause  string
	orderBy      string
	// totalRowsClause string
}

func NewSQLQueryBuilder(db *sql.DB) *SQLQueryBuilder {
	return &SQLQueryBuilder{
		db: db,
	}
}

func (qb *SQLQueryBuilder) Select(columns ...string) *SQLQueryBuilder {
	qb.selectClause = "SELECT " + strings.Join(columns, ", ")
	return qb
}
func (qb *SQLQueryBuilder) TotalRows() (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) %s %s", qb.fromClause, qb.whereClause)
	var totalRows int
	err := qb.db.QueryRow(query).Scan(&totalRows)
	if err != nil {
		return 0, err
	}
	return totalRows, nil
}

func (qb *SQLQueryBuilder) From(table string) *SQLQueryBuilder {
	qb.fromClause = "FROM " + table
	return qb
}

func (qb *SQLQueryBuilder) ApplyDataSourceRequest(request *entities.DataSourceRequest) *SQLQueryBuilder {
	if request.Page > 0 && request.PageSize > 0 {
		qb.Paginate(request.Page, request.PageSize)
	}

	if request.Sort != "" {
		qb.OrderBy(request.Sort, true)
	}

	if request.Filter.Field != "" && request.Filter.Value != "" {
		qb.Where(fmt.Sprintf("%s %s '%s'", request.Filter.Field, request.Filter.Operator, request.Filter.Value))
	}

	return qb
}

func (qb *SQLQueryBuilder) Where(condition string) *SQLQueryBuilder {
	if qb.whereClause == "" {
		qb.whereClause = "WHERE " + condition
	} else {
		qb.whereClause += " AND " + condition
	}
	return qb
}

func (qb *SQLQueryBuilder) OrderBy(column string, ascending bool) *SQLQueryBuilder {
	order := "ASC"
	if !ascending {
		order = "DESC"
	}
	qb.orderBy = "ORDER BY " + column + " " + order
	return qb
}

func (qb *SQLQueryBuilder) Paginate(page, pageSize int) *SQLQueryBuilder {
	start := (page - 1) * pageSize
	qb.fromClause += fmt.Sprintf(" ORDER BY (SELECT 0) OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", start, pageSize)
	return qb
}

func (qb *SQLQueryBuilder) Query() (*sql.Rows, int, error) {
	query := fmt.Sprintf("%s %s %s %s", qb.selectClause, qb.fromClause, qb.whereClause, qb.orderBy)
	rows, err := qb.db.Query(query)
	if err != nil {
		return nil, 0, err
	}
	totalRows, err := qb.TotalRows()
	if err != nil {
		rows.Close()
		return nil, 0, err
	}
	return rows, totalRows, nil
}

// Insert realiza una operación de inserción en la base de datos y retorna los valores insertados.
func (qb *SQLQueryBuilder) Insert(tableName string, data interface{}, excludeProps ...string) (interface{}, error) {
	columns, values, err := getColumnsAndValues(data, excludeProps)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) OUTPUT INSERTED.* VALUES (%s)", tableName, columns, values)
	// fmt.Println(query)

	row := qb.db.QueryRow(query)

	dataType := reflect.TypeOf(data)
	newData := reflect.New(dataType).Interface()

	err = row.Scan(structFields(newData)...)
	if err != nil {
		return nil, err
	}

	return newData, nil
}

// getColumnsAndValues extracts column names and values from a structure.
func getColumnsAndValues(data interface{}, excludeProps []string) (string, string, error) {
	value := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		typ = typ.Elem()
	}

	var columns []string
	var values []string

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		propName := field.Name

		if contains(excludeProps, propName) {
			fmt.Printf("Excluding property: %s\n", propName)
			continue
		}

		fieldType := value.Field(i).Type()
		fieldValue := value.Field(i).Interface()

		var valueStr string

		if fieldValue == nil || fieldValue == reflect.Zero(fieldType).Interface() {
			if fieldType == reflect.TypeOf(sql.NullTime{}) || fieldType == reflect.TypeOf(time.Time{}) {
				valueStr = "NULL"
			} else {
				valueStr = "''"
			}
		} else {
			// Convert non-nil values to strings
			switch fieldValue.(type) {
			case string:
				// Handle string values
				if fieldValue.(string) == "" {
					valueStr = "NULL" // or "''" if you prefer an empty string
				} else {
					valueStr = fmt.Sprintf("'%v'", fieldValue)
				}
			case int, int8, int16, int32, int64,
				uint, uint8, uint16, uint32, uint64:
				valueStr = fmt.Sprintf("%v", fieldValue)
			case bool:
				valueStr = fmt.Sprintf("%t", fieldValue)
			case float32, float64:
				valueStr = fmt.Sprintf("%f", fieldValue)
			case time.Time:
				valueStr = fmt.Sprintf("'%s'", fieldValue.(time.Time).Format("1900-01-01 00:00:00"))
			case *string:
				// Handle *string values
				if fieldValue.(*string) == nil || *(fieldValue.(*string)) == "" {
					valueStr = "NULL"
				} else {
					valueStr = fmt.Sprintf("'%v'", *(fieldValue.(*string)))
				}
			case *int, *int8, *int16, *int32, *int64,
				*uint, *uint8, *uint16, *uint32, *uint64:
				// Handle *int values
				if fieldValue.(*int) == nil {
					valueStr = "NULL"
				} else {
					valueStr = fmt.Sprintf("%v", *(fieldValue.(*int)))
				}
			case *bool:
				// Handle *bool values
				if fieldValue.(*bool) == nil {
					valueStr = "NULL"
				} else {
					valueStr = fmt.Sprintf("%v", *(fieldValue.(*bool)))
				}
			case *float32, *float64:
				// Handle *float values
				if fieldValue.(*float64) == nil {
					valueStr = "NULL"
				} else {
					valueStr = fmt.Sprintf("%v", *(fieldValue.(*float64)))
				}
			case *time.Time:
				// Handle *time.Time values
				if fieldValue.(*time.Time) == nil {
					valueStr = "NULL"
				} else {
					valueStr = fmt.Sprintf("'%s'", fieldValue.(*time.Time).Format("1900-01-01 00:00:00"))
				}

			default:
				return "", "", fmt.Errorf("unsupported field type: %v", fieldType)
			}
		}

		values = append(values, valueStr)

		columns = append(columns, propName)
	}

	columnsStr := strings.Join(columns, ", ")
	valuesStr := strings.Join(values, ", ")

	return columnsStr, valuesStr, nil
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice:
		return v.IsNil()
	default:
		// Otros tipos no son manejados aquí
		return false
	}
}

// structFields devuelve un conjunto de reflect.Value para los campos de una estructura.
func structFields(data interface{}) []interface{} {
	v := reflect.ValueOf(data).Elem()
	fields := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		fields[i] = v.Field(i).Addr().Interface()
	}
	return fields
}

// contains verifica si un elemento está presente en una lista.
func contains(list []string, element string) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}
