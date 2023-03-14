package cli

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"p2pChatApp/blockchain"
	cr "p2pChatApp/crypto"
	"p2pChatApp/network"
	"p2pChatApp/wallet"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address")
	fmt.Println(" printchain -privatekey PRIVATE_KEY - Prints the blocks in the chain")
	fmt.Println(" send -from FROM -to TO -message AMOUNT -mine -publickey PUBLIC_KEY - Send message from one address to another")
	fmt.Println(" createwallet - Creates a new Wallet")
	fmt.Println(" listaddresses - Lists the addresses in our wallet file")
	fmt.Println(" reindexutxo - Rebuilds the UTXO set")
	fmt.Println(" startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining")
	fmt.Println(" search -keyword KEYWORD -privatekey PRIVATE_KEY - Search for messages containing the keyword")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) StartNode(nodeID, minerAddress string) {
	fmt.Printf("Starting Node %s\n", nodeID)

	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	network.StartServer(nodeID, minerAddress)
}

func (cli *CommandLine) reindexUTXO(nodeID string) {
	chain := blockchain.ContinueBlockChain(nodeID)
	defer chain.Database.Close()
	UTXOSet := blockchain.UTXOSet{Blockchain: chain}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}

func (cli *CommandLine) listAddresses(nodeID string) {
	wallets, _ := wallet.CreateWallets(nodeID)
	addresses := wallets.GetAllAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}

}

func (cli *CommandLine) createWallet(nodeID string) {
	wallets, _ := wallet.CreateWallets(nodeID)
	address := wallets.AddWallet()
	wallets.SaveFile(nodeID)

	fmt.Printf("New address is: %s\n", address)
}

func (cli *CommandLine) printChain(nodeID string, privateKey *rsa.PrivateKey) {
	chain := blockchain.ContinueBlockChain(nodeID)
	defer chain.Database.Close()
	iter := chain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			if privateKey == nil {
				TXString(tx)
			} else {
				decodeTX(tx, privateKey)
			}
		}
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func TXString(tx *blockchain.Transaction) {
	fmt.Printf("--- Transaction %x:", tx.ID)
	fmt.Println()

	for i, input := range tx.Inputs {
		fmt.Printf("     Input %d:", i)
		fmt.Println()
		fmt.Printf("     TXID %d:", input.ID)
		fmt.Println()
		fmt.Printf("     Out %d:", input.Out)
		fmt.Println()
		fmt.Printf("     Signature %x:", input.Signature)
		fmt.Println()
		fmt.Printf("     PubKey %x:", input.PubKey)
		fmt.Println()
	}

	for i, output := range tx.Outputs {
		fmt.Printf("	 Output %d", i)
		fmt.Println()
		fmt.Printf("	 Value  %s", output.Value)
		fmt.Println()
		fmt.Printf("	 Pub Key Hash: %s", string(output.PubKeyHash))
		fmt.Println()
	}
}

func decodeTX(tx *blockchain.Transaction, privateKey *rsa.PrivateKey) {
	fmt.Printf("--- Transaction %x:", tx.ID)
	fmt.Println()

	for i, input := range tx.Inputs {
		fmt.Printf("     Input %d:", i)
		fmt.Println()
		fmt.Printf("     TXID %d:", input.ID)
		fmt.Println()
		fmt.Printf("     Out %d:", input.Out)
		fmt.Println()
		fmt.Printf("     Signature %x:", input.Signature)
		fmt.Println()
		fmt.Printf("     PubKey %x:", input.PubKey)
		fmt.Println()
	}
	fmt.Println()

	for i, output := range tx.Outputs {
		fmt.Printf("Output %d", i)
		fmt.Println()
		if privateKey == nil {
			fmt.Printf("Value:  %s", string(output.Value))
		} else {
			message := []byte(output.Value)
			decMessage, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, message, nil)
			fmt.Printf("Value:  %s", string(decMessage))
		}

		fmt.Println()
		fmt.Printf("Pub Key Hash: %s", string(output.PubKeyHash))
		fmt.Println()
	}

}

func (cli *CommandLine) searchNode(keyword string, nodeID string, privateKey *rsa.PrivateKey) {
	chain := blockchain.ContinueBlockChain(nodeID)
	defer chain.Database.Close()
	iter := chain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			if searchKeyword(tx, keyword, privateKey) {
				decodeTX(tx, privateKey)
			}
		}
		fmt.Println()
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func searchKeyword(tx *blockchain.Transaction, keyword string, privatekey *rsa.PrivateKey) bool {

	contains := false
	for _, output := range tx.Outputs {

		if privatekey == nil {
			if strings.Contains(string(output.Value), keyword) {
				contains = true
				break
			}
		} else {
			message := []byte(output.Value)
			decMessage, _ := rsa.DecryptOAEP(sha256.New(), rand.Reader, privatekey, message, nil)
			if strings.Contains(string(decMessage), keyword) {
				contains = true
				break
			}
		}
	}
	return contains
}

