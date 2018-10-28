package common

// server
const (
	CONFIG_SERVER_MAIN   = "Main"
	CONFIG_SERVER_JUNGLE = "Jungle"
	CONFIG_SERVER_KYLIN  = "Kylin"
	CONFIG_SERVER_DEV    = "dev"

	CONFIG_DEFAULT_SERVER_ADDR      = ":8888"
	CONFIG_DEFAULT_READ_TIMEOUT     = 60
	CONFIG_DEFAULT_PID_FILE         = "/../data/eospark.pid"
	CONFIG_DEFAULT_VERSION_FILE     = "/../data/eospark.version"
	CONFIG_DEFAULT_LOG_CONFIG_FILE  = "/../config/log.xml"
	CONFIG_DEFAULT_WSS_CERT_CONFIG_FILE  = "/../config/cert.pem"
	CONFIG_DEFAULT_WSS_KEY_CONFIG_FILE  = "/../config/key.pem"
	CONFIG_DEFAULT_DATA_PATH        = "/../data/"
	CONFIG_DEFAULT_DATA_NAME        = "eospark"
	CONFIG_DEFAULT_DATA_CHANNEL_MAX = 5000

	INTERFACE_GET_BASE_INFO                = "get_base_info"
	INTERFACE_GET_OVERVIEW_STABLE          = "get_overview_low_refresh"
	INTERFACE_GET_OVERVIEW_REFRESH         = "get_overview_high_refresh"
	INTERFACE_GET_ACCOUNT_BLOCK_INFO       = "get_account_block_info"
	INTERFACE_GET_BLOCK_DETAIL_INFO        = "get_block_detail_info"
	INTERFACE_GET_BLOCK_TRANSACTION_INFO   = "get_block_transaction_info"
	INTERFACE_GET_TRANSACTION_DETAIL_INFO  = "get_transaction_detail_info"
	INTERFACE_GET_TRANSACTION_ACTION_INFO  = "get_transaction_action_info"
	INTERFACE_GET_ADDRESS_DETAIL_INFO      = "get_address_detail_info"
	INTERFACE_GET_BP_INFO                  = "get_bp_info"
	INTERFACE_GET_BLOCK_INFO               = "get_block_info"
	INTERFACE_GET_COIN_INFO                = "get_coin_info"
	INTERFACE_GET_TOKEN_PRICE_LIST         = "get_token_price_list"
	INTERFACE_GET_CONTRACT_INFO            = "get_contract_info"
	INTERFACE_GET_CONTRACT_LIST            = "get_contract_list"
	INTERFACE_GET_ACCOUNT_RELATED_TRX_INFO = "get_account_related_trx_info"
	INTERFACE_GET_ACCOUNT_INFO             = "get_account_info"
	INTERFACE_GET_SUB_ACCOUNT_INFO         = "get_sub_account_info"
	INTERFACE_GET_CONTRACT_TRX_INFO        = "get_contract_trx_info"
	INTERFACE_GET_REFLECT_EOS_ACCOUNT      = "get_reflect_eos_account"
	INTERFACE_GET_ACCOUNT_RESOURCE_INFO    = "get_account_resource_info"

	INTERFACE_GET_TOP_BID_LIST         = "get_top_bid_list"
	INTERFACE_VERIFY_CODE              = "verify_code"
	INTERFACE_APPLY_FOR_AUDIT          = "apply_for_audit"
	INTERFACE_PUSH_AUDIT_RESULT        = "push_audit_result"
	INTERFACE_GET_TOKEN_LIST_INFO      = "get_token_list_info"
	INTERFACE_GET_TARGET_TOKEN_DETAILS = "get_target_token_details"
	INTERFACE_GET_TOKEN_LIST_DETAILS   = "get_token_list_details"
	INTERFACE_GET_TOKEN_MAGNATE_RANK   = "get_token_magnate_rank"

	INTERFACE_GET_RAM_INFO    = "get_ram_info"
	INTERFACE_GET_SEARCH_INFO = "get_search_info"

	SUB_INTERFACE_ACCOUNT_RELATED_TRX_INFO = "tab_name"
)

