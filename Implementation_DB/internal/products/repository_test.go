package products

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/internal/domains"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Implementation_DB/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Store_GetByName_tdxdb(t *testing.T) {
	db := utils.InitTxDB()
	defer db.Close()

	repo := NewRepository(db)

	ctx := context.TODO()

	productExpected := domains.Product{
		Name:  "Producto Nuevo",
		Type:  "Laptop",
		Count: 12,
		Price: 2352.33,
	}
	t.Run("Ok", func(t *testing.T) {
		id, err := repo.Store(ctx, productExpected)
		assert.NoError(t, err)

		productExpected.ID = int(id)

		productResult, err := repo.GetByName(ctx, productExpected.Name)

		assert.NoError(t, err)
		assert.NotEmpty(t, productResult)
		assert.Equal(t, productExpected, productResult)
	})

	t.Run("Fail", func(t *testing.T) {
		productResult, err := repo.GetByName(ctx, "Nombre X")

		assert.NotNil(t, err)
		assert.Empty(t, productResult)
	})
}

func Test_Update(t *testing.T) {

	product := domains.Product{
		ID:    1,
		Name:  "ProductoX",
		Type:  "Laptop",
		Count: 12,
		Price: 2352.33,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE))
	mock.ExpectExec(regexp.QuoteMeta(UPDATE)).WithArgs(product.Name, product.Type, product.Count, product.Price, product.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	_, err = repo.Update(context.TODO(), product.ID, product.Name, product.Type, product.Count, product.Price)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_Delete(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 2

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE))
	mock.ExpectExec(regexp.QuoteMeta(DELETE)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	err = repo.Delete(context.TODO(), int64(id))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	t.Run("FAIl", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(DELETE)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 0))
		err = repo.Delete(context.TODO(), int64(1))
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}
