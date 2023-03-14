const SHA256 = require('crypto-js/sha256')
const {Message} = require('./Message.js')

class Block{

    constructor(_messages, _previousHash = ''){
        this.timestamp = Date.now().toString();
        this.previousHash = _previousHash;
        this.messages = _messages;
        this.hash = this.calculateHash();
        this.nonce = 0;
    }

    
    calculateHash(){
        return SHA256(this.timestamp + this.previousHash + JSON.stringify(this.data) + this.nonce).toString();
    }


    mineBlock(difficulty){
        // while(this.hash.substring(0, difficulty) !== Array.from({ length: difficulty }).map(() => 0)){
        while(this.hash.substring(0, difficulty) !== Array(difficulty + 1).join("0")){
            this.nonce++;
            this.hash = this.calculateHash();
        }

        console.log(`Hash : ${this.hash}`);
    }


    hasValidMessages(){
        for(const msg of this.messages){
            if(!msg.isValid()){
                return false;
            }
        }
        return true;
    }
}

module.exports.Block = Block;
