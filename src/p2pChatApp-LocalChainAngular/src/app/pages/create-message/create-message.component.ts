import { Component, OnInit } from '@angular/core';
import { BlockchainService } from 'src/app/services/blockchain.service';

const {Message} = require('../../../assets/Message.js');


@Component({
  selector: 'app-create-message',
  templateUrl: './create-message.component.html',
  styleUrls: ['./create-message.component.scss']
})
export class CreateMessageComponent implements OnInit {

  public newMsg: typeof Message;
  public walletKey: any;

  constructor(private blockchainService : BlockchainService) {

    this.walletKey = blockchainService.walletKeys[0];


   }

  ngOnInit(): void {
  
    this.newMsg = new Message();
  }

  createMessage(){
    this.newMsg.fromAddress = this.walletKey.publicKey;
    this.newMsg.signMessage(this.walletKey.keyObj);

    this.blockchainService.addMessage(this.newMsg);

    this.newMsg = new Message();
  }

}
