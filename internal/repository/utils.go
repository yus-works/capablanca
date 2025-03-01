package repository

import (
	"gorm.io/gorm"
	"time"
)

// TODO: cache metadata queries somehow

// GetTableNames returns a list of all table names in the database.
func GetTableNames(db *gorm.DB) ([]string, error) {
	var tables []string
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()"
	if err := db.Raw(query).Scan(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

func GetTable(db *gorm.DB, tableName string) (*Table, error) {
    columns, err := GetTableColumns(db, tableName)
    if err != nil {
        return nil, err
    }

	data, err := GetTableData(db, tableName)
    if err != nil {
        return nil, err
    }

    return &Table{
        Columns: columns,
        Data:    data,
    }, nil
}

// GetTableData retrieves all rows from the specified table.
func GetTableData(db *gorm.DB, tableName string) ([]map[string]interface{}, error) {
    var data []map[string]interface{}
    if err := db.Table(tableName).Find(&data).Error; err != nil {
        return nil, err
    }

	for _, row := range data {
		for key, value := range row {
			if t, ok := value.(time.Time); ok {
				row[key] = t.Format("2006-01-02 15:04") // Format without seconds
			}
		}
	}
	return data, nil
}

func GetTableColumns(db *gorm.DB, tableName string) ([]Column, error) {
	type columnRow struct {
		ColumnName string
		DataType   string
	}

	var columnRows []columnRow
	query := `
		SELECT
			COLUMN_NAME AS column_name,
			DATA_TYPE AS data_type
		FROM
			information_schema.columns
		WHERE
			table_name = ? AND table_schema = DATABASE()
		ORDER BY
			CASE 
				WHEN COLUMN_NAME = 'id' THEN 1 
				WHEN COLUMN_NAME LIKE '%\_id' THEN 2 
				ELSE 3 
			END,
			ordinal_position
	`

	if err := db.Raw(query, tableName).Scan(&columnRows).Error; err != nil {
		return nil, err
	}

	columns := make([]Column, 0, len(columnRows))
	for _, row := range columnRows {
		columns = append(columns, Column{
			Name: row.ColumnName,
			Type: ToDbType(row.DataType),
		})
	}

	return columns, nil
}

// RecordExists checks if a record exists in the given table based on a condition.
func RecordExists(db *gorm.DB, tableName string, condition string, args ...interface{}) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM " + tableName + " WHERE " + condition + ")"
	if err := db.Raw(query, args...).Scan(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

// CountRows returns the number of rows in a given table.
func CountRows(db *gorm.DB, tableName string) (int64, error) {
	var count int64
	query := "SELECT COUNT(*) FROM " + tableName
	if err := db.Raw(query).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteRecord deletes a record from the specified table with a given condition.
func DeleteRecord(db *gorm.DB, tableName string, condition string, args ...interface{}) error {
	query := "DELETE FROM " + tableName + " WHERE " + condition
	if err := db.Exec(query, args...).Error; err != nil {
		return err
	}
	return nil
}
