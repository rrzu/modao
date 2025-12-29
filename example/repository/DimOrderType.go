package repository

import (
	"context"

	"github.com/rrzu/modao"
)

var dimOrderTypeDaoSingle modao.SingleDao[*DimOrderTypeDao]

type DimOrderTypeDao struct {
	*modao.HologresBaseDao
}

func InstanceDimOrderTypeDao(ctx context.Context) *DimOrderTypeDao {
	return dimOrderTypeDaoSingle.Do(ctx, func(withDebug bool) *DimOrderTypeDao {
		return &DimOrderTypeDao{modao.NewHologresBaseDao(&DimOrderTypeMod{}, withDebug)}
	})
}

// DimOrderTypeMod 订单类型维表
type DimOrderTypeMod struct {
	OrderId             int64  `json:"order_id" gorm:"column:order_id"`                             // 订单id
	IsSupportSign       int32  `json:"is_support_sign" gorm:"column:is_support_sign"`               // 锁机订单 1是0否
	SelfOrderSign       int32  `json:"self_order_sign" gorm:"column:self_order_sign"`               // 自营订单 1是0否
	SelfServerSign      int32  `json:"self_server_sign" gorm:"column:self_server_sign"`             // 自营店铺 1是0否
	ZftOrderSign        int32  `json:"zft_order_sign" gorm:"column:zft_order_sign"`                 // 直付通订单 1是0否
	DepositFreeSign     int32  `json:"deposit_free_sign" gorm:"column:deposit_free_sign"`           // 免押协助订单 1是0否
	DisOrderSign        int32  `json:"dis_order_sign" gorm:"column:dis_order_sign"`                 // 派发订单 1是0否
	BuyoutOrderSign     int32  `json:"buyout_order_sign" gorm:"column:buyout_order_sign"`           // 买断订单 1是0否
	DirectOrderSign     int32  `json:"direct_order_sign" gorm:"column:direct_order_sign"`           // 盲发订单 1是0否
	InstShippingSign    int32  `json:"inst_shipping_sign" gorm:"column:inst_shipping_sign"`         // 代发仓发货订单 1是0否
	SelfShippingSign    int32  `json:"self_shipping_sign" gorm:"column:self_shipping_sign"`         // 自发仓发货订单 1是0否
	DomainPoolOrderSign int32  `json:"domain_pool_order_sign" gorm:"column:domain_pool_order_sign"` // 转化订单 1是0否
	DeliveryRemindSign  int32  `json:"delivery_remind_sign" gorm:"column:delivery_remind_sign"`     // 加急发货 1是0否
	EtlTime             string `json:"etl_time" gorm:"column:etl_time"`                             // etl时间
	ThirdPartyStoreSign int32  `json:"third_party_store_sign" gorm:"column:third_party_store_sign"` // 品牌门店订单 1是0否
}

func (t *DimOrderTypeMod) TableName() string {
	return string(t.Table().TableName)
}

func (t *DimOrderTypeMod) Table() *modao.HologresTbl {
	return &modao.HologresTbl{
		ConnectInformation: modao.ConnectInfo{
			ConnectName: "hologres-sz-dev",
		},
		DatabaseName: "data_analysis_prod",
		PatternName:  "public",
		TableName:    "dim_order_type",
	}
}
