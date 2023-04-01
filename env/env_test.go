package env

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv_String(t *testing.T) {
	e_name := "test"
	orig_e := "abc"

	if err := os.Setenv(e_name, orig_e); err != nil {
		t.Error(err)
	}

	e := Get(e_name, "123")
	if e != orig_e {
		t.Errorf("expected: %v, received: %v", orig_e, e)
	}

	if err := os.Unsetenv(e_name); err != nil {
		t.Error(err)
	}
}

func TestEnv_String_Default(t *testing.T) {
	e_name := "test"
	test_e := "123"

	e := Get(e_name, test_e)
	if e != test_e {
		t.Errorf("expected: %v, received: %v", test_e, e)
	}
}

func TestEnv_Int(t *testing.T) {
	e_name := "test"
	orig_e := 69
	test_e := 42

	if err := os.Setenv(e_name, fmt.Sprint(orig_e)); err != nil {
		t.Error(err)
	}

	e := Get(e_name, test_e)
	if e != orig_e {
		t.Errorf("expected: %v, received: %v", orig_e, e)
	}

	if err := os.Unsetenv(e_name); err != nil {
		t.Error(err)
	}
}

func TestEnv_Int_Default(t *testing.T) {
	e_name := "test"
	test_e := 42

	e := Get(e_name, test_e)
	if e != test_e {
		t.Errorf("expected: %v, received: %v", test_e, e)
	}
}

func TestEnv_Bool(t *testing.T) {
	e_name := "test"
	orig_e := true
	test_e := false

	if err := os.Setenv(e_name, fmt.Sprint(orig_e)); err != nil {
		t.Error(err)
	}

	e := Get(e_name, test_e)
	if e != orig_e {
		t.Errorf("expected: %v, received: %v", orig_e, e)
	}

	if err := os.Unsetenv(e_name); err != nil {
		t.Error(err)
	}
}

func TestEnv_Bool_Default(t *testing.T) {
	e_name := "test"
	test_e := true

	e := Get(e_name, test_e)
	if e != test_e {
		t.Errorf("expected: %v, received: %v", test_e, e)
	}
}
