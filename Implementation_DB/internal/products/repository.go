package products

import (
	"context"
	"database/sql"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domains.Product, error)
	Store(ctx context.Context, p domains.Product) (int, error)
}
type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	GET_BY_NAME  = "SELECT id, name, type, count, price FROM products WHERE name=?;"
	SAVE_PRODUCT = "INSERT INTO products (name, type, count, price) VALUES (?,?,?,?);"
)

func (r *repository) GetByName(ctx context.Context, name string) (domains.Product, error) {
	row := r.db.QueryRow(GET_BY_NAME, name) //CONSULTA
	var product domains.Product
	err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price) //LO AGREGA A LA VARIABLE
	if err != nil {
		return domains.Product{}, err
	}
	return product, nil

}

func (r *repository) Store(ctx context.Context, p domains.Product) (int, error) {
	stmt, err := r.db.Prepare(SAVE_PRODUCT) //SE PREPARA LA CONSULTA
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(p.Name, p.Type, p.Count, p.Price) //SE EJECUTA LA CONSULTA
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId() //SE OBTIENE EL ULTIMO ID
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
