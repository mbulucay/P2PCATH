const {Blockchain} = require('./Blockchain.js');
const {Block} = require('./Block.js');
const {Message} = require('./Message.js')

const EC = require('elliptic').ec;
const ec = new EC('secp256k1');



let chain = new Blockchain();


/* 
// console.log(`Mining 1`);
// chain.addBlock(new Block({sender: "ali", receiver: "veli", content:"naber 1"}))
// chain.addBlock(new Block({sender: "ali", receiver: "veli", content:"nasilsin 1"}))


// console.log(`Mining 2`);
// chain.addBlock(new Block({sender: "veli", receiver: "ali", content:"iyi"}))
// chain.addBlock(new Block({sender: "veli", receiver: "ali", content:"seni sormali"}))
// chain.addBlock(new Block({sender: "veli", receiver: "ali", content:"sen nasilsin"}))

// console.log(JSON.stringify(chain, null, 4));
// console.log(`Is chain valid? : ${chain.isValid()}`);
 */



/* 
// chain.addMessage(new Message ("ali", "veli", "naber AA"));
// chain.addMessage(new Message ("ali", "veli", "nasilsin AA"));

// chain.addMessage(new Message ("veli", "ali", "iyi VV"))

// chain.minePendingTransaction("bedo");

// chain.addMessage(new Message ("veli", "ali", "iyi VV"))
// chain.addMessage(new Message ("veli", "ali", "seni sormali VV"))
// chain.addMessage(new Message ("veli", "ali", "sen nasilsin VV"))

// chain.minePendingTransaction("veli");

// chain.addMessage(new Message ("ali", "veli", "Iyi bizde AA"));
// chain.addMessage(new Message ("ali", "veli", "Hafta sonu mac var AA"));

// chain.minePendingTransaction("ali");

// chain.addMessage(new Message ("ali", "veli", "Gelir misin AA"));
// chain.addMessage(new Message ("veli", "ali", "hangi gun saat kacta VV"))


// chain.minePendingTransaction("ali");

// chain.printBlocks()
// console.log(chain.getSenderMessagesOfAddress("veli"));
// console.log(chain.getReceiverMessagesOfAddress("veli"));

 */

// User 1
// Private Key 475c4ec522314337d3dc33bffdc66a30fd11a3deb747de5f5329d4bb5e998a33
// Public Key 0469342a48c4ce5516c3f593a8d2894a2fd3f7fcdd2ed9a3e3d3069fec5c1743a95cd0b77bda67f87ea82f7316105a1ccb273b0e9bf04a220c419bc535ee79d5b2

const privateKeyU1 = ec.keyFromPrivate("475c4ec522314337d3dc33bffdc66a30fd11a3deb747de5f5329d4bb5e998a33");
const messageAddressU1 = privateKeyU1.getPublic('hex');


// User 2
// Private Key 2bc7688cf201deea3c0266889f86c922fa6dfce64248f83505b535a43556635a
// Public Key 0433f3a4fc76ff47da4a27d64760acd630e29c1660c41fdddc054f3b223d2f5651be6ce7b57d53df5cc5f4c29e01661052f22bf98eb6e853f87fe640400616c18d

const privateKeyU2 = ec.keyFromPrivate("2bc7688cf201deea3c0266889f86c922fa6dfce64248f83505b535a43556635a");
const messageAddressU2 = privateKeyU2.getPublic('hex');

const msg1 = new Message(fromAddress = messageAddressU1, toAddress = messageAddressU2, content = "naber");
msg1.signMessage(privateKeyU1);
chain.addMessage(msg1);

chain.minePendingTransaction(messageAddressU1);

const msg2 = new Message(fromAddress = messageAddressU2, toAddress = messageAddressU1, content = "iyilik senden naber");
msg2.signMessage(privateKeyU2);
chain.addMessage(msg2);

chain.minePendingTransaction(messageAddressU2)

chain.printBlocks()

console.log(chain.isValid());


chain.blockchain[1].messages[0].content = "Dostum olmaz hasmim yasamaz";
console.log(chain.isValid());

