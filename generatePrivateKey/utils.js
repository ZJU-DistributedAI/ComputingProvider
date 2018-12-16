var keyth = require("keythereum");
var Web3 = require("web3");
// console.log(keyth)
// console.log(Web3)

const KEYPATH="/Users/liulifeng/Workspaces/privateChain/data0"


var web3=new Web3(new Web3.providers.HttpProvider("http://localhost:8545"))
console.log(web3.eth.coinbase)
var keyobj=keyth.importFromFile(web3.eth.coinbase, KEYPATH);
var privateKey=keyth.recover("123456", keyobj);
console.log(privateKey.toString('hex'));


// var keythereum = require("keythereum");
// var datadir = "/Users/liulifeng/Workspaces/privateChain/data0";
// var address= "9893e46b95e70035cf11c103d5ca425166b0532b";//要小写
// const password = "123456";
// var keyObject = keythereum.importFromFile(address, datadir);
// var privateKey = keythereum.recover(password, keyObject);
// console.log(privateKey.toString('hex'));