// mysql
const (
	// config
	EMPTY_CONTRACT_HASH                       = "0000000000000000000000000000000000000000000000000000000000000000"
	CONFIG_DEFAULT_MYSQL_MAX_OPEN_CONNECTIONS = 2000
	CONFIG_DEFAULT_MYSQL_MAX_IDLE_CONNECTIONS = 100

	DB_COMMAND_SELECT_LIMIT  = "limit"
	DB_COMMAND_SELECT_OFFSET = "offset"

	DB_TRX_TIMESTAMP  = "trx_timestamp"
	DB_BLOCK_NUM      = "block_num"
	DB_TRANSACTION_ID = "transaction_id"
	DB_ACCOUNT        = "account"

	// table
	DB_TABLE_T_BLOCK_INFO           = "t_block_info"
	DB_TABLE_T_TRANSACTION_INFO     = "t_transaction_info"
	DB_TABLE_T_TRANSACTION_STATUS   = "t_transaction_status"
	DB_TABLE_T_ACTION_INFO          = "t_action_info"
	DB_TABLE_T_INLINE_ACTION_INFO   = "t_inline_action_info"
	DB_TABLE_T_BP_INFO              = "t_bp_info"
	DB_TABLE_T_ACCOUNT_INFO         = "t_account_info"
	DB_TABLE_T_PERMISSION_INFO      = "t_permission_info"
	DB_TABLE_T_TRANSFER_TRX_INFO    = "t_transfer_trx_info"
	DB_TABLE_T_TOKEN_CREATE_INFO    = "t_token_create_info"
	DB_TABLE_T_CONTRACT_TRX_INFO    = "t_contract_trx_info"
	DB_TABLE_T_VOTE_TRX_INFO        = "t_vote_trx_info"
	DB_TABLE_T_BID_TRX_INFO         = "t_bid_trx_info"
	DB_TABLE_T_RAM_TRX_INFO         = "t_ram_info"
	DB_TABLE_T_COIN_INFO            = "t_coin_info"
	DB_TABLE_T_BASE_INFO            = "t_base_info"
	DB_TABLE_T_EOS_REFLECT_INFO     = "t_eos_reflect_info"
	DB_TABLE_T_CONTRACT_CODE_INFO   = "t_contract_code_info"
	DB_TABLE_T_CONTRACT_AUDIT_INFO  = "t_contract_audit_info"
	DB_TABLE_T_AUDIT_CHECK_TRANSFER = "t_audit_check_transfer_info"
	DB_TABLE_T_TOKEN_LIST_SNAPSHOT  = "t_token_list_info"

	// t_block_info
	DB_T_BLOCK_INFO_F_ID                 = "id"
	DB_T_BLOCK_INFO_F_TRANSACTION_NUM    = "transaction_num"
	DB_T_BLOCK_INFO_F_ACTION_NUM         = "action_num"
	DB_T_BLOCK_INFO_F_PRODUCER           = "producer"
	DB_T_BLOCK_INFO_F_BLOCK_TIMESTAMP    = "block_timestamp"
	DB_T_BLOCK_INFO_F_TRANSACTION_MROOT  = "transaction_mroot"
	DB_T_BLOCK_INFO_F_ACTION_MROOT       = "action_mroot"
	DB_T_BLOCK_INFO_F_PRODUCER_SIGNATURE = "producer_signature"
	DB_T_BLOCK_INFO_F_RAW_DATA           = "raw_data"
	DB_T_BLOCK_INFO_F_TIMESTAMP          = "timestamp"

	DB_T_BLOCK_INFO_F_NO_REVERSE = "no_reverse"

	// t_transaction_info
	DB_T_TRANSACTION_INFO_F_ID                  = "id"
	DB_T_TRANSACTION_INFO_F_SIGNATURES          = "signatures"
	DB_T_TRANSACTION_INFO_F_REF_BLOCK_NUM       = "ref_block_num"
	DB_T_TRANSACTION_INFO_F_MAX_NET_USAGE_WORDS = "max_net_usage_words"
	DB_T_TRANSACTION_INFO_F_MAX_KCPU_USAGE      = "max_kcpu_usage"
	DB_T_TRANSACTION_INFO_F_DELAY_SEC           = "delay_sec"
	DB_T_TRANSACTION_INFO_F_PACKED_TRX          = "packed_trx"
	DB_T_TRANSACTION_INFO_F_EXPIRATION          = "expiration"
	DB_T_TRANSACTION_INFO_F_RAW_DATA            = "raw_data"

	DB_T_TRANSACTION_STATUS_F_STATUS = "status"

	// t_action_info
	DB_T_ACTION_INFO_F_ACCOUNT       = "account"
	DB_T_ACTION_INFO_F_ACTION        = "action"
	DB_T_ACTION_INFO_F_ACTION_ID     = "action_id"
	DB_T_ACTION_INFO_F_STATUS        = "status"
	DB_T_ACTION_INFO_F_AUTHORIZATION = "authorization"
	DB_T_ACTION_INFO_F_DATA          = "data"
	DB_T_ACTION_INFO_F_DATA_MD5      = "data_md5"
	DB_T_ACTION_INFO_F_LAYER         = "layer"
	DB_T_ACTION_INFO_F_HEX_DATA      = "hex_data"
	DB_T_ACTION_INFO_F_HAS_CONTRACT  = "has_contract"

	EOS_DEFAULT_ACCOUNT_EOSIO = "eosio"
	ACTION_NAME_NEWACCOUNT    = "newaccount"
	ACTION_NAME_UPDATEAUTH    = "updateauth"
	ACTION_NAME_DELETEAUTH    = "deleteauth"
	ACTION_NAME_LINKAUTH    = "linkauth"
	ACTION_NAME_UNLINKAUTH   = "unlinkauth"
	ACTION_NAME_SETABI        = "setabi"
	ACTION_NAME_SETCODE       = "setcode"
	ACTION_NAME_TRANSFER      = "transfer"
	ACTION_NAME_VOTE          = "voteproducer"
	ACTION_NAME_BID           = "bidname"
	ACTION_NAME_CREATE        = "create"
	ACTION_NAME_ISSUE         = "issue"
	ACTION_NAME_BUYRAM        = "buyram"
	ACTION_NAME_SELLRAM       = "sellram"
	ACTION_NAME_NONE          = "none"

	// t_bp_info
	DB_T_BP_INFO_F_ACCOUNT               = "account"
	DB_T_BP_INFO_F_NAME                  = "name"
	DB_T_BP_INFO_F_WEBSITE               = "website"
	DB_T_BP_INFO_F_INTRODUCTION          = "introduction"
	DB_T_BP_INFO_F_SCORE                 = "score"
	DB_T_BP_INFO_F_TOTAL_VOTE_NUM        = "total_vote_num"
	DB_T_BP_INFO_F_TOTAL_BLOCK_NUM       = "total_block_num"
	DB_T_BP_INFO_F_LNG                   = "lng"
	DB_T_BP_INFO_F_LAT                   = "lat"
	DB_T_BP_INFO_F_BLOCK_TIMESTAMP       = "block_timestamp"
	DB_T_BP_INFO_F_FIRST_BLOCK_TIMESTAMP = "first_block_timestamp"
	DB_T_BP_INFO_F_ACTIVE_TIMESTAMP      = "active_timestamp"
	DB_T_BP_INFO_F_COUNTRY               = "guojia"
	DB_T_BP_INFO_F_HIS_BANK              = "his_rank"
	DB_T_BP_INFO_F_LOGO_URL              = "logo_url"
	DB_T_BP_INFO_F_TOTAL_VOTE            = "vote"

	// t_account_info
	DB_T_ACCOUNT_INFO_F_ACCOUNT          = "account"
	DB_T_ACCOUNT_INFO_F_CREATER_ACCOUNT  = "creater_account"
	DB_T_ACCOUNT_INFO_F_CREATE_TIMESTAMP = "create_timestamp"
	DB_T_ACCOUNT_INFO_F_TRANSACTION_ID = "transaction_id"

	// t_permission_info
	DB_T_PERMISSION_INFO_F_ACCOUNT          = "account"
	DB_T_PERMISSION_INFO_F_PERM_NAME        = "perm_name"
	DB_T_PERMISSION_INFO_F_PERM_PARENT_NAME = "perm_parent"
	DB_T_PERMISSION_INFO_F_THRESHOLD        = "threshold"
	DB_T_PERMISSION_INFO_F_ITEM             = "item"
	DB_T_PERMISSION_INFO_F_ITEM_TYPE        = "item_type"
	DB_T_PERMISSION_INFO_F_WEIGHT           = "weight"

	ITEM_TYPE_ADDRESS = 1
	ITEM_TYPE_ACCOUNT = 2

	// t_transfer_trx
	DB_T_TRANSFER_TRX_F_INDEX         = "index"
	DB_T_TRANSFER_TRX_F_ISSUE_ACCOUNT = "issue_account"
	DB_T_TRANSFER_TRX_F_FROM_ACCOUNT  = "from_account"
	DB_T_TRANSFER_TRX_F_TO_ACCOUNT    = "to_account"
	DB_T_TRANSFER_TRX_F_QUANTITY      = "quantity"
	DB_T_TRANSFER_TRX_F_SYMBOL        = "symbol"
	DB_T_TRANSFER_TRX_F_MEMO          = "memo"
	DB_T_TRANSFER_TRX_F_TRX_TYPE      = "transfer_type"
	DB_T_TRANSFER_TRX_F_TRX_STATUS    = "status"
	DB_T_WHERE_HIDE_SMALL_QUANTITY    = "hide_small_quantity" //检查条件,非字段
	DB_T_TRANSFER_TRX_F_DATA_MD5      = "data_md5"

	// t_token_create_trx
	DB_T_TOKEN_CREATE_F_EXEC_ACCOUNT  = "exec_account"
	DB_T_TOKEN_CREATE_F_ISSUE_ACCOUNT = "issue_account"
	DB_T_TOKEN_CREATE_F_SUPPLY        = "supply"

	// t_contract_trx
	DB_T_CONTRACT_TRX_F_ACCOUNT   = "account"
	DB_T_CONTRACT_TRX_F_VMTYPE    = "vmtype"
	DB_T_CONTRACT_TRX_F_VMVERSION = "vmversion"

	// t_vote_trx
	DB_T_VOTE_TRX_F_VOTER     = "voter"
	DB_T_VOTE_TRX_F_PRODUCER  = "producer"
	DB_T_VOTE_TRX_F_PROXY     = "proxy"
	DB_T_VOTE_TRX_F_VOTE_TYPE = "vote_type"

	// t_bid_trx
	DB_T_BID_TRX_F_BIDDER   = "bidder"
	DB_T_BID_TRX_F_NEW_NAME = "newname"
	DB_T_BID_TRX_F_QUANTITY = "quantity"
	DB_T_BID_TRX_F_SYMBOL   = "symbol"
	DB_T_BID_TRX_F_STATE    = "state"

	// t_ram_trx
	DB_T_RAM_TRX_F_ACCOUNT    = "account"
	DB_T_RAM_TRX_F_TO_ACCOUNT = "to_account"
	DB_T_RAM_TRX_F_RAM_PRICE  = "ram_price"
	DB_T_RAM_TRX_F_EOS_PRICE  = "eos_price"
	DB_T_RAM_TRX_F_QUANTITY   = "quantity"
	DB_T_RAM_TRX_F_TYPE       = "type"

	// t_coin_info
	DB_T_COIN_INFO_F_PRICE      = "price"
	DB_T_COIN_INFO_F_MARKET_CAP = "market_cap"
	DB_T_COIN_INFO_F_TIMESTAMP  = "timestamp"

	DB_T_COIN_INFO_F_TIMESTAMP_START = "start"
	DB_T_COIN_INFO_F_TIMESTAMP_END   = "end"

	// t_ram_info
	DB_T_RAM_INFO_F_PRICE = "price"
	//DB_T_RAM_INFO_F_MARKET_CAP = "market_cap"
	DB_T_RAM_INFO_F_TIMESTAMP = "timestamp"

	DB_T_RAM_INFO_F_TIMESTAMP_START = "start"
	DB_T_RAM_INFO_F_TIMESTAMP_END   = "end"
	DB_T_RAM_INFO_F_TIMESTAMP_COUNT = "count"

	//t_base_info
	DB_T_BASE_INFO_F_TRX_NUM     = "trx_num"
	DB_T_BASE_INFO_F_ACTION_NUM  = "action_num"
	DB_T_BASE_INFO_F_ACCOUNT_NUM = "account_num"

	// t_eos_reflect_info
	DB_T_EOS_REFLECT_F_ETH_ACCOUNT = "eth_account"
	DB_T_EOS_REFLECT_F_EOS_ACCOUNT = "eos_account"
	DB_T_EOS_REFLECT_F_BLANCE      = "balance"

	//t_contract_code_info
	DB_T_CONTRACT_CODE_HASH        = "code_hash"
	DB_T_CONTRACT_CODE_JSON        = "code_json"
	DB_T_CONTRACT_CODE_UPLOAD_TIME = "upload_timestamp"
	DB_T_CONTRACT_CODE_URLS        = "code_file_urls"
	DB_T_CONTRACT_VERSION          = "version"
	DB_T_CONTRACT_OS               = "os"
	DB_T_CONTRACT_COMPILER_OPTION  = "compiler_option"

	//t_contract_audit_info
	DB_T_AUDIT_TASK_ID         = "task_id"
	DB_T_AUDIT_TASK_STATE      = "task_state"
	DB_T_AUDIT_ACCOUNT_NAME    = "account_name"
	DB_T_AUDIT_DEVELOPER_EMAIL = "developer_email"
	DB_T_AUDIT_DEVELOPER_NAME  = "developer_name"
	DB_T_AUDIT_DEVELOPER_PHONE = "developer_phone"
	DB_T_AUDIT_AUDITOR_ID      = "auditor_id"

	SEARCH_IFNO = "content"
)

