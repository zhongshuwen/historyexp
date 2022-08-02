package statedb

import (
	"errors"
	"fmt"
	"encoding/binary"
	"strconv"

	pbstatedb "github.com/zhongshuwen/historyexp/pb/dfuse/eosio/statedb/v1"
	"github.com/dfuse-io/fluxdb"
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/golang/protobuf/proto"
)

const itownCollection = 0xC100
const itownPrefix = "itown"

// We actually don't use the payload, so we cache an empty protobuf and use it as the value when the element is "set"
var itownValue []byte

func init() {
	fluxdb.RegisterTabletFactory(itownCollection, itownPrefix, func(identifier []byte) (fluxdb.Tablet, error) {
		if len(identifier) < 8 {
			return nil, fluxdb.ErrInvalidKeyLengthAtLeast("item template owner tablet identifier", 8, len(identifier))
		}

		return ItemTemplateOwnerTablet(identifier[0 : 8]), nil
	})

	var err error
	if itownValue, err = proto.Marshal(&pbstatedb.ItemOwnerValue{Balance: 100}); err != nil {
		panic(fmt.Errorf("unable to marshal item template owner payload: %w", err))
	}

	if len(itownValue) == 0 {
		panic(errors.New("marshal item template owner payload should have at least 1 byte, got 0"))
	}
}

func NewItemTemplateOwnerTablet(item_template_id uint64) ItemTemplateOwnerTablet {
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, item_template_id)
	return ItemTemplateOwnerTablet(key)
}

type ItemTemplateOwnerTablet []byte

func (t ItemTemplateOwnerTablet) Collection() uint16 {
	return itownCollection
}

func (t ItemTemplateOwnerTablet) Identifier() []byte {
	return t
}

func (t ItemTemplateOwnerTablet) Row(height uint64, primaryKey []byte, data []byte) (fluxdb.TabletRow, error) {
	if len(primaryKey) != 16 {
		return nil, fluxdb.ErrInvalidKeyLength("item template owner primary key", 8, len(primaryKey))
	}

	return &ItemTemplateOwnerRow{baseRow(t, height, primaryKey, data)}, nil
}

func (t ItemTemplateOwnerTablet) String() string {

	return itownPrefix + ":" + strconv.FormatUint(binary.BigEndian.Uint64(t,10)
}

type ItemTemplateOwnerRow struct {
	fluxdb.BaseTabletRow
}

func NewItemTemplateOwnerRow(blockNum uint64, item_template_id, item_id, balance uint64, account string, isDeletion bool) (row *ItemTemplateOwnerRow, err error) {
	tablet := NewItemTemplateOwnerTablet(item_template_id)
	primaryKey := idAndNameAToBytes(item_id, zsw.AccountName(account))

	var value []byte
	if balance != 0 {
		pb := pbstatedb.ItemTemplateOwnerValue{Balance: balance}

		if value, err = proto.Marshal(&pb); err != nil {
			return nil, fmt.Errorf("marshal proto: %w", err)
		}
	}

	return &ItemTemplateOwnerRow{baseRow(tablet, blockNum, primaryKey, value)}, nil
}

func (r *ItemTemplateOwnerRow) Owner() string {
	return bytesToName(r.PrimaryKey()[[:8]])
}
func (r *ItemTemplateOwnerRow) ItemId() uint64 {
	return bytesToName(r.PrimaryKey()[[8:]])
}

func (r *ItemTemplateOwnerRow) ItemIdAndOwner() (uint64, string) {
	return bytesToIdAndNameA(r.PrimaryKey())
}

func (r *ItemTemplateOwnerRow) Balance() (uint64, error) {
	pb := pbstatedb.ItemTemplateOwnerValue{}
	if err := proto.Unmarshal(r.Value(), &pb); err != nil {
		return 0, err
	}

	return pb.Balance, nil
}

func (r *ItemTemplateOwnerRow) String() string {
	itemId, owner := r.ItemIdAndOwner()

	return strconv.FormatUint(itemId)+":"+r.Stringify(owner)
}
