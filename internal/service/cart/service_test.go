package cart

import (
	"context"
	"fmt"
	"interview/internal/dto"
	"interview/internal/entity"
	cartservicerepomock "interview/internal/mocks/cartservice_repo_mock"
	productservicemock "interview/internal/mocks/productservice_mock"
	"interview/internal/pkg/richerror"
	"testing"

	"go.uber.org/mock/gomock"
)

func setup(t *testing.T) (*Service, *cartservicerepomock.MockRepository, *productservicemock.MockProductService) {
	ctrl := gomock.NewController(t)

	cartRepo := cartservicerepomock.NewMockRepository(ctrl)
	productService := productservicemock.NewMockProductService(ctrl)
	// productCache := productservicecache.NewMockCachRepository(ctrl)

	cartService := New(cartRepo, productService)

	return cartService, cartRepo, productService
}

func TestAddItemToCart(t *testing.T) {
	t.Run("find cart with unexpected error", func(t *testing.T) {
		service, repo, _ := setup(t)

		req := dto.AddItemToCartRequest{}

		repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(entity.Cart{}, false, fmt.Errorf("unexpected error"))

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart does not exist and cannot create a new one", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
		}

		cart := entity.Cart{SessionID: req.SessionID}

		findQuery := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(entity.Cart{}, false, nil)
		repo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(cart, fmt.Errorf("unexpected error")).After(findQuery)

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart does not exist and getting the product fails", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
		}

		cart := entity.Cart{
			SessionID: req.SessionID,
		}

		findQuery := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(entity.Cart{}, false, nil)
		repo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(cart, nil).After(findQuery)

		productService.EXPECT().GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(dto.ProductGetItemResponse{}, fmt.Errorf("unexpected"))

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart does not exist and creating a new cart item fails", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
		}

		cart := entity.Cart{
			SessionID: req.SessionID,
		}

		findQuery := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(entity.Cart{}, false, nil)
		createItem := repo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(cart, nil).After(findQuery)
		repo.EXPECT().CreateCartItem(gomock.Any(), gomock.Any()).Return(entity.CartItem{}, fmt.Errorf("unexpected")).After(createItem)

		productService.EXPECT().GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       300,
		}, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}
	})

	t.Run("cart does not exist and the operation is successful", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		product := dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       200,
		}

		cartItem := entity.CartItem{
			ProductName: req.Product,
			Quantity:    req.Quantity,
			CartID:      cart.ID,
			Price:       product.Price,
		}

		findQuery := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(entity.Cart{}, false, nil)
		createItem := repo.EXPECT().CreateCart(gomock.Any(), gomock.Any()).Return(cart, nil).After(findQuery)
		repo.EXPECT().CreateCartItem(gomock.Any(), cartItem).Return(cartItem, nil).After(createItem)

		productService.EXPECT().
			GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(product, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	})

	t.Run("cart exists and find_cart returns unexpected error", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		findCart := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(cart, true, nil)
		repo.EXPECT().FindCartItemByProduct(gomock.Any(), cart.ID, req.Product).Return(entity.CartItem{}, false, fmt.Errorf("unexpected")).After(findCart)

		productService.EXPECT().GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       300,
		}, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart exists and create new cart item fails", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		cartItem := entity.CartItem{
			ProductName: req.Product,
			Quantity:    req.Quantity,
			CartID:      cart.ID,
			Price:       200,
		}

		findCart := repo.EXPECT().FindOpenCartBySessionID(gomock.Any(), gomock.Any()).Return(cart, true, nil)
		findCartItem := repo.EXPECT().FindCartItemByProduct(gomock.Any(), cart.ID, req.Product).Return(entity.CartItem{}, false, nil).After(findCart)
		repo.EXPECT().CreateCartItem(gomock.Any(), cartItem).Return(entity.CartItem{}, fmt.Errorf("unexpected error")).After(findCartItem)

		productService.EXPECT().GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       200,
		}, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart exists and create new cart item is successful", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		cartItem := entity.CartItem{
			ProductName: req.Product,
			Quantity:    req.Quantity,
			CartID:      cart.ID,
			Price:       200,
		}

		findCart := repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), gomock.Any()).
			Return(cart, true, nil)
		findCartItem := repo.EXPECT().
			FindCartItemByProduct(gomock.Any(), cart.ID, req.Product).
			Return(entity.CartItem{}, false, nil).After(findCart)
		repo.EXPECT().
			CreateCartItem(gomock.Any(), cartItem).
			Return(cartItem, nil).
			After(findCartItem)

		productService.EXPECT().GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).Return(dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       200,
		}, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	})

	t.Run("cart exists, cart item exists and updating cart item fails", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		product := dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       200,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		cartItem := entity.CartItem{
			ProductName: req.Product,
			Quantity:    req.Quantity,
			CartID:      cart.ID,
			Price:       200,
		}

		result := entity.CartItem{
			ProductName: req.Product,
			Quantity:    cartItem.Quantity + 1,
			CartID:      cart.ID,
			Price:       product.Price * (float64(cartItem.Quantity) + 1),
		}

		findCart := repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), gomock.Any()).
			Return(cart, true, nil)
		findCartItem := repo.EXPECT().
			FindCartItemByProduct(gomock.Any(), cart.ID, req.Product).
			Return(cartItem, true, nil).After(findCart)
		repo.EXPECT().
			UpdateCartItem(gomock.Any(), result).
			Return(fmt.Errorf("unexpected error")).
			After(findCartItem)

		productService.EXPECT().
			GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).
			Return(product, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart exists, cart item exists and updating cart item is successful", func(t *testing.T) {
		service, repo, productService := setup(t)

		sessionID := "session_id"

		req := dto.AddItemToCartRequest{
			SessionID: sessionID,
			Product:   "Shoe",
			Quantity:  1,
		}

		product := dto.ProductGetItemResponse{
			ProductName: req.Product,
			Price:       200,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		cartItem := entity.CartItem{
			ProductName: req.Product,
			Quantity:    req.Quantity,
			CartID:      cart.ID,
			Price:       200,
		}

		result := entity.CartItem{
			ProductName: req.Product,
			Quantity:    cartItem.Quantity + 1,
			CartID:      cart.ID,
			Price:       product.Price * (float64(cartItem.Quantity) + 1),
		}

		findCart := repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), gomock.Any()).
			Return(cart, true, nil)
		findCartItem := repo.EXPECT().
			FindCartItemByProduct(gomock.Any(), cart.ID, req.Product).
			Return(cartItem, true, nil).After(findCart)
		repo.EXPECT().
			UpdateCartItem(gomock.Any(), result).
			Return(nil).
			After(findCartItem)

		productService.EXPECT().
			GetProduct(gomock.Any(), dto.ProductGetItemRequest{ProductName: req.Product}).
			Return(product, nil)

		err := service.AddItemToCart(context.Background(), req)
		if err != nil {
			t.Error("no error is returned")
			t.FailNow()
		}
	})
}

