package paginator

import (
	"database/sql"
	"fmt"
	"math"
)

type Paginator struct {
	db *sql.DB
}

type Pagination struct {
	Rows      *[]Row
	Page      int
	PageSize  int
	PageCount int
	RowCount  int
}

type query struct {
	sql  string
	args []interface{}
}

const maxPageSize = 100

func New(db *sql.DB) *Paginator {
	return &Paginator{
		db: db,
	}
}

func (p *Paginator) Paginate(q *query, page, pageSize int) (*Pagination, error) {
	rowCount, err := p.countRows(q)
	if err != nil {
		return nil, err
	}
	if rowCount == 0 {
		return new(Pagination), nil
	}

	if page < 1 {
		page = 1
	}

	pageCount := int(math.Ceil(float64(rowCount) / float64(pageSize)))

	if page > pageCount {
		page = pageCount
	}

	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	sql := fmt.Sprintf("Select t.* From (%s) t Limit %d Offset %d", q.sql, pageSize, (page-1)*pageSize)
	rows, err := doQuery(p.db, sql, q.args...)
	if err != nil {
		return nil, err
	}

	return &Pagination{
		Rows:      rows,
		Page:      page,
		PageSize:  pageSize,
		PageCount: pageCount,
		RowCount:  rowCount,
	}, nil
}

func (p *Paginator) countRows(q *query) (count int, err error) {
	sql := fmt.Sprintf("Select Count(1) From (%s) t", q.sql)
	err = p.db.QueryRow(sql, q.args...).Scan(&count)
	return
}

func doQuery(db *sql.DB, sqlInfo string, args ...interface{}) (*[]Row, error) {
	dbRows, err := db.Query(sqlInfo, args...)
	if err != nil {
		return nil, err
	}

	columns, _ := dbRows.Columns()
	columnLength := len(columns)
	scanArgs := make([]interface{}, columnLength)
	for i, _ := range scanArgs {
		var scanArg string
		scanArgs[i] = &scanArg
	}

	var rows []Row
	for dbRows.Next() {
		_ = dbRows.Scan(scanArgs...)

		var row Row
		row = make(Row)
		for i, scanArg := range scanArgs {
			row[columns[i]] = *scanArg.(*string)
		}

		rows = append(rows, row)
	}

	defer dbRows.Close()

	return &rows, nil
}

func (p *Paginator) CreateQuery(sql string, args ...interface{}) *query {
	return &query{
		sql,
		args,
	}
}

func (p *Pagination) RowIndex(i int) Row {
	return (*p.Rows)[i]
}
