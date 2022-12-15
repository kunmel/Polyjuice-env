package config

// NOTICE: 在本project中，所有TxType，AccessID应遵循从1开始的规律

const TEST_RUN_TIME = 30

const MICRO_TXTYPECOUNT = 10
const MICRO_ACCESSEACH = 8
const MICRO_BACKOFF_TYPE = 3
const MICRO_ACTION_NUM = 3 + MICRO_TXTYPECOUNT
const MICRO_ACCESS_SUM = MICRO_ACCESSEACH * MICRO_TXTYPECOUNT

const THREAD_COUNT = 4

//合约地址，在部署合约时，使用contract.address可以看到
const ETH_CONTRACT_ADDR = "0x4cc9483998ffce05bb0c5fac0a8b04b44535f1ed"

//你的账户的密码
const ETH_ACCOUNT_PASSWD = "0000"

// 下列两个参数组成的地址+文件是你的秘钥文件，地址需换成你所部署的以太坊/data/keystore地址，
// filename为keystore文件夹下文件名的最后一段是你账户地址的文件。
const ETH_KEY_FILENAME = "UTC--2022-12-15T08-21-49.022019637Z--290baa5de37f48760573f74d67433dc568f9f7c5"
const ETH_KEY_FILEPATH = "/home/guotiezheng/go/src/eth-private/data/keystore/"
const ETHCHAINID = 981106

const POLICYPATHFRONT = ""
const TXFILE_PATH = "./transaction/txFile/"
const TXFILE_SIZE = 10000

const RunTime = 100
