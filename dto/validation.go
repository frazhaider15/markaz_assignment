package dto

import "fmt"

func (r *OrderRequest) Validate() error {
	if r.UserId == 0 {
		return fmt.Errorf("invalid user id ")
	}
	if len(r.Products) == 0 {
		return fmt.Errorf("no products")
	}
	return nil
}
