package tarantool_test

import (
	"encoding/json"
	"fmt"
	"time"
)

//UnixTime is
type UnixTime uint64

// MarshalJSON implements json.Marshaler.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON implements json.Marshaler.
func (t *UnixTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != "" {
		ti, err := time.ParseInLocation("2006-01-02 15:04:05", x, time.Local)
		if err != nil {
			return err
		}
		*t = UnixTime(ti.Unix())
	} else {
		*t = UnixTime(time.Now().Unix())
	}
	return nil
}

//Product is
type Test struct {
	_msgpack struct{} `msgpack:",asArray,omitempty"`
	ID       uint64   `json:"id"`
}

//Product is
type Product struct {
	_msgpack       struct{} `msgpack:",asArray,omitempty"`
	Productid      uint64   `json:"productid"`
	Productname    string   `json:"productname"`
	Productkey     string   `json:"productkey"`
	Producttype    uint64   `json:"producttype"`
	Productaddtime UnixTime `json:"productaddtime"`
	Productmodel   string   `json:"productmodel"`
	Productsecret  string   `json:"productsecret"`
	Productdesc    string   `json:"productdesc"`
}

//Node is
type Node struct {
	_msgpack      struct{} `msgpack:",asArray,omitempty"`
	Nodeid        uint64   `json:"nodeid"`
	Nodepid       uint64   `json:"nodepid"`
	Nodename      string   `json:"nodename"`
	Nodelabel     string   `json:"nodelabel"`
	Nodetype      uint64   `json:"nodetype"`
	Nodeip        string   `json:"nodeip"`
	Nodeport      uint64   `json:"nodeport"`
	Nodewsport    uint64   `json:"nodewsport"`
	Nodestatus    uint64   `json:"nodestatus"`
	Nodedriver    string   `json:"nodedriver"`
	Nodeaddr      string   `json:"nodeaddr"`
	Nodecontacts  string   `json:"nodecontacts"`
	Nodephone     string   `json:"nodephone"`
	Nodedesc      string   `json:"nodedesc"`
	Nodemodel     string   `json:"nodemodel"`
	Nodeconfig    string   `json:"nodeconfig"`
	Nodepnames    string   `json:"-"`
	ChildrenNodes []Node   `json:"children"`
	Items         []Item   `json:"items"`
}

//Item is
type Item struct {
	_msgpack        struct{} `msgpack:",asArray,omitempty"`
	Itemid          uint64   `json:"itemid"`
	Nodeid          uint64   `json:"nodeid"`
	Itemname        string   `json:"itemname"`
	Itemdatatype    string   `json:"itemdatatype"`
	Itemaddr        string   `json:"itemaddr"`
	Itemlength      uint64   `json:"itemlength"`
	Itemroundrule   uint64   `json:"itemroundrule"`
	Itemdigitnum    uint64   `json:"itemdigitnum"`
	Itemdesc        string   `json:"itemdesc"`
	Itemrw          uint64   `json:"itemrw"`
	Itemscanrate    uint64   `json:"itemscanrate"`
	Itemscantype    string   `json:"itemscantype"`
	Itemstatus      uint64   `json:"itemstatus"`
	Itemallownull   uint64   `json:"itemallownull"`
	Itemallowupdate uint64   `json:"itemallowupdate"`
}

//NodeStatus is
type NodeStatus struct {
	_msgpack      struct{} `msgpack:",asArray,omitempty"`
	Nodename      string   `json:"nodename"`
	Nodestatus    uint64   `json:"nodestatus"`
	Nodechecktime UnixTime `json:"nodechecktime"`
	Nodefeedback  string   `json:"nodefeedback"`
}

//LiveData is
type LiveData struct {
	_msgpack struct{}    `msgpack:",asArray"`
	Tagname  string      `json:"tagname"`
	Value    interface{} `json:"value"`
	DateTime UnixTime    `json:"datetime"`
	Source   string      `json:"source"`
	Nodeid   uint64      `json:"nodeid"`
	Nodename string      `json:"nodename"`
	Tagaddr  string      `json:"tagaddr"`
	Quality  uint64      `json:"quality"`
}
