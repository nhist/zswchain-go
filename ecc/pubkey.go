package ecc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/zhongshuwen/gmsm/x509"
	"github.com/eoscanada/eos-go/libbsuite/btcd/btcec"
	"github.com/eoscanada/eos-go/libbsuite/btcutil/base58"
)

var PublicKeyPrefix = "PUB_"
var PublicKeyK1Prefix = "PUB_K1_"
var PublicKeyR1Prefix = "PUB_R1_"
var PublicKeyWAPrefix = "PUB_WA_"
var PublicKeyGMPrefix = "PUB_GM_"
var PublicKeyPrefixCompat = "EOS"
var PublicKeyPrefixK1Output = "PUB_K1_"

var publicKeyDataSize = new(int)

func init() { *publicKeyDataSize = 33 }

type innerPublicKey interface {
	key(content []byte) (*btcec.PublicKey, error)
	prefix() string
	keyMaterialSize() *int
}

type PublicKey struct {
	Curve   CurveID
	Content []byte

	inner innerPublicKey
}

func (p PublicKey) IsEmpty() bool {
	return p.Curve == 0 && p.Content == nil && p.inner == nil
}

func MustNewPublicKeyFromData(data []byte) PublicKey {
	key, err := NewPublicKeyFromData(data)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func NewPublicKeyFromData(data []byte) (out PublicKey, err error) {
	if len(data) <= 0 {
		return out, errors.New("data must have at least one byte, got 0")
	}

	out = PublicKey{
		Curve:   CurveID(data[0]), // 1 byte
		Content: data[1:],         // 33 bytes for K1 & R1 keys, variable size for WA
	}

	switch out.Curve {
	case CurveK1:
		out.inner = &innerK1PublicKey{}
	case CurveR1:
		out.inner = &innerR1PublicKey{}
	case CurveWA:
		out.inner = &innerWAPublicKey{}
	case CurveGM:
		out.inner = &innerGMPublicKey{}
	default:
		return out, fmt.Errorf("unsupported curve prefix %q", out.Curve)
	}

	return out, out.Validate()
}

func MustNewPublicKey(pubKey string) PublicKey {
	key, err := NewPublicKey(pubKey)
	if err != nil {
		panic(err.Error())
	}
	return key
}

func NewPublicKey(pubKey string) (out PublicKey, err error) {
	if len(pubKey) < 8 {
		return out, fmt.Errorf("invalid format")
	}

	// We had a for/loop using a map before, this had two disavantadges. The ordering was
	// not constant so we were not optimizing for the fact that compat keys appears way more
	// often than all others. Also, updating the compat prefix to something else required to
	// update the map, which was harder to maintain.
	//
	// We now have an unrolled for/loop specially ordered so that the most occurring prefix
	// is checked first.

	if strings.HasPrefix(pubKey, PublicKeyPrefixCompat) {
		return newPublicKey(CurveK1, pubKey[len(PublicKeyPrefixCompat):], newInnerK1PublicKey)
	}

	if strings.HasPrefix(pubKey, PublicKeyK1Prefix) {
		return newPublicKey(CurveK1, pubKey[len(PublicKeyK1Prefix):], newInnerK1PublicKey)
	}

	if strings.HasPrefix(pubKey, PublicKeyR1Prefix) {
		return newPublicKey(CurveR1, pubKey[len(PublicKeyR1Prefix):], newInnerR1PublicKey)
	}

	if strings.HasPrefix(pubKey, PublicKeyWAPrefix) {
		return newPublicKey(CurveWA, pubKey[len(PublicKeyWAPrefix):], newInnerWAPublicKey)
	}

	if strings.HasPrefix(pubKey, PublicKeyGMPrefix) {
		return newPublicKey(CurveGM, pubKey[len(PublicKeyGMPrefix):], newInnerGMPublicKey)
	}

	return out, fmt.Errorf("public key should start with %q, %q, %q, %q, or the old %q", PublicKeyK1Prefix, PublicKeyR1Prefix, PublicKeyWAPrefix, PublicKeyGMPrefix, PublicKeyPrefixCompat)
}

func newPublicKey(curveID CurveID, keyMaterial string, innerFactory func() innerPublicKey) (out PublicKey, err error) {
	payload, err := decodePublicKeyMaterial(keyMaterial, curveID)
	if err != nil {
		return out, err
	}

	return PublicKey{Curve: curveID, Content: payload, inner: innerFactory()}, nil
}

func (p PublicKey) Validate() error {
	if p.inner == nil {
		return fmt.Errorf("the inner public key structure must be present, was nil")
	}

	if p.Curve == CurveK1 || p.Curve == CurveR1 || p.Curve == CurveGM {
		size := p.inner.keyMaterialSize()
		if size == nil {
			return fmt.Errorf("R1, K1, GM public keys must have a fixed key material size")
		}

		if len(p.Content) != *size {
			return fmt.Errorf("public key data must have a length of %d, got %d", *size, len(p.Content))
		}
	}

	return nil
}
func SM2PemToEOSPublicKeyString(sm2PublicKeyPemData []byte) (string, error) {
	pem, err := x509.ReadPublicKeyFromPem(sm2PublicKeyPemData)
	if err != nil {
		return "", err
	}

	baseBytes := []byte{byte(CurveGM)}
	baseBytes = append(baseBytes, CompressReal(pem)...)
	pubKey, err := NewPublicKeyFromData(baseBytes)
	if err != nil {
		return "", err
	}
	return pubKey.String(), nil

}
func (p PublicKey) Key() (*btcec.PublicKey, error) {
	return p.inner.key(p.Content)
}

var emptyKeyMaterial = make([]byte, 33)

func (p PublicKey) GetCompoundPublicKeyASN1SignatureData(asn1SignatureData []byte) (*Signature, error) {
	if p.Curve != CurveGM {
		return nil, fmt.Errorf("GetCompoundPublicKeyASN1SignatureData: currenly only supported for SM2 (GM)")
	}
	sigData := []byte{byte(CurveGM)}
	sigData = append(sigData, p.Content[0:33]...)
	sigData = append(sigData, asn1SignatureData...)
	if len(sigData) > 106 {
		return nil, fmt.Errorf("signature too long, total size of the signature including the curve byte and 33 byte compressed public key should be 106 bytes or less")
	} else if len(sigData) < 106 {
		sigData = append(sigData, bytes.Repeat([]byte{0}, 106-len(sigData))...)
	}
	signatureInstance, err := NewSignatureFromData(sigData)
	return &signatureInstance, err
}
func (p PublicKey) String() string {
	if p.IsEmpty() {
		return ""
	}

	data := p.Content
	if len(data) == 0 {
		// Nothing really to do, just output some garbage
		return p.inner.prefix() + base58.Encode(emptyKeyMaterial)
	}

	hash := ripemd160checksum(data, p.Curve)
	size := p.KeyMaterialSize()

	rawKey := make([]byte, size+4)
	copy(rawKey, data[:size])
	copy(rawKey[size:], hash[:4])

	return p.inner.prefix() + base58.Encode(rawKey)
}

func (p PublicKey) KeyMaterialSize() int {
	size := p.inner.keyMaterialSize()
	if size != nil {
		return *size
	}

	return len(p.Content)
}

func (p PublicKey) MarshalJSON() ([]byte, error) {
	s := p.String()
	return json.Marshal(s)
}

func (p *PublicKey) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	newKey, err := NewPublicKey(s)
	if err != nil {
		return err
	}

	*p = newKey

	return nil
}
