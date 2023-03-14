const EC = require('elliptic').ec;

const ec = new EC('secp256k1');




while(1){
    const key = ec.genKeyPair();
    const publicK = key.getPublic('hex');
    const privateK = key.getPrivate('hex');

    console.log();
    console.log(`Private Key ${privateK}`);

    console.log();
    console.log(`Public Key ${publicK}`);

    console.log();
    console.log("=================================================");
    console.log();
    // break;"
}


