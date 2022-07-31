package statedb

import (
	"errors"
	"fmt"
	"encoding/binary"

	pbstatedb "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/statedb/v1"
	"github.com/dfuse-io/fluxdb"
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/golang/protobuf/proto"
)

const iownCollection = 0xC000
const iownPrefix = "iown"

// We actually don't use the payload, so we cache an empty protobuf and use it as the value when the element is "set"
var iownValue []byte

func init() {
	fluxdb.RegisterTabletFactory(iownCollection, iownPrefix, func(identifier []byte) (fluxdb.Tablet, error) {
		if len(identifier) < 8 {
			return nil, fluxdb.ErrInvalidKeyLengthAtLeast("item owner tablet identifier", 8, len(identifier))
		}

		return ItemOwnerTablet(identifier[0 : 8]), nil
	})

	var err error
	if iownValue, err = proto.Marshal(&pbstatedb.ItemOwnerValue{Balance: 100}); err != nil {
		panic(fmt.Errorf("unable to marshal item owner payload: %w", err))
	}

	if len(iownValue) == 0 {
		panic(errors.New("marshal item owner payload should have at least 1 byte, got 0"))
	}
}

func NewItemOwnerTablet(item_id uint64) ItemOwnerTablet {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, item_id)
	return ItemOwnerTablet(key)
}

type ItemOwnerTablet []byte

func (t ItemOwnerTablet) Collection() uint16 {
	return iownCollection
}

func (t ItemOwnerTablet) Identifier() []byte {
	return t
}

func (t ItemOwnerTablet) Row(height uint64, primaryKey []byte, data []byte) (fluxdb.TabletRow, error) {
	if len(primaryKey) != 8 {
		return nil, fluxdb.ErrInvalidKeyLength("item owner primary key", 8, len(primaryKey))
	}

	return &ItemOwnerRow{baseRow(t, height, primaryKey, data)}, nil
}

func (t ItemOwnerTablet) String() string {

	return iownPrefix + ":" + (binary.BigEndian.Uint64(t))
}

type ItemOwnerRow struct {
	fluxdb.BaseTabletRow
}

func NewItemOwnerRow(blockNum uint64, item_id, balance uint64, account string, isDeletion bool) (row *ItemOwnerRow, err error) {
	tablet := NewItemOwnerTablet(item_id)
	primaryKey := nameaToBytes(zsw.AccountName(account))

	var value []byte
	if balance != 0 {
		pb := pbstatedb.ItemOwnerValue{Balance: balance}

		if value, err = proto.Marshal(&pb); err != nil {
			return nil, fmt.Errorf("marshal proto: %w", err)
		}
	}

	return &ItemOwnerRow{baseRow(tablet, blockNum, primaryKey, value)}, nil
}

func (r *ItemOwnerRow) Owner() string {
	return bytesToName(r.PrimaryKey())
}

func (r *ItemOwnerRow) Balance() (uint64, error) {
	pb := pbstatedb.ItemOwnerValue{}
	if err := proto.Unmarshal(r.Value(), &pb); err != nil {
		return 0, err
	}

	return pb.Balance, nil
}

func (r *ItemOwnerRow) String() string {
	return r.Stringify(bytesToName(r.PrimaryKey()))
}
