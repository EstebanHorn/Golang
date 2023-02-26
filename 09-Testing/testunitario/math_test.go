package testunitario

import "testing"

/*
func TestSuma(t *testing.T) {
	total := Suma(5, 6)

	if total != 11 {
		t.Errorf("Suma incorrecta tiene %d se esperaba %d", total, 11)
	}*}*/

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)
		if total != item.c {
			t.Errorf("Suma incorrecta tiene %d se esperaba %d", total, item.c)
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{5, 3, 5},
		{5, 6, 6},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("Maximo incorrecto, dio %d se esperaba %d", max, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		a int
		b int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		f := Fibonacci(item.a)
		if f != item.b {
			t.Errorf("Fibonacci incorrecto, dio %d se esperaba %d", f, item.b)
		}
	}
}
