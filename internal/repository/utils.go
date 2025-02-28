package repository

import (
	"gorm.io/gorm"
	"time"
)

// TODO: cache metadata queries somehow
// TODO: test

// GetTableNames returns a list of all table names in the database.
func GetTableNames(db *gorm.DB) ([]string, error) {
	var tables []string
	query := "SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()"
	if err := db.Raw(query).Scan(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

// GetTableData retrieves all rows from the specified table.
func GetTableData(db *gorm.DB, tableName string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	query := "SELECT * FROM " + tableName
	if err := db.Raw(query).Scan(&results).Error; err != nil {
		return nil, err
	}
	for _, row := range results {
		for key, value := range row {
			if t, ok := value.(time.Time); ok {
				row[key] = t.Format("2006-01-02 15:04") // Format without seconds
			}
		}
	}
	return results, nil
}

// GetTableColumns returns the column names of a specific table.
func GetTableColumns(db *gorm.DB, tableName string) ([]string, error) {
	var columns []string
	query := `
		SELECT
			COLUMN_NAME
		FROM
			information_schema.columns
		WHERE
			table_name = ? AND table_schema = DATABASE()
	`
	if err := db.Raw(query, tableName).Scan(&columns).Error; err != nil {
		return nil, err
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
