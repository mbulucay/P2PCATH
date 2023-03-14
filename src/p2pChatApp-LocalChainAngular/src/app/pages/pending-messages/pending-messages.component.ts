import { Component, OnInit } from '@angular/core';
import { BlockchainService } from 'src/app/services/blockchain.service';

const { Message } = require("../../../assets/Message.js")

@Component({
  selector: 'app-pending-messages',
  templateUrl: './pending-messages.component.html',
  styleUrls: ['./pending-messages.component.scss']
})
export class PendingMessagesComponent implements OnInit {

  public pendingMessages = [] as typeof Message[];

  constructor(private blockchainService : BlockchainService) { 
    this.pendingMessages = blockchainService.getPendingMessages();
  }

  ngOnInit(): void {
  }

  minePendingMessages(){
    this.blockchainService.minePendingMessages();
  }

}
