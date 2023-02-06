package info

import (
	"errors"

	"github.com/mrexmelle/go-psql-size-dist/config"
)

func (info *Table) FindLabelNameInRule(rule config.Rule) (string, error) {
	for _, r := range rule.Labels {
		for _, t := range r.Tables {
			if t == info.Name {
				return t, nil
			}
		}
	}

	return "", errors.New("Cannot find info name")
}

func GetSizeOfTableName(tables []Table, tableName string) (uint64, error) {
	for _, table := range tables {
		if table.Name == tableName {
			return table.Size, nil
		}
	}

	return 0, errors.New("cannot find table name")
}

func Delete(tables []Table, tableName string) []Table {
	for i := len(tables) - 1; i >= 0; i-- {
		if tables[i].Name == tableName {
			tables = append(tables[:i], tables[i+1:]...)
		}
	}

	return tables
}
