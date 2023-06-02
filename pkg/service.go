package pkg

type C2COrderListPageForm struct {
	AdNo            string `json:"adNo"`            // 广告编号
	Asset           string `json:"asset"`           // 数字货币，如"BTC"
	OrderStatus     string `json:"orderStatus"`     // 订单状态
	PayType         string `json:"payType"`         // 支付方式类型，如 "BANK", "WECHAT"
	OrderStatusList string `json:"orderStatusList"` // 订单状态
	StartDate       string `json:"startDate"`       // 订单创建开始时间，以毫秒为单位
	EndDate         string `json:"endDate"`         // 订单创建结束时间，以毫秒为单位
	Page            string `json:"page"`            // 页码，最小值=1
	Rows            string `json:"rows"`            // 每页的条数，最小值=1，最大值=2000
	Order           string `json:"order"`           // 按哪个参数排序
	Sort            string `json:"sort"`            // 升序降序，可选值："Asc", "Desc"
}

type C2COrderListHistoryPageForm struct {
	TradeType      string `json:"tradeType"`      // 交易类型，可选值："BUY", "SELL"
	StartTimestamp string `json:"startTimestamp"` // 订单创建开始时间，以毫秒为单位
	EndTimestamp   string `json:"endTimestamp"`   // 订单创建结束时间，以毫秒为单位
	Page           string `json:"page"`           // 页码，最小值=1
	Rows           string `json:"rows"`           // 每页的条数，最小值=1，最大值=2000
	Order          string `json:"order"`          // 按哪个参数排序
	Sort           string `json:"sort"`           // 升序降序，可选值："Asc", "Desc"
}

type C2CAdPriceForm struct {
	Assets       string `json:"assets"`       // 数字货币
	FiatCurrency string `json:"fiatCurrency"` // 法币
	PayType      string `json:"payType"`      // 支付方式类型，如 "BANK", "WECHAT"
	TradeType    string `json:"tradeType"`    // 交易类型，可选值："BUY", "SELL"
}

type C2CAdDetailByAdNoForm struct {
	AdNo string `json:"adNo"` // 广告编号
}

type C2CUserAdListPageForm struct {
	AdNo      string `json:"adNo"`      // 广告编号
	AdStatus  string `json:"adStatus"`  // 广告状态, 可选值："1", "2", "3"
	Asset     string `json:"asset"`     // 数字货币，如"BTC"
	TradeType string `json:"tradeType"` // 交易类型，可选值："BUY", "SELL"
	StartDate string `json:"startDate"` // 广告创建开始时间
	EndDate   string `json:"endDate"`   // 广告创建结束时间
	FiatUnit  string `json:"fiatUnit"`  // 法币
	Page      string `json:"page"`      // 页码
	Rows      string `json:"rows"`      // 每页条目数
}

type C2CAdPushForm struct {
	Asset                 string              `json:"asset"`                 // 数字货币，如"BTC"
	Classify              string              `json:"classify"`              // 可选值："mass", "profession", mass 普通广告，profession 专业广告
	FiatUnit              string              `json:"fiatUnit"`              // 法币
	InitAmount            string              `json:"initAmount"`            // 初始金额，单位：数字货币
	MaxSingleTransAmount  string              `json:"maxSingleTransAmount"`  // 最大单笔法币金额
	MinSingleTransAmount  string              `json:"minSingleTransAmount"`  // 最小单笔法币金额
	OnlineNow             string              `json:"onlineNow"`             // 是否立即上线
	PayTimeLimit          string              `json:"payTimeLimit"`          // 支付时间限制，单位：分钟
	Price                 string              `json:"price"`                 // 出售价格，单位：法币
	PriceType             string              `json:"priceType"`             // 可选值："1", "2", 1 固定价格 2 浮动价格
	TradeType             string              `json:"tradeType"`             // 交易类型，可选值："BUY", "SELL"
	PriceFloatingRatio    string              `json:"priceFloatingRatio"`    // 浮动价格比例，仅在浮动价格类型广告生效
	TradeMethods          []map[string]string `json:"tradeMethods"`          // 支付方式信息
	BuyerRegDaysLimit     string              `json:"buyerRegDaysLimit"`     // taker账号已注册的最少天数，不传表示无限制
	BuyerBtcPositionLimit string              `json:"buyerBtcPositionLimit"` // taker持有数字货币的最低总价值（以BTC计算）
	Remarks               string              `json:"remarks"`               // 广告备注
	AutoReplyMsg          string              `json:"autoReplyMsg"`          // 自动回复内容，下单后该内容会自动通过chat发送给taker
}

