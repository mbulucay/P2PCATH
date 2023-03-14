package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func ImportPublicKeyFromPemStr(pemStr string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pemStr))
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	return pub
}

func ImportPrivateKeyFromPemStr(pemStr string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemStr))
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv
}

func ExportPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	pubASN1, _ := x509.MarshalPKIXPublicKey(pubkey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	}))
}

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

// Export public key as a string in PEM format
func exportPubKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubkey),
		},
	))
	return pubKeyPem
}

// Export private key as a string in PEM format
func exportPrivKeyAsPEMStr(privkey *rsa.PrivateKey) string {
	privKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privkey),
		},
	))
	return privKeyPem

}
