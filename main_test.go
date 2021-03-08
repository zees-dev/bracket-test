package main

import (
	"testing"
)

func TestValidInputs(t *testing.T) {
	t.Run("test empty input passes", func(t *testing.T) {
		pass := validateBrackets("")
		if !pass {
			t.Errorf("failed to validate empty string")
		}
	})

	t.Run("test {} passes", func(t *testing.T) {
		pass := validateBrackets("{}")
		if !pass {
			t.Errorf("failed to validate `{}`")
		}
	})

	t.Run("test [{{}}] passes", func(t *testing.T) {
		pass := validateBrackets("[{{}}]")
		if !pass {
			t.Errorf("failed to validate `[{{}}]`")
		}
	})

	t.Run("test [{{}()[]}]({}) passes", func(t *testing.T) {
		pass := validateBrackets("[{{}()[]}]({})")
		if !pass {
			t.Errorf("failed to validate `[{{}()[]}]({})`")
		}
	})
}

func TestInValidInputs(t *testing.T) {
	t.Run("test } fails", func(t *testing.T) {
		pass := validateBrackets("}")
		if pass {
			t.Errorf("validate incorrectly `}`")
		}
	})

	t.Run("test [ fails", func(t *testing.T) {
		pass := validateBrackets("[")
		if pass {
			t.Errorf("validate incorrectly `[`")
		}
	})

	t.Run("test [{} fails", func(t *testing.T) {
		pass := validateBrackets("[{}")
		if pass {
			t.Errorf("validate incorrectly `[{}`")
		}
	})

	t.Run("test [{{}()[]}]({}){ fails", func(t *testing.T) {
		pass := validateBrackets("[{{}()[]}]({}){")
		if pass {
			t.Errorf("validate incorrectly `[{{}()[]}]({}){`")
		}
	})
}
