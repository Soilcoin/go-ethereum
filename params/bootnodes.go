// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{

	"enode://12217da9b5d13d06b4cfec1fb6bec30912a644752c67fe1cab4bb22f7ca10e57f9da9330fcec37f31386ead09cdab62618f8fbcd6fc8fa600543d6a9ded3f8bb@37.187.76.123:39420"
	"enode://17677980aa2fb04f3910b16e552c7ba3b6840f46085f750608f339935d7b6f9b7da113316b908c148e18b31c3933f5caee6ee610827a61f87c8f27838c4e2c6f@95.215.46.238:39420"
	"enode://36667a60f555c254957b56aca013b507bb9e2a2f7b2484be7c70471dee268492907a11ee00329af761574ae2be57562660b95738654a2ca96bcf8ae88974033f@51.15.130.241:34102"
	"enode://92dac473a9f3a50b5f4e91b59fa13087b49bd93fba224dbb12dc812740c337ddf0c3af7407134d18cdfb694c0c60c036ef54320d47f3f9ec9f12b0830a2075a3@201.242.173.58:23486"
	"enode://a1c00725424f0159ae716db7c3236f5aa93fe9a7015592eef24ac563d735c71b09ff27af260271f6ae61bc3ba10656619df4f08961261cf3f10e049d73ef4d1e@37.187.76.123:60505"
	"enode://b9456fcc4d24602dd1306fae582f0625378519acc23af4d178424c00d911baaadd899fb96b1e431e5ca66245087b641324208f3058839644bacb64b56814b115@51.15.130.151:32850"
	"enode://ed65072f342efe1e95b61b6544a7f2a64fceeab4a448ecd56c671015c7ee7721f5e919adf00e63111636136cc36ed1128b1ae2e6da98b52f520404cf3e9284b4@85.62.20.153:48510"

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
}
