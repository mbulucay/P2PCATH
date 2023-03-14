package blockchain

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// string to rsa.PublicKey
func ImportPublicKeyFromPemStr(pemStr string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pemStr))
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	return pub
}

// string to rsa.PrivateKey
func ImportPrivateKeyFromPemStr(pemStr string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemStr))
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv
}

// rsa.PublicKey to string
func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	pubASN1, _ := x509.MarshalPKIXPublicKey(pubkey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}))
}

// rsa.PrivateKey to string
func ExportPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privBytes := x509.MarshalPKCS1PrivateKey(privkey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}))

}

func ExportMsgAsPemStr(msg []byte) string {
	msg_pem := string(pem.EncodeToMemory(&pem.Block{Type: "MESSAGE", Bytes: msg}))
	return msg_pem
}
