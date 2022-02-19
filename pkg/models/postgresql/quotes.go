package postgresql

import (
	"database/sql"
)

type QuoteModel struct {
	DB *sql.DB
}

// TODO

// func (m *QuoteModel) Checker() (int, error) {
// 	stmt := `select 2 + 2;`


// 	rows, err := m.DB.Query(stmt)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer rows.Close()

// 	var n int

// 	for rows.Next() {
		
// 		err = rows.Scan(&n)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return n, nil
// }