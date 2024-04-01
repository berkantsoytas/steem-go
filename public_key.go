package steemgo

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/ripemd160"
)

const (
	// ADDRESS_PREFIX = "your_address_prefix_here"
	ChecksumLength = 4
	PrefixLength   = len(ADDRESS_PREFIX)
)

type PublicKey struct {
	Key *secp256k1.PublicKey
}

func (p *PublicKey) FromBytes(key []byte) error {
	publicKey, err := btcec.ParsePubKey(key)
	if err != nil {
		return fmt.Errorf("parse public key failed: %v", err)
	}

	p.Key = (*secp256k1.PublicKey)(publicKey)

	return nil
}

func (p *PublicKey) FromString(key string) error {
	prefix := key[:PrefixLength]
	if prefix != ADDRESS_PREFIX {
		return fmt.Errorf("invalid address prefix: %s", prefix)
	}

	publicKeyWithoutPrefix := key[PrefixLength:]
	publicKeyBytes := base58.Decode(publicKeyWithoutPrefix)

	publicKeyOriginal := publicKeyBytes[:len(publicKeyBytes)-ChecksumLength]
	checksum := publicKeyBytes[len(publicKeyBytes)-ChecksumLength:]
	cChecksum := hashsum(publicKeyOriginal)

	if !bytes.Equal(checksum, cChecksum[:ChecksumLength]) {
		return fmt.Errorf("invalid checksum: %v", checksum)
	}

	return p.FromBytes(publicKeyOriginal)
}

func (p *PublicKey) ToString() string {
	checksum := hashsum(p.Key.SerializeCompressed())
	key := append(p.Key.SerializeCompressed(), checksum[:ChecksumLength]...)
	return ADDRESS_PREFIX + base58.Encode(key)
}

func (p *PublicKey) ToBytes() []byte {
	return p.Key.SerializeCompressed()
}

func (p *PublicKey) FromPrivateKey(privateKey *PrivateKey) {
	*p = *privateKey.ToPublicKey()
}

func (p *PublicKey) FromWIF(wif string) error {
	privateKey := new(PrivateKey)
	if err := privateKey.FromWIF(wif); err != nil {
		return err
	}

	p.FromPrivateKey(privateKey)

	return nil
}

func hashsum(buff []byte) []byte {
	hasher := ripemd160.New()
	hasher.Write(buff)
	return hasher.Sum(nil)
}