type C2CAdUpdateForm struct {
	AdNo                  string              `json:"adNo"`                  // 广告编号
	AdStatus              string              `json:"adStatus"`              // 广告状态, 可选值："1", "2", "3", "4"
	InitAmount            string              `json:"initAmount"`            // 初始金额，单位：数字货币
	MaxSingleTransAmount  string              `json:"maxSingleTransAmount"`  // 最大单笔法币金额
	MinSingleTransAmount  string              `json:"minSingleTransAmount"`  // 最小单笔法币金额
	PayTimeLimit          string              `json:"payTimeLimit"`          // 支付时间限制，单位：分钟
	Price                 string              `json:"price"`                 // 出售价格，单位：法币
	PriceFloatingRatio    string              `json:"priceFloatingRatio"`    // 浮动价格比例
	TradeMethods          []map[string]string `json:"tradeMethods"`          // 支付方式信息
	BuyerRegDaysLimit     string              `json:"buyerRegDaysLimit"`     // taker账号已注册的最少天数，不传表示无限制
	BuyerBtcPositionLimit string              `json:"buyerBtcPositionLimit"` // taker持有数字货币的最低总价值（以BTC计算）
	Remarks               string              `json:"remarks"`               // 广告备注
	AutoReplyMsg          string              `json:"autoReplyMsg"`          // 自动回复内容，下单后该内容会自动通过chat发送给taker
}

type C2CAdUpdateStatusForm struct {
	AdNo     string `json:"adNo"`     // 广告编号
	AdStatus string `json:"adStatus"` // 广告状态, 可选值："1", "2", "3", "4"
}

type C2CAdSearchForm struct {
	Asset         string `json:"asset"`         // 数字货币，如 "BTC"
	FiatUnit      string `json:"fiatUnit"`      // 法币
	TradeType     string `json:"tradeType"`     // 交易类型，可选值："BUY", "SELL"
	PayTypes      string `json:"payTypes"`      // 支持的支付方式类型, 如 "BANK", "WECHAT"
	PublisherType string `json:"publisherType"` // 交易类型，可选值："merchant", "user"，不限制请传空
	TransAmount   string `json:"transAmount"`   // 期望交易的法币金额
	Page          string `json:"page"`          // 页码，最小值=1
	Rows          string `json:"rows"`          // 每页的条数，最小值=1，最大值=2000
}

type C2COrderAddForm struct {
	Asset       string `json:"asset"`       // 数字货币，如"BTC"
	FiatUnit    string `json:"fiatUnit"`    // 法币
	TradeType   string `json:"tradeType"`   // 交易类型，可选值："BUY", "SELL"
	BuyType     string `json:"buyType"`     // 按法币/数字货币数量下单，可选值："BY_MONEY", "BY_AMOUNT"
	TotalAmount string `json:"totalAmount"` // 下单数量
	PayId       string `json:"payId"`       // 下卖单时必传，传用于收款的支付方式的id
	PayType     string `json:"payType"`     // 下买单时必传，传下单广告支持的支付方式的标识之一，如 "WECHAT"
	AdNo        string `json:"adNo"`        // 下单目标广告的广告编号

}

type C2COrderMarkForm struct {
	OrderNo string `json:"orderNo"` // 订单编号
	PayId   string `json:"payId"`   // 买单必传，支付的目标支付方式的id

}

type C2COrderPassForm struct {
	OrderNo string `json:"orderNo"` // 订单编号

}

type C2COrderCancelForm struct {
	OrderNo                   string `json:"orderNo"`                   // 订单编号
	OrderCancelReasonCode     string `json:"orderCancelReasonCode"`     // 取消订单原因码，可选值："1", "2", "3", "4", "5"
	OrderCancelAdditionalInfo string `json:"orderCancelAdditionalInfo"` // 取消原因选择 "5" 时，用此字段提交额外的取消原因信息

}
