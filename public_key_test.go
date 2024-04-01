package steemgo_test

import (
	"testing"

	steemgo "github.com/berkantsoytas/steem-go"
)

func TestPublicKey_FromString(t *testing.T) {
	for _, td := range keystd {
		p := new(steemgo.PublicKey)

		if err := p.FromString(td["PublicKey"]); err != nil {
			t.Error(err)
		}

		expected := td["PublicKey"]
		actual := p.ToString()

		if expected != actual {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	}
}

func TestPublicKey_FromWIF(t *testing.T) {
	for _, td := range keystd {
		p := new(steemgo.PublicKey)

		if err := p.FromWIF(td["WIF"]); err != nil {
			t.Error(err)
		}

		expected := td["PublicKey"]
		actual := p.ToString()

		if expected != actual {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	}
}
