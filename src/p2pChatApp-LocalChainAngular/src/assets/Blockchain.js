const {Block} = require('./Block.js');
const {Message} = require('./Message.js')

const EC = require('elliptic').ec;
const ec = new EC('secp256k1');


class Blockchain{

    // Private Key a7a16d3a0626037f0b7efc52ee97df24ea447a0ad206b3cc7743f842d1265b7d
    // Public Key 04c629278469148aca6cd79da4f4e3e8f0248010464b6e8624dfb8d6ef00fb5156f7f6899fb5eeaa78bf80b78f54f16fb9128f90b8804c90c43a908e57cb4c5cb3
    constructor(){
        this.blockchain = [this.generateGenesis()]
        this.difficulty = 2;
        this.pendingMessages = [];
    }


    generateGenesis(){
            
        let privateKey = ec.keyFromPrivate("a7a16d3a0626037f0b7efc52ee97df24ea447a0ad206b3cc7743f842d1265b7d");
        let messageAddress = privateKey.getPublic('hex');

        let msg = new Message(messageAddress , "mbulucay", `mbulucay thank you for your services`);
        msg.signMessage(privateKey);

        let arr = []
        arr.push(msg);
        return new Block(arr, "0");
    }


    getSize(){ return this.blockchain.length; }

    
    getLastBlock(){ return this.blockchain[this.blockchain.length - 1]; }

    
    addBlock(newBlock){
        newBlock.previousHash = this.getLastBlock().hash;
        // newBlock.hash = newBlock.calculateHash();
        newBlock.mineBlock(this.difficulty)
        this.blockchain.push(newBlock);
    }


    minePendingMessages(miningRewardAddress){
        let block = new Block(this.pendingMessages, this.getLastBlock().hash);

        block.mineBlock(this.difficulty);
        console.log(`Block successfully mined`);
        console.log(`================================================`);

        // console.log(JSON.stringify(this.blockchain,  null, 5))
        this.blockchain.push(block);
        
        let privateKey = ec.keyFromPrivate("a7a16d3a0626037f0b7efc52ee97df24ea447a0ad206b3cc7743f842d1265b7d");
        let messageAddress = privateKey.getPublic('hex');

        let msg = new Message(messageAddress , miningRewardAddress, `${miningRewardAddress} thank you for your services`);
        msg.signMessage(privateKey);

        this.pendingMessages = [ msg ];
    }


    addMessage(message){

        if(!message.fromAddress || !message.toAddress){
            return new Error(`The message need from address and to address to be valid message `);
        }

        if(!message.isValid()){
            return new Error(`Cannot add invalid message in to chain`);
        }

        this.pendingMessages.push(message);
    }


    getSenderMessagesOfAddress(address){

        let messageList = [];
        for(const block of this.blockchain){

            for(const message of block["messages"]){
                if(message["fromAddress"] === address){
                    messageList.push(message);
                }
            }
        }
        return messageList;
    }


    getReceiverMessagesOfAddress(address){

        let messageList = [];
        for(const block of this.blockchain){
            for(const message of block["messages"]){
                if(message["toAddress"] === address){
                    messageList.push(message);
                }
            }
        }
        return messageList;
    }


    printBlocks(){

        // console.log(JSON.stringify(this.blockchain[1], null, 4));

        for(const block of this.blockchain){
            for (const [key, value] of Object.entries(block)) {
                console.log(key, value);
            }
            console.log();
            console.log("=================================");
            console.log("=================================");
            console.log("=================================");
            console.log();
        }
    }

    
    isValid(){

        for(let i=1; i<this.getSize(); i++){

            if(!this.blockchain[i].hasValidMessages()){
                return false;
            }

            if(this.blockchain[i].hash !== this.blockchain[i].calculateHash()){
                console.log(1);
                console.log(this.blockchain[i].hash);
                console.log(this.blockchain[i].calculateHash());
                return false;
            } 
            if(this.blockchain[i].previousHash !== this.blockchain[i-1].hash){
                console.log(2);
                console.log(this.blockchain[i].hash);
                console.log(this.blockchain[i].calculateHash());
                return false;
            }
        }

        return true;
    }
}

module.exports.Blockchain = Blockchain;