func TestDeleteCartItem(t *testing.T) {
	t.Run("find cart is unsuccessful", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.DeleteCartItemRequest{
			SessionID:  sessionID,
			CartItemID: 1,
		}

		repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), sessionID).
			Return(entity.Cart{}, false, fmt.Errorf("unexpected"))

		err := service.DeleteCartItem(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("cart does not exist", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.DeleteCartItemRequest{
			SessionID:  sessionID,
			CartItemID: 1,
		}

		repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), sessionID).
			Return(entity.Cart{}, false, nil)

		err := service.DeleteCartItem(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindNotFound {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("deletion is unsuccessful", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.DeleteCartItemRequest{
			SessionID:  sessionID,
			CartItemID: 1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		findCart := repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), sessionID).
			Return(cart, true, nil)
		repo.EXPECT().
			DeleteCartItemByID(gomock.Any(), req.CartItemID).
			Return(fmt.Errorf("unexpected")).
			After(findCart)

		err := service.DeleteCartItem(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindUnexpected {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})

	t.Run("deletion is successful", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.DeleteCartItemRequest{
			SessionID:  sessionID,
			CartItemID: 1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
		}

		findCart := repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), sessionID).
			Return(cart, true, nil)
		repo.EXPECT().
			DeleteCartItemByID(gomock.Any(), req.CartItemID).
			Return(nil).
			After(findCart)

		err := service.DeleteCartItem(context.Background(), req)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	})

	t.Run("cart is closed", func(t *testing.T) {
		service, repo, _ := setup(t)

		sessionID := "session_id"

		req := dto.DeleteCartItemRequest{
			SessionID:  sessionID,
			CartItemID: 1,
		}

		cart := entity.Cart{
			ID:        1,
			SessionID: req.SessionID,
			Status:    entity.CartClosed,
		}

		repo.EXPECT().
			FindOpenCartBySessionID(gomock.Any(), sessionID).
			Return(cart, true, nil)

		err := service.DeleteCartItem(context.Background(), req)
		if err == nil {
			t.Error("no error is returned")
			t.FailNow()
		}

		if richErr, ok := err.(richerror.RichError); !ok {
			t.Error("returned error is not of type richerror")
			t.FailNow()
		} else {
			if richErr.Kind() != richerror.KindInvalid {
				t.Error("error must be unexpected")
				t.FailNow()
			}
		}
	})
}
