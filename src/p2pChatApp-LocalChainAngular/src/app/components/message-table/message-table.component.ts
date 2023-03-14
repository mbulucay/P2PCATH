import { Component, Input, OnInit } from '@angular/core';

const {Message} = require("../../../assets/Message.js")

@Component({
  selector: 'app-message-table',
  templateUrl: './message-table.component.html',
  styleUrls: ['./message-table.component.scss']
})
export class MessageTableComponent implements OnInit {

  @Input() public messages = [] as typeof Message[];;

  constructor() { }

  ngOnInit(): void {
  }

}
