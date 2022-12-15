package ethConnecter

import (
	"Polyjuice-env/config"
	smartcontract "Polyjuice-env/smartContract"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wonderivan/logger"
)

var initOnce sync.Once

// 这个是智能合约编译出来的golang文件中的，在编译时输入的-type的参数(一般为首字母大写)+Raw，建议自己去那个文件里搜一下，确定一下大小写
var ethTxRaw smartcontract.StoreAndGetRaw
var client *ethclient.Client
var nonceLock sync.Mutex
var nonceFromEth int64
var fromAddress common.Address

func init() {
	initOnce.Do(initClient)
}

func initClient() {
	var err error
	//自己的以太坊的ip及端口，端口在启动网络时设定
	client, err = ethclient.Dial("http://127.0.0.1:7777")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("eth network connected")
	// 智能合约的地址
	address := common.HexToAddress(config.ETH_CONTRACT_ADDR)
	// 这两个函数需要去智能合约生成的golang文件中查看一下。
	// 一般情况下为"New+编译时输入的-type的参数"
	storeObj, err := smartcontract.NewStoreAndGet(address, client)
	// 一般情况下为"编译时输入的-type的参数(一般为首字母大写)+Raw"
	ethTxRaw = smartcontract.StoreAndGetRaw{storeObj}
}

func newAuth() *bind.TransactOpts {
	keyJson, err := ioutil.ReadFile(config.ETH_KEY_FILEPATH + config.ETH_KEY_FILENAME)
	if err != nil {
		logger.Fatal(err)
	}
	key, err := keystore.DecryptKey(keyJson, config.ETH_ACCOUNT_PASSWD)
	if err != nil {
		logger.Fatal(err)
	}
	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logger.Fatal("error casting public key to ECDSA", err)
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	//nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	//gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		logger.Fatal(err)
	}
	//auth := bind.NewKeyedTransactor(key.PrivateKey)
	// 搭建启动以太坊时确认的chainID
	chainID := big.NewInt(config.ETHCHAINID)
	//新版需要chainID，因此不再使用NewKeyedTransactor
	// 生成auth用于transact调用
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		logger.Fatal(err)
	}
	//auth.Nonce = big.NewInt(int64(nonce))
	auth.Nonce = big.NewInt(getNonceFromETH())
	auth.Value = big.NewInt(0)
	// 发起事务如果出现错误可能是因为gas不足，设置大一点就好了
	// 但也不可设置过大，因为在启动链时设置了gas limit，超过这个也不行
	auth.GasLimit = uint64(300000000)
	auth.GasPrice = big.NewInt(1)
	return auth
}

// 分别使用以下两个函数调用智能合约，其中只有调用Call、Transact的区别
// 这两个函数实际上不以是读还是写来进行区分，而是在编写智能合约.sol时，你在编写对应函数时，是否给函数加了view字段来区分
// 如果你添加了view，则你可以在生成的.go文件中，同名函数前发现 *** is a free data retrieval call binding the contract method注释，则你可以使用Call
// 如果没添加，可以发现 *** is a paid mutator transaction binding the contract method，则可以使用Transact

func QueryETH(key string) (string, string) {
	var result []interface{}
	err := ethTxRaw.Call(nil, &result, "retrieve", key)
	if err != nil {
		logger.Error("call query of eth failed", err)
	}
	//fmt.Println("query done")
	value := result[1].(string)
	version := result[2].(string)
	return value, version
}

func BatchQueryETH(keys []string) ([]string, []string) {
	var result []interface{}
	err := ethTxRaw.Call(nil, &result, "batchRetrieve", keys)
	if err != nil {
		logger.Error("call query of eth failed", err)
	}
	return result[0].([]string), result[1].([]string)
}

func WriteETH(key string, value string, version string) error {
	auth := newAuth()
	_, err := ethTxRaw.Transact(auth, "store", key, value, version)
	if err != nil {
		fmt.Println(err)
		return err
	}
	logger.Trace("Write to blockChain success: ", key)
	return nil
}

func BatchWriteETH(key []string, value []string, version []string) error {
	nonceLock.Lock()
	auth := newAuth()
	_, err := ethTxRaw.Transact(auth, "batchStore", key, value, version)
	nonceLock.Unlock()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InitETH() {
	value := "dummyValue"

	// 一定要注意这里的gas设置数量，上链内容多的话，gas很容易超过。
	for j := 0; j <= 200; j++ {
		keys, values, versions := []string{}, []string{}, []string{}
		for n := j * 250; n < (j+1)*250; n++ {
			newKey := strconv.Itoa(n)
			keys = append(keys, newKey)
			values = append(values, value)
			versions = append(versions, strconv.Itoa(-1))
		}
		BatchWriteETH(keys, values, versions)
		fmt.Println(strconv.Itoa(j*200) + "~" + strconv.Itoa(j*200+200) + "write done")
	}

	for i := 1; i <= config.MICRO_TXTYPECOUNT; i++ {
		key := strconv.Itoa(i) + "-"
		for j := 0; j <= 10; j++ {
			keys, values, versions := []string{}, []string{}, []string{}
			for n := j * 100; n < (j+1)*100; n++ {
				keys = append(keys, key+strconv.Itoa(n))
				values = append(values, value)
				versions = append(versions, strconv.Itoa(-1))
			}
			//fmt.Println(keys)
			BatchWriteETH(keys, values, versions)
			fmt.Println(key + strconv.Itoa(j*100) + "~" + strconv.Itoa(j*100+100) + "write done")
		}
	}
	fmt.Println("All kv initialized")
}

func getNonceFromETH() int64 {
	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	return int64(nonce)
}
