package book

import (
	"bookstore.weber.edu/domain"
	"bookstore.weber.edu/domain/book"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}
type MysqlRow struct {
	Rows *sql.Rows
}

func NewBookMysqlRepository(u string, p string, h string, prt string, dbn string) (*MysqlRepository, error) {
	dsn := u + ":" + p + "@tcp(" + h + ":" + prt + ")/" + dbn + "?parseTime=true&timeout=10s&collation=utf8_general_ci"
	dbh, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = dbh.Ping(); err != nil {
		return nil, err
	}
	dbh.SetMaxIdleConns(10000)
	return &MysqlRepository{db: dbh}, nil
}

func (r *MysqlRepository) Write(b *book.Book) (domain.ID, error) {
	stmt, err := r.db.Prepare("INSERT INTO books (isbn, title, author, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return b.ID, err
	}
	_, err = stmt.Exec(b.ID, b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		return b.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return b.ID, err
	}
	return b.ID, nil
}

func (r *MysqlRepository) ReadAll() ([]book.Book, error) {
	stmt, err := r.db.Prepare("SELECT isbn, title, author, price FROM books")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var bks []book.Book
	for rows.Next() {
		var b book.Book
		err = rows.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price)
		if err != nil {
			return nil, err
		}
		bks = append(bks, b)
	}
	return bks, nil
}