const (
	// coinmarketcap
	CONFIG_DEFAULT_COIN_MARKET_CAP_ADDR           = "https://api.coinmarketcap.com/v2/ticker/1765/"
	CONFIG_DEFAULT_COIN_GECKO_ADDR                = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=eos"
	CONFIG_DEFAULT_COIN_MARKET_CAP_MAX_CONNECTION = 3
	CONFIG_DEFAULT_COIN_MARKET_CAP_TIMEOUT        = 300
	CONFIG_DEFAULT_COIN_MARKET_CAP_TIMER_INTERVAL = 1000

	// eos
	CONFIG_DEFAULT_EOS_GET_BLOCK_INTERFACE            = "http://127.0.0.1:8888/v1/chain/get_block"
	CONFIG_DEFAULT_EOS_GET_MAKE_UP_BLOCK_INTERFACE    = "http://47.75.132.47:8888/v1/chain/get_block"
	CONFIG_DEFAULT_EOS_GET_INFO_INTERFACE             = "http://127.0.0.1:8888/v1/chain/get_info"
	CONFIG_DEFAULT_EOS_GET_BALANCE_INTERFACE          = "http://127.0.0.1:8888/v1/chain/get_currency_balance"
	CONFIG_DEFAULT_EOS_GET_CODE_INTERFACE             = "http://127.0.0.1:8888/v1/chain/get_code"
	CONFIG_DEFAULT_EOS_GET_ABI_INTERFACE              = "http://127.0.0.1:8888/v1/chain/get_abi"
	CONFIG_DEFAULT_EOS_GET_RAW_CODE_INTERFACE         = "http://127.0.0.1:8888/v1/chain/get_raw_code_and_abi"
	CONFIG_DEFAULT_EOS_GET_BP_INTERFACE               = "http://127.0.0.1:8888/v1/chain/get_producers"
	CONFIG_DEFAULT_EOS_GET_TABLE_INTERFACE            = "http://127.0.0.1:8888/v1/chain/get_table_rows"
	CONFIG_DEFAULT_EOS_GET_ACCOUNT_RESOURCE_INTERFACE = "http://127.0.0.1:8888/v1/chain/get_account"
	CONFIG_DEFAULT_EOS_GET_CURRENCY_STATS_INTERFACE   = "http://127.0.0.1:8888/v1/chain/get_currency_stats"
	CONFIG_DEFAULT_EOS_GET_ETH_MAIN_NET_URL           = "https://mainnet.infura.io/Zay1Pf3PXNEThFtpupz9"

	CONFIG_DEFAULT_EOS_ACCOUNT_LIST       = "tokendapppub"
	CONFIG_DEFAULT_IRREVERSIBLE_BLOCK_NUM = 10
	CONFIG_DEFAULT_EOS_MAX_CONNECTION     = 3
	CONFIG_DEFAULT_EOS_TIMEOUT            = 300
	CONFIG_DEFAULT_EOS_TIMER_INTERVAL     = 1000
	CONFIG_DEFAULT_EOS_FREQUENCY_LIMIT    = 30
	CONFIG_DEFAULT_MYSQL_CACHE_TIME       = 30
	CONFIG_DEFAULT_EOS_TOOL_PATH          = "/root/eos/build/programs/cleos/cleos"
	CONFIG_DEFAULT_BP_TIMER_INTERVAL      = 60000

	CONFIG_DEFAULT_EOS_MAIN_SYS = "EOS"

	CONFIG_SHORT_CACHE_TIME       = 30
	CONFIG_LONG_CACHE_TIME        = 3600
	CONFIG_EVERY_DAY_REFRESH_TIME = 3600 //即凌晨1点
	CONFIG_CACHE_SVR_LIST         = "47.75.1.86:11211,47.75.132.47:11211,47.52.106.54:11211"
	CONFIG_NODE_API_HOST_LIST     = "https://eu1.eosdac.io" //,http://api.hkeos.com,https://api.eosdetroit.io

	CONFIG_CONTRACT_SOURCE_PATH = "/var/log/eos/contract"
	CONFIG_ACCOUNT_BLACK_LIST   = "blocktwitter"
	VERIFY_CONTRACT_URL         = "http://172.31.96.228:10003/verify_contract" //微服务在park2上
	AUDIT_CONTRACT_URL          = "http://localhost:11000/audit/invite"
)

// business
const (
	REPORT_DATA_TYPE_REQUEST_RECORD  = 1
	REPORT_DATA_TYPE_MYSQL_OPERATION = 2
	REPORT_DATA_TYPE_MYSQL_ROLLBACK  = 3
	REPORT_DATA_TYPE_BLOCK_SYNC      = 4

	DEFAULT_MAX_LIMIT                         = 50
	DEFAULT_ROLLBACK_AGAIN_SLEEP_MILLI_SECOND = 100
)

var SystemContracts = []string{
	"eosio.token", "eosio.bios",
	"eosio.sudo", "eosio.system",
	"eosio.msig", "eosio",
	"eosio.unregd",
}
