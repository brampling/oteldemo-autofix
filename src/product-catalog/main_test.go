// SPDX-License-Identifier: Apache-2.0
package main

import (
	"testing"

	pb "github.com/opentelemetry/opentelemetry-demo/src/product-catalog/genproto/oteldemo"
)

func TestValidatePrice(t *testing.T) {
	tests := []struct {
		name    string
		price   *pb.Money
		wantErr bool
	}{
		{name: "whole dollars", price: &pb.Money{CurrencyCode: "USD", Units: 175}},
		{name: "dollars and cents", price: &pb.Money{CurrencyCode: "USD", Units: 129, Nanos: 950_000_000}},
		{name: "zero", price: &pb.Money{CurrencyCode: "USD"}, wantErr: true},
		{name: "negative", price: &pb.Money{CurrencyCode: "USD", Units: -5}, wantErr: true},
		{name: "nanos out of range", price: &pb.Money{CurrencyCode: "USD", Units: 1, Nanos: 1_000_000_000}, wantErr: true},
		{name: "missing", price: nil, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePrice(tt.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePrice(%v) error = %v, wantErr %v", tt.price, err, tt.wantErr)
			}
		})
	}
}
