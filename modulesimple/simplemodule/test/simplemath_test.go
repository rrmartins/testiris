package simplemath

import (
  "testing"
  "imartins/modulesimple/simplemodule/simplemath"
)

func TestAddOK(t *testing.T) {
	t.Log("Called simplemath class and Add() method and calc is OK")
  if c := simplemath.Add(5,8); c != 13 {
		t.Errorf("Expected c of 13, but it was %d instead.", c)
	}
}

func TestSubtractOK(t *testing.T) {
	t.Log("Called simplemath class and Subtract() method and calc is OK")
  if c := simplemath.Subtract(5,8); c != -3 {
		t.Errorf("Expected c of -3, but it was %d instead.", c)
	}
}

func TestMultiplyOK(t *testing.T) {
	t.Log("Called simplemath class and Multiply() method and calc is OK")
  if c := simplemath.Multiply(2,2); c != 4 {
		t.Errorf("Expected c of 4, but it was %d instead.", c)
	}
}
