import { Component, OnInit } from '@angular/core';
import { BlockchainService } from 'src/app/services/blockchain.service'

const {Blockchain} = require("../../../assets/Blockchain.js")
const {Block} = require("../../../assets/Block.js")
@Component({
  selector: 'app-blockchain-viewer',
  templateUrl: './blockchain-viewer.component.html',
  styleUrls: ['./blockchain-viewer.component.scss']
})
export class BlockchainViewerComponent implements OnInit {

  public blocks = [];
  public  selectedBlock = new Block();

  constructor(private blockchainService: BlockchainService) { 
    this.blocks = blockchainService.getBlocks();
    this.selectedBlock = this.blocks[0];
  }

  ngOnInit(): void {

  }

  showMessages(block: typeof Block) {
    this.selectedBlock = block;
  }
}
