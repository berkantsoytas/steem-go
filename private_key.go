package steemgo

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

type PrivateKey struct {
	Key *btcutil.WIF
}

func (p *PrivateKey) FromWIF(wif string) error {
	wifKey, err := btcutil.DecodeWIF(wif)
	if err != nil {
		return fmt.Errorf("decode WIF failed: %v", err)
	}

	p.Key = wifKey

	return nil
}

func (p *PrivateKey) ToWIF() string {
	return p.Key.String()
}

func (p *PrivateKey) FromString(key string) error {
	return p.FromWIF(key)
}

func (p *PrivateKey) ToString() string {
	return p.ToWIF()
}

func (p *PrivateKey) FromBytes(key []byte) error {
	keyWIF, _ := btcec.PrivKeyFromBytes(key)

	privateKey, err := btcutil.NewWIF(keyWIF, &chaincfg.Params{}, true)
	if err != nil {
		return fmt.Errorf("decode private key failed: %v", err)
	}

	p.Key = privateKey

	return nil
}

func (p *PrivateKey) ToBytes() []byte {
	return p.Key.PrivKey.Serialize()
}

func (p *PrivateKey) ToPublicKey() *PublicKey {
	return &PublicKey{
		Key: p.Key.PrivKey.PubKey(),
	}
}

func (p *PrivateKey) ToPublicKeyString() string {
	return p.ToPublicKey().ToString()
}
