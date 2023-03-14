const SHA256 = require('crypto-js/sha256')
const EC = require('elliptic').ec;
const ec = new EC('secp256k1');
const EncryptRsa = require('encrypt-rsa').default;

class Message{

    constructor(fromAddress, toAddress, content){
        this.fromAddress = fromAddress;
        this.toAddress = toAddress;
        this.content = content;
    }


    calculateHash(){
        return SHA256(this.fromAddress + this.toAddress + this.content).toString();
    }

    signMessage(signKey){
        if(signKey.getPublic('hex') !== this.fromAddress){ throw new Error(`You can not sign for other address !!!`); }

        const hashMsg = this.calculateHash();
        const sig = signKey.sign(hashMsg, 'base64');
        this.signature = sig.toDER('hex');
        
    }



    isValid(){
        if(this.fromAddress === null) return true;

        if(!this.signature || this.signature === 0){
            throw new Error(`No signature in this message`)
        }

        const publicKey = ec.keyFromPublic(this.fromAddress, 'hex');
        return publicKey.verify(this.calculateHash(), this.signature);
    }
}

module.exports.Message = Message;
