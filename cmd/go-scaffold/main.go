// main.go
package main

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

func Program() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: ./executable.exe <nombre_de_tabla>")
		os.Exit(1)
	}
	dsn := "sqlserver://sysugosrv02:UGO!dev0823@96.88.124.75:1433?database=CenturionNotes&connection+timeout=30&encrypt=disable"
	db, err := sqlx.Connect("mssql", dsn)
	if err != nil {
		panic(err)
	}
	tableName := os.Args[1]

	GenerateStructFromTable(db, tableName)

}

func GenerateStructFromTable(db *sqlx.DB, tableName string) {

	var columns []struct {
		ColumnName string `db:"column_name"`
		DataType   string `db:"data_type"`
		IsNullable string `db:"is_nullable"`
	}
	consult := strings.Replace("USE CenturionNotes;  SELECT data_type, column_name, is_nullable FROM INFORMATION_SCHEMA.COLUMNS  WHERE TABLE_SCHEMA = 'dbo' AND TABLE_NAME = 'tableName'", "tableName", tableName, 1)
	fmt.Print("consult", consult)
	err := db.Select(&columns, consult)
	if err != nil {
		panic(err)
	}

	code := generateGoCode(tableName, columns)

	filePath := fmt.Sprintf("tmp/models/%s.go", tableName)
	err = writeToFile(filePath, code)
	if err != nil {
		panic(err)
	}

}

func toCamelCase(s string) string {
	// parts := strings.Split(s, "_")
	// for i, part := range parts {
	// 	parts[i] = strings.Title(part)
	// }
	// return strings.Join(parts, "")
	return s
}

func sqlTypeToGoType(sqlType string) string {
	switch sqlType {
	case "int":
		return "int"
	case "bigint":
		return "int"
	case "text":
		return "string"
	case "varchar":
		return "string"
	case "datetime":
		return "time.Time"
	case "smallint":
		return "int16"
	case "decimal":
		return "float64"
	case "float":
		return "float64"
	case "double":
		return "float64"
	case "money":
		return "float64"
	case "bit":
		return "bool"
	case "timestamp":
		return "[]uint8"
	default:
		return "interface{}"
	}
}

func generateGoCode(tableName string, columns []struct {
	ColumnName string `db:"column_name"`
	DataType   string `db:"data_type"`
	IsNullable string `db:"is_nullable"`
}) string {
	var code strings.Builder

	// Genera la estructura en Go
	code.WriteString(fmt.Sprintf("package %s\n\n", "models"))
	code.WriteString(fmt.Sprintf("import (\n\t%s\n)\n\n", "\"time\""))
	code.WriteString(fmt.Sprintf("type %s struct {\n", tableName))
	for _, col := range columns {
		fmt.Printf("column:\t%s\n", col)
		code.WriteString(fmt.Sprintf("    %s %s%s\n", toCamelCase(col.ColumnName), validateIfNullable(col.IsNullable), sqlTypeToGoType(col.DataType)))
	}
	code.WriteString("}\n")

	return code.String()
}

func validateIfNullable(s string) string {
	if s == "NO" {
		return ""
	}
	return "*"
}

func writeToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	fmt.Print("\n")
	fmt.Print(" ---------------------------------------------------------- \n")
	fmt.Print(" ---------------------------------------------------------- ")
	fmt.Printf("\n Scaffolding Completed!!!, \n The Model is in %s \n", filePath)
	fmt.Print(" ---------------------------------------------------------- \n")
	fmt.Print(" ---------------------------------------------------------- \n")
	return nil

}

func main() {
	Program()
}
