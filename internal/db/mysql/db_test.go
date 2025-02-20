package mysql

import (
	"context"
	"interview/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) *Mysql {
	repo := New(Config{
		Host:     "localhost",
		Port:     "4002",
		Username: "ice_user",
		Password: "9xz3jrd8wf",
		DBName:   "ice_db",
	})

	repo.db.AutoMigrate(&Cart{}, &CartItem{})

	t.Cleanup(func() {
		err := repo.db.Migrator().DropTable(&CartItem{})
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		err = repo.db.Migrator().DropTable(&Cart{})
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	})

	return repo
}

func TestCreateCart(t *testing.T) {
	db := setup(t)

	sessionID := "sample_id"

	cart := entity.Cart{
		ID:        0,
		Total:     0,
		SessionID: sessionID,
		Status:    entity.CartOpen,
	}

	var err error
	cart, err = db.CreateCart(context.Background(), cart)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, cart.ID)
}

func TestUpdateCart(t *testing.T) {
	// TODO:
}

func TestFindCartByID(t *testing.T) {
	db := setup(t)

	t.Run("fail on finding cart", func(t *testing.T) {
		_, exist, err := db.FindCartByID(context.Background(), 1)
		assert.Nil(t, err)
		assert.Equal(t, false, exist)
	})

	t.Run("successful on finding cart", func(t *testing.T) {
		sessionID := "sample_id"

		cart := entity.Cart{
			ID:        0,
			Total:     0,
			SessionID: sessionID,
			Status:    entity.CartOpen,
		}

		cart, err := db.CreateCart(context.Background(), cart)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		_, exist, err := db.FindCartByID(context.Background(), cart.ID)
		assert.Nil(t, err)
		assert.Equal(t, true, exist)
	})
}

func TestFindOpenCartBySessionID(t *testing.T) {
	// TODO:
}

func TestCreateCartItem(t *testing.T) {
	// TODO:
}

func TestUpdateCartItem(t *testing.T) {
	// TODO:
}

func TestFindCartItemByID(t *testing.T) {
	// TODO:
}

func TestFindCartItemByProduct(t *testing.T) {
	// TODO:
}

func TestFindCartItemsByCartID(t *testing.T) {
	// TODO:
}

func TestDeleteCartItemByID(t *testing.T) {
	// TODO:
}
