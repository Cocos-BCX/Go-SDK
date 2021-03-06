package CocosSDK // import "CocosSDK"

import (
	"CocosSDK/chain"
	"CocosSDK/rpc"
	"CocosSDK/wallet"
	"log"
	"sync"
)

var once *sync.Once = &sync.Once{}
var Wallet *wallet.Wallet
var Chain *chain.Chain

/*
*初始化SDK
 */
func InitSDK(host string, use_ssl bool, port ...int) {
	once.Do(
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Panicln("SDK Init Error:", err)
				}
			}()
			if err := rpc.InitClient(host, use_ssl, port...); err != nil {
				log.Panicln("SDK Init Error:", err)
			}
			chain.InitChain()
			Chain = chain.CocosBCXChain
			Wallet = wallet.CreateWallet()
		})
}
