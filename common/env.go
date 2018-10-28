package common

import (
	"eospart_websocket/module"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

// server
var ServerAddr string
var ReadTimeout int64
var HostName string
var DataPath string
var DataName string
var DataFile string
var P_DataFile *os.File
var DataChannel chan string

// business
var AllowRequestRecord bool
var AllowMessageRecord bool

// coinmarketcap
var CoinmarketcapAddress string
var CoinmarketcapTransport *http.Transport
var CoinmarketcapTimeout int64
var CoinmarketcapTimerInterval int64
var CoinmarketcapClient *http.Client

// coingecko
var CoinGeckoAddress string
var CoinGeckoTransport *http.Transport
var CoinGeckoTimeout int64
var CoinGeckoTimerInterval int64
var CoinGeckoClient *http.Client

// eos
var EOSGetBlockInterface string
var EOSGetMakeUpBlockInterface string
var EOSGetInfoInterface string
var EOSGetBalanceInterface string
var EOSGetCodeInterface string
var EOSGetABIInterface string
var EOSGetRawCodeInterface string
var EOSGetBPInterface string
var EOSGetTableInterface string
var EOSGetAccountResourceInterface string
var EOSGetCurrencyStatsInterface string

var EOSTransport *http.Transport
var EOSTimeout int64
var EOSTimerInterval int64
var EOSFrequencyLimit int64
var FrequencyLimitTimestamp int64
var BPTimerInterval int64
var EOSClient *http.Client
var EOSToolPath string
var ETHMainNetUrl string
var EOSMainSymbol string
var UseBPRPC bool

var CacheServerList string
var UseCache bool
var SemiPermanentCache, PermanentCache map[string]string
var ShortCacheTime, LongCacheTime int64

var InlineTokenExecAccountList string
var BPNodeHostList string

//source code dir
var SourceCodeDir string

var AccountBlackList []string

//mysql
var MysqlCacheTime int64
var MysqlCacheInsertTimestamp int64

var IrreversibleBlockNum int64

var MaxPageSize int64

//contract
var VerifyContractInterface string
var ContractAuditInterface string
var AuditCheckList map[string]string

//指定每日时间更新缓存
var EveryDayTime map[string]int64

var SeverFlag string

//转账记录条数白名单
var WhiteList string

func InitEnv() (Err module.Error) {
	var err error

	ServerAddr = Config.Section("service").Key("addr").
		MustString(CONFIG_DEFAULT_SERVER_ADDR)

	ReadTimeout = Config.Section("service").Key("read_timeout").
		MustInt64(CONFIG_DEFAULT_READ_TIMEOUT)

	DataPath = Path +
		Config.Section("service").Key("data_path").MustString(CONFIG_DEFAULT_DATA_PATH)

	DataName = Config.Section("service").Key("data_name").
		MustString(CONFIG_DEFAULT_DATA_NAME)

	HostName, err = os.Hostname()
	if err != nil {
		Logger.Error("HostName failed")
		return module.Error{ErrCode:ERROR_SYSTEM, ErrMsg:err}
	}

	DataChannelMax := int(Config.Section("service").Key("data_channel_max").
		MustInt64(CONFIG_DEFAULT_DATA_CHANNEL_MAX))

	DataChannel = make(chan string, DataChannelMax)

	AllowRequestRecord =
		Config.Section("business").Key("allow_request_record").MustBool(false)

	AllowMessageRecord =
		Config.Section("business").Key("allow_message_record").MustBool(false)

	MaxPageSize =
		Config.Section("business").Key("max_page_size").MustInt64(DEFAULT_MAX_LIMIT)

	//coinmarketcap
	CoinmarketcapAddress = Config.Section("coinmarketcap").Key("addr").MustString(CONFIG_DEFAULT_COIN_MARKET_CAP_ADDR)
	CoinmarketcapMaxConnection := int(Config.Section("coinmarketcap").Key("max_connection").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_MAX_CONNECTION))
	CoinmarketcapTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     60 * time.Second,
		MaxIdleConns:        CoinmarketcapMaxConnection,
		MaxIdleConnsPerHost: CoinmarketcapMaxConnection,
	}
	CoinmarketcapTimeout = Config.Section("coinmarketcap").Key("timeout").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_TIMEOUT)
	CoinmarketcapTimerInterval = Config.Section("coinmarketcap").Key("timer_interval").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_TIMER_INTERVAL)
	coinmarketcap_timeout := time.Duration(CoinmarketcapTimeout) * time.Second
	CoinmarketcapClient = &http.Client{
		Timeout:   coinmarketcap_timeout,
		Transport: CoinmarketcapTransport,
	}

	//coingecko
	CoinGeckoAddress = Config.Section("coingecko").Key("addr").MustString(CONFIG_DEFAULT_COIN_GECKO_ADDR)
	CoinGeckoMaxConnection := int(Config.Section("coingecko").Key("max_connection").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_MAX_CONNECTION))
	CoinGeckoTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     60 * time.Second,
		MaxIdleConns:        CoinGeckoMaxConnection,
		MaxIdleConnsPerHost: CoinGeckoMaxConnection,
	}
	CoinGeckoTimeout = Config.Section("coingecko").Key("timeout").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_TIMEOUT)
	CoinGeckoTimerInterval = Config.Section("coingecko").Key("timer_interval").MustInt64(CONFIG_DEFAULT_COIN_MARKET_CAP_TIMER_INTERVAL)
	coingecko_timeout := time.Duration(CoinGeckoTimeout) * time.Second
	CoinGeckoClient = &http.Client{
		Timeout:   coingecko_timeout,
		Transport: CoinGeckoTransport,
	}

	EOSGetBlockInterface = Config.Section("eos").Key("get_block").MustString(CONFIG_DEFAULT_EOS_GET_BLOCK_INTERFACE)
	EOSGetMakeUpBlockInterface = Config.Section("eos").Key("make_up_get_block").MustString(CONFIG_DEFAULT_EOS_GET_MAKE_UP_BLOCK_INTERFACE)
	EOSGetInfoInterface = Config.Section("eos").Key("get_info").MustString(CONFIG_DEFAULT_EOS_GET_INFO_INTERFACE)
	EOSGetBalanceInterface = Config.Section("eos").Key("get_balance").MustString(CONFIG_DEFAULT_EOS_GET_BALANCE_INTERFACE)
	EOSGetCodeInterface = Config.Section("eos").Key("get_code").MustString(CONFIG_DEFAULT_EOS_GET_CODE_INTERFACE)
	EOSGetABIInterface = Config.Section("eos").Key("get_abi").MustString(CONFIG_DEFAULT_EOS_GET_ABI_INTERFACE)
	EOSGetRawCodeInterface = Config.Section("eos").Key("get_raw_code").MustString(CONFIG_DEFAULT_EOS_GET_RAW_CODE_INTERFACE)

	EOSGetBPInterface = Config.Section("eos").Key("get_bp").MustString(CONFIG_DEFAULT_EOS_GET_BP_INTERFACE)
	EOSGetTableInterface = Config.Section("eos").Key("get_table").MustString(CONFIG_DEFAULT_EOS_GET_TABLE_INTERFACE)
	EOSGetAccountResourceInterface = Config.Section("eos").Key("get_account").MustString(CONFIG_DEFAULT_EOS_GET_ACCOUNT_RESOURCE_INTERFACE)
	EOSGetCurrencyStatsInterface = Config.Section("eos").Key("get_currency_stats").MustString(CONFIG_DEFAULT_EOS_GET_CURRENCY_STATS_INTERFACE)

	ETHMainNetUrl = Config.Section("eos").Key("get_eth_main_net").MustString(CONFIG_DEFAULT_EOS_GET_ETH_MAIN_NET_URL)
	EOSMainSymbol = Config.Section("eos").Key("main_sys").MustString(CONFIG_DEFAULT_EOS_MAIN_SYS)
	UseBPRPC = Config.Section("eos").Key("use_bp_rpc").MustBool(false)

	CacheServerList = Config.Section("cache").Key("server_list").MustString(CONFIG_CACHE_SVR_LIST)
	UseCache = Config.Section("cache").Key("use_cache").MustBool(false)
	ShortCacheTime = Config.Section("cache").Key("short_cache_time").MustInt64(CONFIG_SHORT_CACHE_TIME)
	LongCacheTime = Config.Section("cache").Key("long_cache_time").MustInt64(CONFIG_LONG_CACHE_TIME)
	MysqlCacheTime = Config.Section("cache").Key("mysql_cache_time").MustInt64(CONFIG_DEFAULT_MYSQL_CACHE_TIME)

	InlineTokenExecAccountList = Config.Section("eos").Key("account_list").MustString(CONFIG_DEFAULT_EOS_ACCOUNT_LIST)
	BPNodeHostList = Config.Section("bp_api_node").Key("server_list").MustString(CONFIG_NODE_API_HOST_LIST)

	IrreversibleBlockNum = Config.Section("eos").Key("irreversible_block_num").MustInt64(CONFIG_DEFAULT_IRREVERSIBLE_BLOCK_NUM)

	EOSMaxConnection := int(Config.Section("eos").Key("max_connection").MustInt64(CONFIG_DEFAULT_EOS_MAX_CONNECTION))
	EOSTimeout = Config.Section("eos").Key("timeout").MustInt64(CONFIG_DEFAULT_EOS_TIMEOUT)
	EOSClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        EOSMaxConnection,
			MaxIdleConnsPerHost: EOSMaxConnection,
			IdleConnTimeout:     time.Duration(30) * time.Second,
		},
		Timeout: time.Duration(EOSTimeout) * time.Second,
	}
	EOSFrequencyLimit = Config.Section("eos").Key("frequency_limit").MustInt64(CONFIG_DEFAULT_EOS_FREQUENCY_LIMIT)
	EOSTimerInterval = Config.Section("eos").Key("timer_interval").MustInt64(CONFIG_DEFAULT_EOS_TIMER_INTERVAL)
	EOSToolPath = Config.Section("eos").Key("tool_path").MustString(CONFIG_DEFAULT_EOS_TOOL_PATH)
	BPTimerInterval = Config.Section("bp").Key("timer_interval").MustInt64(CONFIG_DEFAULT_BP_TIMER_INTERVAL)
	SourceCodeDir = Config.Section("eos").Key("contract_code_path").MustString(CONFIG_CONTRACT_SOURCE_PATH)
	accountBlackList := Config.Section("eos").Key("account_black_list").MustString(CONFIG_ACCOUNT_BLACK_LIST)
	AccountBlackList = strings.Split(accountBlackList, ",")

	VerifyContractInterface = Config.Section("eos").Key("verify_contract_url").MustString(VERIFY_CONTRACT_URL)
	ContractAuditInterface = Config.Section("eos").Key("contract_audit_url").MustString(AUDIT_CONTRACT_URL)
	AuditCheckList = make(map[string]string)

	//cache
	SemiPermanentCache = make(map[string]string)
	PermanentCache = make(map[string]string)
	SemiPermanentCache[INTERFACE_GET_ACCOUNT_BLOCK_INFO] = INTERFACE_GET_ACCOUNT_BLOCK_INFO
	SemiPermanentCache[INTERFACE_GET_BP_INFO] = INTERFACE_GET_BP_INFO
	SemiPermanentCache[INTERFACE_GET_ACCOUNT_INFO] = INTERFACE_GET_ACCOUNT_INFO
	SemiPermanentCache[INTERFACE_GET_CONTRACT_TRX_INFO] = INTERFACE_GET_CONTRACT_TRX_INFO
	SemiPermanentCache[INTERFACE_GET_ACCOUNT_RELATED_TRX_INFO] = INTERFACE_GET_ACCOUNT_RELATED_TRX_INFO
	SemiPermanentCache[INTERFACE_GET_SUB_ACCOUNT_INFO] = INTERFACE_GET_SUB_ACCOUNT_INFO
	SemiPermanentCache[INTERFACE_GET_TOP_BID_LIST] = INTERFACE_GET_TOP_BID_LIST
	SemiPermanentCache[INTERFACE_GET_REFLECT_EOS_ACCOUNT] = INTERFACE_GET_REFLECT_EOS_ACCOUNT
	SemiPermanentCache[INTERFACE_GET_TOKEN_LIST_INFO] = INTERFACE_GET_TOKEN_LIST_INFO
	// SemiPermanentCache[INTERFACE_GET_CONTRACT_INFO] = INTERFACE_GET_CONTRACT_INFO;
	// SemiPermanentCache[INTERFACE_VERIFY_CODE] = INTERFACE_VERIFY_CODE;

	PermanentCache[INTERFACE_GET_BLOCK_DETAIL_INFO] = INTERFACE_GET_BLOCK_DETAIL_INFO
	PermanentCache[INTERFACE_GET_BLOCK_TRANSACTION_INFO] = INTERFACE_GET_BLOCK_TRANSACTION_INFO
	PermanentCache[INTERFACE_GET_TRANSACTION_ACTION_INFO] = INTERFACE_GET_TRANSACTION_ACTION_INFO
	PermanentCache[INTERFACE_GET_TRANSACTION_DETAIL_INFO] = INTERFACE_GET_TRANSACTION_DETAIL_INFO
	PermanentCache[INTERFACE_GET_ADDRESS_DETAIL_INFO] = INTERFACE_GET_ADDRESS_DETAIL_INFO

	EveryDayTime = make(map[string]int64)
	EveryDayTime[INTERFACE_GET_TARGET_TOKEN_DETAILS] = 1
	EveryDayTime[INTERFACE_GET_OVERVIEW_STABLE] = 1

	SeverFlag = Config.Section("server").Key("server_flag").MustString(CONFIG_SERVER_MAIN)

	WhiteList = Config.Section("trx").Key("white_list").MustString("")
	Logger.Info("evn init success.")
	return module.Error{ErrCode:SUCCESS_CODE, ErrMsg:nil}
}
