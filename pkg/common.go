package pkg

import (
	"github.com/ZzzWClock/pexpay-go-sdk/utils"
)

// QueryC2COrderHistoryListPage
// 按页获取历史订单
// @param form  *C2COrderListHistoryPageForm
// @param pexPay PexPayInterface
func QueryC2COrderHistoryListPage(form *C2COrderListHistoryPageForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.queryC2COrderHistoryListPage(param)
}

// QueryC2COrderListPage
// 按页获取订单列表
// @param form  *C2COrderListPageForm
// @param pexPay PexPayInterface
func QueryC2COrderListPage(form *C2COrderListPageForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.queryC2COrderListPage(param)
}

// QueryC2CPaymentMethodAll
// 获取所有有效的付款方式
// @param pexPay PexPayInterface
func QueryC2CPaymentMethodAll(pexPay PexPayInterface) ([]byte, error) {
	return pexPay.queryC2CPaymentMethodAll()
}

// QueryC2CAdClass
// 获取 C2C 广告类别
// @param pexPay PexPayInterface
func QueryC2CAdClass(pexPay PexPayInterface) ([]byte, error) {
	return pexPay.queryC2CAdClass()
}

// QueryC2CAdPrice
// 获取 C2C 广告报价
// @param form  *C2CAdPriceForm
// @param pexPay PexPayInterface
func QueryC2CAdPrice(form *C2CAdPriceForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.queryC2CAdPrice(param)
}

// QueryC2CAdDetailByAdNo
// 获取 C2C 广告详情
// @param form  *C2CAdDetailByAdNoForm
// @param pexPay PexPayInterface
func QueryC2CAdDetailByAdNo(form *C2CAdDetailByAdNoForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.queryC2CAdDetailByAdNo(param)
}

// QueryC2CUserAdListPage
// 获取 C2C 用户广告列表
// @param form  *C2CUserAdListPageForm
// @param pexPay PexPayInterface
func QueryC2CUserAdListPage(form *C2CUserAdListPageForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.queryC2CUserAdListPage(param)
}

// C2CAdPush
// 发布 C2C 广告
// @param form  *C2CAdPushForm
// @param pexPay PexPayInterface
func C2CAdPush(form *C2CAdPushForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2CAdPush(param)
}

// C2CAdUpdate
// 更新 C2C 广告
// @param form  *C2CAdUpdateForm
// @param pexPay PexPayInterface
func C2CAdUpdate(form *C2CAdUpdateForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2CAdUpdate(param)
}

// C2CAdUpdateStatus
// 更新 C2C 广告状态
// @param form  *C2CAdUpdateStatusForm
// @param pexPay PexPayInterface
func C2CAdUpdateStatus(form *C2CAdUpdateStatusForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2CAdUpdateStatus(param)
}

// C2CAdSearch
// 搜索 C2C 广告
// @param form  *C2CAdSearchForm
// @param pexPay PexPayInterface
func C2CAdSearch(form *C2CAdSearchForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2CAdSearch(param)
}

// C2COrderAdd
// C2C 下单
// @param form  *C2COrderAddForm
// @param pexPay PexPayInterface
func C2COrderAdd(form *C2COrderAddForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2cOrderAdd(param)
}

// C2COrderMark
// C2C 标记订单已支付
// @param form  *C2COrderPassForm
// @param pexPay PexPayInterface
func C2COrderMark(form *C2COrderMarkForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2cOrderMark(param)
}

// C2COrderPass
// C2C 放币
// @param form  *C2COrderPassForm
// @param pexPay PexPayInterface
func C2COrderPass(form *C2COrderPassForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2cOrderPass(param)
}

// C2COrderCancel
// C2C 取消订单
// @param form  *C2COrderCancelForm
// @param pexPay PexPayInterface
func C2COrderCancel(form *C2COrderCancelForm, pexPay PexPayInterface) ([]byte, error) {

	param, err := utils.RefStructChangeMap(form)
	if err != nil {
		return nil, err
	}
	return pexPay.c2cOrderCancel(param)
}
