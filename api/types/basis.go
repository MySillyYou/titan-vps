package types

import (
	"time"

	"github.com/LMF709268224/titan-vps/lib/trxbridge/core"
)

type Hellos struct {
	Msg string
}

// OrderState represents the state of an order in the process of being pulled.
type OrderState int64

// Constants defining various states of the order process.
const (
	// Created order
	Created OrderState = iota
	// WaitingPayment Waiting for user to payment order
	WaitingPayment
	// BuyGoods buy goods
	BuyGoods
	// Done the order done
	Done
)

// Int returns the int representation of the order state.
func (s OrderState) Int() int64 {
	return int64(s)
}

// User user info
type User struct {
	UUID      string    `db:"uuid" json:"uuid"`
	UserName  string    `db:"user_name" json:"user_name"`
	PassHash  string    `db:"pass_hash" json:"pass_hash"`
	Address   string    `db:"address" json:"address"`
	Public    string    `db:"public" json:"public"`
	Token     string    `db:"token" json:"token"`
	Role      int32     `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type DescribePriceReq struct {
	RegionId                string
	InstanceType            string
	PriceUnit               string
	ImageID                 string
	InternetChargeType      string
	SystemDiskCategory      string
	SystemDiskSize          int32
	Period                  int32
	Amount                  int32
	InternetMaxBandwidthOut int32
	DataDisk                []*DescribePriceRequestDataDisk
}

type DescribePriceRequestDataDisk struct {
	Category         *string
	PerformanceLevel *string
	Size             *int64
}

type DescribePriceResponse struct {
	Currency      string
	OriginalPrice float32
	TradePrice    float32
	USDPrice      float32
}

type DescribeImageResponse struct {
	ImageId      string
	ImageName    string
	ImageFamily  string
	Platform     string
	OSType       string
	OSName       string
	Architecture string
}
type CreateOrderReq struct {
	CreateInstanceReq
	Amount int32
}

type CreateInstanceReq struct {
	Id                      string  `db:"id"`
	RegionId                string  `db:"region_id"`
	InstanceId              string  `db:"instance_id"`
	UserID                  string  `db:"user_id"`
	OrderID                 string  `db:"order_id"`
	InstanceType            string  `db:"instance_type"`
	DryRun                  bool    `db:"dry_run"`
	ImageId                 string  `db:"image_id"`
	SecurityGroupId         string  `db:"security_group_id"`
	InstanceChargeType      string  `db:"instance_charge_type"`
	PeriodUnit              string  `db:"period_unit"`
	Period                  int32   `db:"period"`
	InternetMaxBandwidthOut int32   `db:"bandwidth_out"`
	InternetMaxBandwidthIn  int32   `db:"bandwidth_in"`
	IpAddress               string  `db:"ip_address"`
	TradePrice              float32 `db:"trade_price"`
	SystemDiskCategory      string  `db:"system_disk_category"`
	InternetChargeType      string  `db:"internet_charge_type"`
	SystemDiskSize          int32   `db:"system_disk_size"`
	DataDisk                []*DescribePriceRequestDataDisk
}

type CreateInstanceResponse struct {
	InstanceID       string  `db:"instance_id"`
	OrderId          string  `db:"order_id"`
	RequestId        string  `db:"request_id"`
	TradePrice       float32 `db:"trade_price"`
	PublicIpAddress  string  `db:"public_ip_address"`
	PrivateKey       string
	PrivateKeyStatus int `db:"private_key_status"`
}
type DescribeInstanceTypeReq struct {
	RegionId         string
	MemorySize       float32
	CpuArchitecture  string
	InstanceCategory string
	CpuCoreCount     int32
	MaxResults       int64
	NextToken        string
}

type DescribeRecommendInstanceTypeReq struct {
	RegionId           string
	Memory             float32
	Cores              int32
	InstanceChargeType string
}

type DescribeRecommendInstanceResponse struct {
	Memory             int32
	Cores              int32
	InstanceType       string
	InstanceTypeFamily string
}

type DescribeInstanceTypeResponse struct {
	InstanceTypes []*DescribeInstanceType
	NextToken     string
}

type DescribeInstanceType struct {
	InstanceTypeId         string
	MemorySize             float32
	CpuArchitecture        string
	InstanceCategory       string
	CpuCoreCount           int32
	InstanceTypeFamily     string
	PhysicalProcessorModel string
	NextToken              string
}

type CreateKeyPairResponse struct {
	KeyPairID      string
	KeyPairName    string
	PrivateKeyBody string
}

type AttachKeyPairResponse struct {
	Code       string
	InstanceId string
	Message    string
	Success    string
}

// OrderRecord represents information about an order record
type OrderRecord struct {
	OrderID       string     `db:"order_id"`
	From          string     `db:"from_addr"`
	UserID        string     `db:"user_id"`
	To            string     `db:"to_addr"`
	Value         string     `db:"value"`
	State         OrderState `db:"state"`
	TradePrice    float32    `db:"trade_price"`
	DoneState     int64      `db:"done_state"`
	CreatedHeight int64      `db:"created_height"`
	CreatedTime   time.Time  `db:"created_time"`
	DoneTime      time.Time  `db:"done_time"`
	DoneHeight    int64      `db:"done_height"`
	VpsID         int64      `db:"vps_id"`
	Msg           string     `db:"msg"`
	TxHash        string     `db:"tx_hash"`
}

type OrderRecordResponse struct {
	Total int
	List  []*OrderRecord
}

// RechargeState Recharge order state
type RechargeState int64

// Constants defining various states of the recharge process.
const (
	// RechargeCreate Recharge create
	RechargeCreate RechargeState = iota
	// RechargeDone Recharge done
	RechargeDone
	// RechargeRefund Recharge Refund
	RechargeRefund
)

// WithdrawState Withdraw order state
type WithdrawState int64

// Constants defining various states of the recharge process.
const (
	// WithdrawCreate Withdraw create
	WithdrawCreate WithdrawState = iota
	// WithdrawDone Withdraw done
	WithdrawDone
)

type MyInstanceKeyState int64
type MyInstanceState int64

const (
	KeyNoSet MyInstanceKeyState = iota
	KeyHaveSet
)

const (
	InstanceCreate MyInstanceState = iota
	InstanceRunning
	InstanceStop
)

type LoginType int64

// Constants defining various states of the recharge process.
const (
	// LoginTypeMetaMask
	LoginTypeMetaMask LoginType = iota
	// LoginTypeTron
	LoginTypeTron
)

type RechargeResponse struct {
	Total int
	List  []*RechargeRecord
}

// RechargeRecord represents information about an recharge record
type RechargeRecord struct {
	OrderID       string        `db:"order_id"`
	From          string        `db:"from_addr"`
	UserID        string        `db:"user_id"`
	To            string        `db:"to_addr"`
	Value         string        `db:"value"`
	State         RechargeState `db:"state"`
	CreatedHeight int64         `db:"created_height"`
	CreatedTime   time.Time     `db:"created_time"`
	DoneTime      time.Time     `db:"done_time"`
	DoneHeight    int64         `db:"done_height"`
}

type WithdrawResponse struct {
	Total int
	List  []*WithdrawRecord
}

// WithdrawRecord represents information about an withdraw record
type WithdrawRecord struct {
	OrderID       string        `db:"order_id"`
	From          string        `db:"from_addr"`
	User          string        `db:"user_id"`
	To            string        `db:"to_addr"`
	Value         string        `db:"value"`
	State         WithdrawState `db:"state"`
	CreatedHeight int64         `db:"created_height"`
	CreatedTime   time.Time     `db:"created_time"`
	DoneTime      time.Time     `db:"done_time"`
	DoneHeight    int64         `db:"done_height"`
	WithdrawAddr  string        `db:"withdraw_addr"`
	WithdrawHash  string        `db:"withdraw_hash"`
	Executor      string        `db:"executor"`
}

type PaymentCompletedReq struct {
	OrderID       string
	TransactionID string
}

type UserReq struct {
	UserId    string
	Signature string
	Type      LoginType
}

type UserResponse struct {
	UserId   string
	SignCode string
	Token    string
}

type UserInfoTmp struct {
	UserLogin UserResponse
	OrderInfo OrderRecord
}

type Token struct {
	TokenString string
	UserId      string
	Expiration  time.Time
}

// EventTopics represents topics for pub/sub events
type EventTopics string

const (
	// EventFvmTransferWatch node online event
	EventFvmTransferWatch EventTopics = "fvm_transfer_watch"
	// EventTronTransferWatch node online event
	EventTronTransferWatch EventTopics = "tron_transfer_watch"
)

func (t EventTopics) String() string {
	return string(t)
}

type FvmTransferWatch struct {
	TxHash string
	From   string
	To     string
	Value  string
}

type TronTransferWatch struct {
	TxHash string
	From   string
	To     string
	Value  string
	State  core.Transaction_ResultContractResult
	Height int64
	UserID string
}

type ExchangeRateRsp struct {
	Code int32               `json:"code"`
	Data ExchangeRateDataRsp `json:"data"`
}
type ExchangeRateDataRsp struct {
	Rate float32 `json:"rate"`
}

type RechargeAddress struct {
	Addr   string `db:"addr"`
	UserID string `db:"user_id"`
}

type MyInstanceResponse struct {
	Total int
	List  []*MyInstance
}

type MyInstance struct {
	ID                 string             `db:"id"`
	InstanceId         string             `db:"instance_id"`
	OrderID            string             `db:"order_id"`
	UserID             string             `db:"user_id"`
	PrivateKeyStatus   MyInstanceKeyState `db:"private_key_status"`
	InstanceName       string             `db:"instance_name"`
	InstanceSystem     string             `db:"instance_system"`
	Location           string             `db:"location"`
	Price              float32            `db:"price"`
	State              string             `db:"state"`
	InternetChargeType string             `db:"internet_charge_type"`
	CreatedTime        time.Time          `db:"created_time"`
}

type InstanceDetails struct {
	ID         string `db:"id"`
	InstanceId string `db:"instance_id"`
}
