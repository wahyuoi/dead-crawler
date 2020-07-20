package main

import "testing"

func TestCompletingLink(t *testing.T) {
	t.Run("Both OK", func(t *testing.T) {
		exp := "https://gedewahyu.com"
		link, err := completeLink(exp, exp)
		if err != nil {
			t.Fail()
		}

		if link != exp {
			t.Fail()
		}
	})
	t.Run("Without schema", func(t *testing.T) {
		exp := "https://gedewahyu.com/gedewahyu.com"
		parent := "https://gedewahyu.com/"
		curr := "gedewahyu.com"
		link, err := completeLink(parent, curr)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if link != exp {
			t.Fatalf("exp: %v; got: %v", exp, link)
		}
	})
	t.Run("Without host, with /", func(t *testing.T) {
		exp := "https://gedewahyu.com/com"
		parent := "https://gedewahyu.com/asd"
		curr := "/com"
		link, err := completeLink(parent, curr)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if link != exp {
			t.Fatalf("exp: %v; got: %v", exp, link)
		}
	})
	t.Run("Without host, without /", func(t *testing.T) {
		exp := "https://gedewahyu.com/asd/com"
		parent := "https://gedewahyu.com/asd"
		curr := "com"
		link, err := completeLink(parent, curr)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if link != exp {
			t.Fatalf("exp: %v; got: %v", exp, link)
		}
	})
	t.Run("Without host, without /, parent is html", func(t *testing.T) {
		exp := "https://gedewahyu.com/asd/com"
		parent := "https://gedewahyu.com/asd/index.html"
		curr := "com"
		link, err := completeLink(parent, curr)
		if err != nil {
			t.Fatalf("%v", err)
		}

		if link != exp {
			t.Fatalf("exp: %v; got: %v", exp, link)
		}
	})
}
