package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderA := NewOrder(true, 5)
	buyOrderB := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 10)

	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderC)

	fmt.Println(l)

	l.RemoveOrder(buyOrderB)

	fmt.Println(l)
}

func TestOrderbook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 10)
	buyOrderB := NewOrder(true, 2000)
	buyOrderC := NewOrder(true, 500)

	ob.PlaceOrder(18_000, buyOrderA)
	ob.PlaceOrder(18_000, buyOrderB)
	ob.PlaceOrder(19_000, buyOrderC)

	fmt.Printf("%+v", ob.Bids)
}
