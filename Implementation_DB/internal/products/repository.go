package products

import (
	"context"
	"database/sql"
	"errors"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domains.Product, error)
	Store(ctx context.Context, p domains.Product) (int, error)
	GetAll(ctx context.Context) ([]domains.Product, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int, name, ptype string, count int, price float64) (domains.Product, error)
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
	UPDATE       = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?;"
	GET_ALL      = "SELECT id, name, type, count, price FROM products;"
	DELETE       = "DELETE FROM products WHERE id=?;"
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

func (r *repository) GetAll(ctx context.Context) ([]domains.Product, error) {
	var products []domains.Product
	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return []domains.Product{}, err
	}

	for rows.Next() {
		var product domains.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return []domains.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	smt, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer smt.Close()

	result, err := smt.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}

func (r *repository) GetAll_Context(ctx context.Context) ([]domains.Product, error) {
	var products []domains.Product
	rows, err := r.db.QueryContext(ctx, GET_ALL)
	if err != nil {
		return []domains.Product{}, err
	}
	for rows.Next() {
		var product domains.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price)
		if err != nil {
			return []domains.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *repository) Update(ctx context.Context, id int, name, ptype string, count int, price float64) (domains.Product, error) {
	stmt, err := r.db.Prepare(UPDATE)

	if err != nil {
		return domains.Product{}, err
	}
	defer stmt.Close()

	product := domains.Product{ID: id, Name: name, Type: ptype, Count: count, Price: price}
	_, err = stmt.Exec(name, ptype, count, price, id)
	if err != nil {
		return domains.Product{}, err
	}
	return product, nil
}
