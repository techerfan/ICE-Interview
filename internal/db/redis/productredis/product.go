package productredis

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func (d DB) SetProduct(ctx context.Context, productName string, cost float64, expTime time.Duration) error {
	_, err := d.client.Set(ctx, productName, cost, expTime).Result()
	if err != nil {
		return fmt.Errorf("could not store the product in the cache store: %v", err)
	}

	return nil
}

func (d DB) GetProduct(ctx context.Context, productName string) (float64, bool) {
	prictStr, err := d.client.Get(ctx, productName).Result()
	if err != nil {
		return -1, false
	}

	price, err := strconv.ParseFloat(prictStr, 64)
	if err != nil {
		return -1, false
	}

	return price, true
}

func (d DB) DeleteProduct(ctx context.Context, productName string) error {
	_, err := d.client.Del(ctx, productName).Result()
	if err != nil {
		return fmt.Errorf("could not delete the product from cache store: %v", err)
	}

	return nil
}
