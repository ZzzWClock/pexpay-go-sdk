package pkg

type PexPayInterface interface {
	queryC2COrderHistoryListPage(map[string]string) ([]byte, error) // queryC2COrderHistoryListPage
	queryC2COrderListPage(map[string]string) ([]byte, error)        // queryC2COrderListPage
	queryC2CPaymentMethodAll() ([]byte, error)                      // queryC2CPaymentMethodAll
	queryC2CAdClass() ([]byte, error)                               // queryC2CAdClass
	queryC2CAdPrice(map[string]string) ([]byte, error)              // queryC2CAdPrice
	queryC2CAdDetailByAdNo(map[string]string) ([]byte, error)       // queryC2CAdDetailByAdNo
	queryC2CUserAdListPage(map[string]string) ([]byte, error)       // queryC2CUserAdListPage
	c2CAdPush(map[string]string) ([]byte, error)                    // c2CAdPush
	c2CAdUpdate(map[string]string) ([]byte, error)                  // c2CAdUpdate
	c2CAdUpdateStatus(map[string]string) ([]byte, error)            // c2CAdUpdateStatus
	c2CAdSearch(map[string]string) ([]byte, error)                  // c2CAdSearch
	c2cOrderAdd(map[string]string) ([]byte, error)                  // c2cOrderAdd
	c2cOrderMark(map[string]string) ([]byte, error)                 // c2cOrderMark
	c2cOrderPass(map[string]string) ([]byte, error)                 // c2cOrderPass
	c2cOrderCancel(map[string]string) ([]byte, error)               // c2cOrderCancel
}
