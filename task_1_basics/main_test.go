package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	tests := []struct {
		name        string
		itemPrice   float64
		money       float64
		wantSuccess bool
		wantAmount  float64
	}{
		{
			name:        "exact payment returns true and zero change",
			itemPrice:   25000,
			money:       25000,
			wantSuccess: true,
			wantAmount:  0,
		},
		{
			name:        "overpayment returns true and correct change",
			itemPrice:   25000,
			money:       30000,
			wantSuccess: true,
			wantAmount:  5000,
		},
		{
			name:        "underpayment returns false and correct shortfall",
			itemPrice:   25000,
			money:       20000,
			wantSuccess: false,
			wantAmount:  5000,
		},
		{
			name:        "large overpayment",
			itemPrice:   1500,
			money:       100000,
			wantSuccess: true,
			wantAmount:  98500,
		},
		{
			name:        "small shortfall",
			itemPrice:   10000,
			money:       9999.50,
			wantSuccess: false,
			wantAmount:  0.50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSuccess, gotAmount := processPayment(tt.itemPrice, tt.money)

			if gotSuccess != tt.wantSuccess {
				t.Errorf("processPayment(%v, %v) success = %v, want %v",
					tt.itemPrice, tt.money, gotSuccess, tt.wantSuccess)
			}

			if math.Abs(gotAmount-tt.wantAmount) > 0.001 {
				t.Errorf("processPayment(%v, %v) amount = %v, want %v",
					tt.itemPrice, tt.money, gotAmount, tt.wantAmount)
			}
		})
	}
}

func TestReadPositiveFloat(t *testing.T) {
	tests := []struct {
		name      string
		input     string // simulated stdin lines, separated by \n
		wantValue float64
	}{
		{
			name:      "valid number on first try",
			input:     "25000\n",
			wantValue: 25000,
		},
		{
			name:      "invalid then valid",
			input:     "abc\n15000\n",
			wantValue: 15000,
		},
		{
			name:      "empty then negative then valid",
			input:     "\n-500\n42000\n",
			wantValue: 42000,
		},
		{
			name:      "zero then valid",
			input:     "0\n7500\n",
			wantValue: 7500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			scanner := bufio.NewScanner(reader)

			got := readPositiveFloat(scanner, "test prompt: ")

			if math.Abs(got-tt.wantValue) > 0.001 {
				t.Errorf("readPositiveFloat() = %v, want %v", got, tt.wantValue)
			}
		})
	}
}