func (cli *CommandLine) createBlockChain(address, nodeID string) {
	if !wallet.ValidateAddress(address) {
		log.Panic("Address is not Valid")
	}
	chain := blockchain.InitBlockChain(address, nodeID)
	defer chain.Database.Close()

	UTXOSet := blockchain.UTXOSet{Blockchain: chain}
	UTXOSet.Reindex()

	fmt.Println("Finished!")
}

func (cli *CommandLine) send(from, to string, message string, nodeID string, mineNow bool) {
	if !wallet.ValidateAddress(to) {
		log.Panic("Address is not Valid")
	}
	if !wallet.ValidateAddress(from) {
		log.Panic("Address is not Valid")
	}
	chain := blockchain.ContinueBlockChain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: chain}
	defer chain.Database.Close()

	wallets, err := wallet.CreateWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := blockchain.NewTransaction(&wallet, to, message, &UTXOSet)
	if mineNow {
		cbTx := blockchain.CoinbaseTx(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}
		block := chain.MineBlock(txs)
		UTXOSet.Update(block)
	} else {
		network.SendTx(network.KnownNodes[0], tx)
		fmt.Println("send tx")
	}

	fmt.Println("Success!")
}

func (cli *CommandLine) CreateKeyPair() {
	bits := 1024
	PrivateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	PublicKey := &PrivateKey.PublicKey

	publicKey := cr.ExportPublicKeyAsPemStr(PublicKey)
	privateKey := cr.ExportPrivateKeyAsPemStr(PrivateKey)

	fmt.Println("PUBLIC_KEY=" + publicKey)
	fmt.Println("PRIVATE_KEY=" + privateKey)
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env is not set!")
		runtime.Goexit()
	}

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("createwallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("listaddresses", flag.ExitOnError)
	reindexUTXOCmd := flag.NewFlagSet("reindexutxo", flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet("startnode", flag.ExitOnError)
	searchNodeCmd := flag.NewFlagSet("search", flag.ExitOnError)

	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendMessage := sendCmd.String("message", "No content found", "Message to send")
	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")
	publicKeyPath := sendCmd.String("publickey", "", "Public key of the sender")
	privateKeyPath := printChainCmd.String("privatekey", "", "Private key of the receiver")
	keyword := searchNodeCmd.String("keyword", "", "Keyword to search")
	privateKeyPath2 := searchNodeCmd.String("privatekey", "", "Keyword to search")

	switch os.Args[1] {
	case "reindexutxo":
		err := reindexUTXOCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "startnode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "listaddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createwallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "search":
		err := searchNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createBlockchainAddress, nodeID)
	}

	if printChainCmd.Parsed() {
		if *privateKeyPath == "" {
			cli.printChain(nodeID, nil)
		} else {

			privateKeyPEM := readKeyFromFile(*privateKeyPath)
			privateKey := exportPEMStrToPrivKey(privateKeyPEM)

			cli.printChain(nodeID, privateKey)
		}
	}

	if createWalletCmd.Parsed() {
		cli.createWallet(nodeID)
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses(nodeID)
	}

	if reindexUTXOCmd.Parsed() {
		cli.reindexUTXO(nodeID)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}

		if *publicKeyPath == "" {
			cli.send(*sendFrom, *sendTo, *sendMessage, nodeID, *sendMine)
		} else {
			pubKeyPEM := readKeyFromFile(*publicKeyPath)
			pubKey := exportPEMStrToPubKey(pubKeyPEM)

			message := []byte(*sendMessage)
			cipherText, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, message, nil)
			*sendMessage = string(cipherText[:])

			cli.send(*sendFrom, *sendTo, *sendMessage, nodeID, *sendMine)
		}
	}

	if startNodeCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startNodeCmd.Usage()
			runtime.Goexit()
		}
		cli.StartNode(nodeID, *startNodeMiner)
	}

	if searchNodeCmd.Parsed() {
		fmt.Println("Searching for keyword: " + *keyword)
		if *keyword == "" {
			searchNodeCmd.Usage()
			runtime.Goexit()
		}
		if *privateKeyPath2 == "" {
			cli.searchNode(*keyword, nodeID, nil)
		} else {
			privateKeyPEM := readKeyFromFile(*privateKeyPath2)
			privateKey := exportPEMStrToPrivKey(privateKeyPEM)
			cli.searchNode(*keyword, nodeID, privateKey)
		}

	}
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
