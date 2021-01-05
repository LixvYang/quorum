package api

import qApi "github.com/huo-ju/quorum/internal/pkg/api"

type PagerOpt struct {
	StartTrxId string
	Reverse    bool
	Page       int
}

// /api/v1/node
type NodeInfoStruct struct {
	NodeId      string              `json:"node_id"`
	NodePubKey  string              `json:"node_publickey"`
	NodeStatus  string              `json:"node_status"`
	NodeType    string              `json:"node_type"`
	NodeVersion string              `json:"node_version"`
	Peers       map[string][]string `json:"peers"`
}

// /api/v1/network
type NetworkInfoGroupStruct struct {
	GroupId   string   `json:"GroupId"`
	GroupName string   `json:"GroupName"`
	Peers     []string `json:"Peers"`
}

type NetworkInfoNodeStruct struct {
	Addrs      []string `json:"addrs"`
	EthAddr    string   `json:"ethaddr"`
	NatEnabled bool     `json:"nat_enabled"`
	NatType    string   `json:"nat_type"`
	PeerId     string   `json:"peerid"`
}

type NetworkInfoStruct struct {
	Groups []NetworkInfoGroupStruct `json:"groups"`
	Node   NetworkInfoNodeStruct    `json:"node"`
}

// /api/v1/groups
type GroupInfoStruct struct {
	OwnerPubKey    string `json:"OwnerPubKey"`
	GroupId        string `json:"GroupId"`
	GroupName      string `json:"GroupName"`
	LastUpdate     int64  `json:"LastUpdate"`
	LatestBlockNum int64  `json:"LatestBlockNum"`
	LatestBlockId  string `json:"LatestBlockId"`
	GroupStatus    string `json:"GroupStatus"`
}

type GroupInfoListStruct struct {
	Groups []GroupInfoStruct `json:"groups"`
}

func (a GroupInfoListStruct) Len() int { return len(a.Groups) }
func (a GroupInfoListStruct) Less(i, j int) bool {
	return a.Groups[i].GroupName < a.Groups[j].GroupName
}
func (a GroupInfoListStruct) Swap(i, j int) { a.Groups[i], a.Groups[j] = a.Groups[j], a.Groups[i] }

// /api/v1/group/$group/content
type ContentInnerStruct map[string]interface{}

type ReplyToStruct struct {
	TrxId string `json:"trxid"`
}

type ContentInnerMsgStruct struct {
	Name    string        `json:"name"`
	Content string        `json:"content"`
	ReplyTo ReplyToStruct `json:"inreplyto"`
}

type WalletStruct struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ContentInnerProfileStruct struct {
	Name   string         `json:"name"`
	Wallet []WalletStruct `json:"wallet"`
}

type ContentStruct struct {
	TrxId     string             `json:"TrxId"`
	Publisher string             `json:"Publisher"`
	Content   ContentInnerStruct `json:"Content"`
	TypeUrl   string             `json:"TypeUrl"`
	TimeStamp int64              `json:"TimeStamp"`
}

type ContentList []ContentStruct

func (a ContentList) Len() int           { return len(a) }
func (a ContentList) Less(i, j int) bool { return a[i].TimeStamp < a[j].TimeStamp }
func (a ContentList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// /api/v1/trx/$id
type TrxStruct struct {
	TrxId     string `json:"TrxId"`
	GroupId   string `json:"GroupId"`
	Sender    string `json:"Sender"`
	Pubkey    string `json:"Pubkey"`
	Data      []byte `json:"Data"`
	TimeStamp int64  `json:"TimeStamp"`
	Version   string `json:"Version"`
	Expired   int64  `json:"Expired"`
	Signature string `json:"Signature"`
}

type TrxRespStruct struct {
	TrxId string `json:"trx_id"`
}

// /api/v1/group/profile
type QuorumPersonStruct struct {
	Name string `json:"name"`
}
type QuorumTargetStruct struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
type NickReqStruct struct {
	Person QuorumPersonStruct `json:"person"`
	Target QuorumTargetStruct `json:"target"`
	Type   string             `json:"type"`
}

type NickRespStruct TrxRespStruct

type TokenRespStruct struct {
	Token string `json:"token"`
}

// POST /api/v1/group/content
type ContentReqObjectStruct struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type ContentReqTargetStruct struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ContentReqStruct struct {
	Object ContentReqObjectStruct `json:"object"`
	Target ContentReqTargetStruct `json:"target"`
	Type   string                 `json:"type"`
}

type ContentRespStruct TrxRespStruct

// POST /api/v1/group/join
type JoinRespStruct struct {
	GroupId   string `json:"group_id"`
	Signature string `json:"signature"`
}

// POST /api/v1/group
type CreateGroupReqStruct struct {
	Name string `json:"group_name"`
}

type GenesisBlockStruct struct {
	BlockId        string `json:"BlockId"`
	GroupId        string `json:"GroupId"`
	BlockNum       int    `json:"BlockNum"`
	Timestamp      int64  `json:"Timestamp"`
	ProducerId     string `json:"ProducerId"`
	ProducerPubKey string `json:"ProducerPubKey"`
	Hash           string `json:"Hash"`
	Signature      string `json:"Signature"`
}

type GroupSeedStruct struct {
	GenesisBlock GenesisBlockStruct `json:"genesis_block"`
	GroupId      string             `json:"group_id"`
	GroupName    string             `json:"group_name"`
	OwnerPubKey  string             `json:"owner_pubkey"`
	Signature    string             `json:"signature"`
}

type LeaveGroupReqStruct struct {
	GroupId string `json:"group_id"`
}

type GroupLeaveRetStruct struct {
	GroupId   string `json:"group_id"`
	Signature string `json:"signature"`
}

type DeleteGroupReqStruct struct {
	GroupId string `json:"group_id"`
}

type GroupDelRetStruct struct {
	GroupId     string `json:"group_id"`
	OwnerPubKey string `json:"owner_pubkey"`
	Signature   string `json:"signature"`
}

type GroupForceSyncRetStruct struct {
	GroupId string `json:"GroupId"`
	Error   string `json:"Error"`
}

// Get /v1/block/:group_id/:block_num and /v1/block/:block_id
type BlockStruct struct {
	BlockId        string       `json:"BlockId,omitempty"`
	GroupId        string       `json:"GroupId,omitempty"`
	PrevBlockId    string       `json:"PrevBlockId,omitempty"`
	BlockNum       int64        `json:"BlockNum,omitempty"`
	Timestamp      int64        `json:"Timestamp,omitempty"`
	PreviousHash   []byte       `json:"PreviousHash,omitempty"`
	ProducerId     string       `json:"ProducerId,omitempty"`
	ProducerPubKey string       `json:"ProducerPubKey,omitempty"`
	Trxs           []*TrxStruct `json:"Trxs,omitempty"`
	Hash           []byte       `json:"Hash,omitempty"`
	Signature      []byte       `json:"Signature,omitempty"`
}

// GET /api/v1/network/peers/ping
type PingInfoItemStruct struct {
	Addrs       []string             `json:"addrs"`
	Protocols   []string             `json:"protocols"`
	RTT         string               `json:"rtt"`
	Connections []qApi.AddrProtoPair `json:"connections"`
}
