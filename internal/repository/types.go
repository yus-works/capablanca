package repository

type Column struct {
	Name string
	Type DbType
}

type Table struct {
	Name string
	Columns []Column
	Data    []map[string]interface{}
}

// DbType represents the type of a database column.
type DbType int

const (
	DbUnknown DbType = iota // Unknown type (default)
	DbInt
	DbBigInt
	DbFloat
	DbDouble
	DbDecimal
	DbVarchar
	DbText
	DbBoolean
	DbDate
	DbDateTime
	DbTimestamp
	DbJSON
	DbEnum
)

func (c DbType) String() string {
	switch c {
	case DbInt:
		return "INT"
	case DbBigInt:
		return "BIGINT"
	case DbFloat:
		return "FLOAT"
	case DbDouble:
		return "DOUBLE"
	case DbDecimal:
		return "DECIMAL"
	case DbVarchar:
		return "VARCHAR"
	case DbText:
		return "TEXT"
	case DbBoolean:
		return "BOOLEAN"
	case DbDate:
		return "DATE"
	case DbDateTime:
		return "DATETIME"
	case DbTimestamp:
		return "TIMESTAMP"
	case DbJSON:
		return "JSON"
	case DbEnum:
		return "ENUM"
	default:
		return "UNKNOWN"
	}
}

// ToDbType converts a database column type string to a ColumnType enum.
func ToDbType(dbType string) DbType {
	switch dbType {
	case "int", "integer":
		return DbInt
	case "bigint":
		return DbBigInt
	case "float":
		return DbFloat
	case "double":
		return DbDouble
	case "decimal", "numeric":
		return DbDecimal
	case "varchar", "char", "character varying":
		return DbVarchar
	case "text":
		return DbText
	case "boolean", "bool":
		return DbBoolean
	case "date":
		return DbDate
	case "datetime":
		return DbDateTime
	case "timestamp":
		return DbTimestamp
	case "json":
		return DbJSON
	case "enum":
		return DbEnum
	default:
		return DbUnknown
	}
}
