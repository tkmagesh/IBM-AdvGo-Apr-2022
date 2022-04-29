package models

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %v, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func (products Products) Filter(predicate func(product Product) bool) Products {
	result := Products{}
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func Filter[T any](items []T, predicate func(item T) bool) []T {
	result := []T{}
	for _, p := range items {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

func (products Products) All(predicate func(product Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(product Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}
