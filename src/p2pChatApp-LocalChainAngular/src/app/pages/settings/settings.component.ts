import { Component, OnInit } from '@angular/core';
import { BlockchainService } from 'src/app/services/blockchain.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent implements OnInit {

  public blockchain;

  constructor(private blockchainServices: BlockchainService) { 

    this.blockchain = blockchainServices.blockchain;

  }

  ngOnInit(): void {
  }

}
