

The program can createing wallet. With the uniqe address on the ./tmp folder for the each user
{
    $ go run main.go createwallet
}


The program can list all the wallets in the ./tmp folder
{
    $ go run main.go listwallets
}


The program can send the message to the address
{
    $ go run main.go send -from FROM_ADDRESS -to TO_ADDRESS -message CONTENT -mine -publickey PUBLIC_KEY_FILE_PATH
}

The program can create the blockchain
{
    $ go run main.go createblockchain -address ADDRESS
}

The program can print the blockchain
{
    $ go run main.go printchain -privatekey PRIVATE_KEY_FILE_PATH
}

The program can create the new node as miner
{
    $ go run main.go startnode -miner ADDRESS
}
