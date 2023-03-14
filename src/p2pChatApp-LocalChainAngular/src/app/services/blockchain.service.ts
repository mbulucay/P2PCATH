import { Injectable } from '@angular/core';

const EC = require('elliptic').ec;

const {Block} = require('../../assets/Block.js');
const {Blockchain} = require('../../assets/Blockchain.js');
const {Message} = require('../../assets/Message.js');


@Injectable({
  providedIn: 'root'
})
export class BlockchainService {

  public blockchain = new Blockchain();
  public walletKeys : any[] = [];

  constructor() { 

    this.blockchain.difficulty = 2;
    this.generateWalletKeys();
  }

  getBlocks(){
    return this.blockchain.blockchain;
  }

  addMessage(message:typeof Message){
    this.blockchain.addMessage(message);
    // this.blockchain.minePendingMessages("Owner Of The System Address");
  }

  getPendingMessages(){
    return this.blockchain.pendingMessages;
  }

  minePendingMessages(){
    this.blockchain.minePendingMessages(this.walletKeys[0].publicKey);
  }

  private generateWalletKeys(){
    const ec = new EC('secp256k1');
    const keyPair = ec.genKeyPair();
    
    this.walletKeys.push({
      keyObj: keyPair,
      publicKey: keyPair.getPublic('hex'),
      privateKey: keyPair.getPrivate('hex'),
    });

    console.log(this.walletKeys);

  }
}
