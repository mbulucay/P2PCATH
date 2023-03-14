import { Component, OnInit, Input } from '@angular/core';
const {Block} = require('../../../assets/Block.js');

@Component({
  selector: 'app-block-view',
  templateUrl: './block-view.component.html',
  styleUrls: ['./block-view.component.scss']
})
export class BlockViewComponent implements OnInit {

  @Input() public block:typeof Block;

  constructor() { }

  ngOnInit(): void {
  }

}
