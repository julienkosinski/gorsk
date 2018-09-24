package pgsql

import "github.com/go-pg/pg/orm"

// Keeps using singular table names
func init() {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})
}
