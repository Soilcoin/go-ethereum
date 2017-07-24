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
	"enode://058730ce2c546d59292c33da3ae2347e8588ff28e4f35645d5280006129f5b6f99efa8075fde9db887726436ea802f6970eac0223cc31ee864a14c5a7a3b914d@46.200.159.232:39420",
	"enode://275e82030d664c13459bfe7cde8b51ae8dbdad5f6383c737752b8f004b50faf26f251c8afbb4f936c9909af2f29ae9532fc05466b787e5d4cbe2b8698122055e@80.99.252.96:39420",
	"enode://6d3a030b4bb68e559085e34fc80abc53371cc89e62cd806e0f9feefdd175f672f68866d3fa007d994390485d872bd58a5f2008dd6a67a741c170dc64b4bf9b1a@213.136.76.42:39420",
	"enode://7968705a5a1c9c84a64e1b3985e8605140a10bc8e68e9fd867605a7385b880b7ce945421fc92ad35a5caa573b988a22f5382cf478f8bc785d93324a5b89b71aa@37.187.76.123:39420",
	"enode://aebece2fb7897d71e15146fadf3d5dca5fbd3a67ab8e900423befab250c0a79543cd6c66f62f1937b1f2c0268e83d1e7589effba24dd9bfbe68f383438ec4458@77.78.26.74:39420",
	"enode://b9456fcc4d24602dd1306fae582f0625378519acc23af4d178424c00d911baaadd899fb96b1e431e5ca66245087b641324208f3058839644bacb64b56814b115@51.15.130.151:39420",
	"enode://cb4be3c752e7a231059ab427074af59b07a463c03c470ef34d2fd74c127487db65c4042866fc5445724096e6d9522a6a707fbf2fb02e830a3334ee422e6d0f0a@37.120.184.125:39420",
	"enode://d0708361c50e38eacb8f79d3956eefa0f1f9800e99a4759461a690568bd73e3b48b6e72edb11f51a8d80d1c6bbe9819f0334210a0d7c70a7691b344e84eea2c9@37.187.76.123:39420",
	"enode://46fce6fdc328596df97c8c2e4d1744b9bf5bfd540ff8cf95e4880d797bd92e2cb52a81f38a9a49e51079f875c1adbdc2d4d6f503253e2d7188a494d421bd968b@178.62.133.174:30403",	

}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://6ce05930c72abc632c58e2e4324f7c7ea478cec0ed4fa2528982cf34483094e9cbc9216e7aa349691242576d552a2a56aaeae426c5303ded677ce455ba1acd9d@13.84.180.240:30303", // US-TX
	"enode://20c9ad97c081d63397d7b685a412227a40e23c8bdc6688c6f37e97cfbc22d2b4d1db1510d8f61e6a8866ad7f0e17c02b14182d37ea7c3c8b9c2683aeb6b733a1@52.169.14.227:30303", // IE
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303?discport=30304", // IE
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://0cc5f5ffb5d9098c8b8c62325f3797f56509bff942704687b6530992ac706e2cb946b90a34f1f19548cd3c7baccbcaea354531e5983c7d1bc0dee16ce4b6440b@40.118.3.223:30305",
	"enode://1c7a64d76c0334b0418c004af2f67c50e36a3be60b5e4790bdac0439d21603469a85fad36f2473c9a80eb043ae60936df905fa28f1ff614c3e5dc34f15dcd2dc@40.118.3.223:30308",
	"enode://85c85d7143ae8bb96924f2b54f1b3e70d8c4d367af305325d30a61385a432f247d2c75c45c6b4a60335060d072d7f5b35dd1d4c45f76941f62a4f83b6e75daaf@40.118.3.223:30309",
}
