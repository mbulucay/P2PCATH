package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
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

// Save string to a file
func saveKeyToFile(keyPem, filename string) {
	pemBytes := []byte(keyPem)
	ioutil.WriteFile(filename, pemBytes, 0400)
}

// Decode private key struct from PEM string
func exportPEMStrToPrivKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

// Decode public key struct from PEM string
func exportPEMStrToPubKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}

// Read data from file
func readKeyFromFile(filename string) []byte {
	key, _ := ioutil.ReadFile(filename)
	return key
}

func generateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	// This method requires a random number of bits.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// The public key is part of the PrivateKey struct
	return privateKey, &privateKey.PublicKey
}

func main() {

	// Generate a 2048-bits key
	privateKey, publicKey := generateKeyPair(2048)

	// fmt.Printf("Private key: %v\n", privateKey)
	// fmt.Printf("Public Key: %v", publicKey)

	// Create PEM string
	privKeyStr := exportPrivKeyAsPEMStr(privateKey)
	pubKeyStr := exportPubKeyAsPEMStr(publicKey)

	fmt.Println(privKeyStr)
	fmt.Println(pubKeyStr)

	saveKeyToFile(privKeyStr, "privkey.pem")
	saveKeyToFile(pubKeyStr, "pubkey.pem")

	// privKeyPEM := readKeyFromFile("privkey.pem")
	// privateKey = exportPEMStrToPrivKey(privKeyPEM)
	// fmt.Printf("Private key: %v\n", privKeyFile)

	// pubKeyPEM := readKeyFromFile("pubkey.pem")
	// publicKey = exportPEMStrToPubKey(pubKeyPEM)
	// fmt.Printf("Public key: %v\n", pubKeyFile)

	// message := []byte("super secret message")
	// cipherText, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, message, nil)
	// fmt.Println("Encrypted message: ", cipherText)

	// decMessage, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, cipherText, nil)
	// fmt.Printf("Original: %s\n", string(decMessage))

}
