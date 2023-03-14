// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract p2pChatApp{

    struct Mesaj{
        address from;
        string content;
    }
    mapping(uint => Mesaj) m_mesaj;

    mapping(address => Mesaj[]) public addres_box;

    uint public counter = 0;

    constructor(){

    }

    function addMesaj(address _to , string memory _content) public returns(bool){

        addres_box[_to].push(Mesaj(msg.sender, _content));
        counter++;

        return true;
    }

}