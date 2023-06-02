package pkg

import (
	"bytes"
	"github.com/ZzzWClock/pexpay-go-sdk/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type PexPayStruct struct {
	ApiKey    string
	ApiSecret string
	address   string
}

func (p *PexPayStruct) queryC2COrderHistoryListPage(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/listUserOrderHistory")
	return p.doGet(param)
}

func (p *PexPayStruct) queryC2COrderListPage(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/listOrders")
	return p.doPost(param)
}

func (p *PexPayStruct) queryC2CPaymentMethodAll() ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/paymentMethod/listAll")
	return p.doPost(nil)
}

func (p *PexPayStruct) queryC2CAdClass() ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/getAvailableAdClassify")
	return p.doGet(nil)
}

func (p *PexPayStruct) queryC2CAdPrice(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/getReferencePrice")
	return p.doPost(param)
}

func (p *PexPayStruct) queryC2CAdDetailByAdNo(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/getDetailByNo")
	return p.doPost(param)
}

func (p *PexPayStruct) queryC2CUserAdListPage(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/listWithPagination")
	return p.doPost(param)
}

func (p *PexPayStruct) c2CAdPush(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/post")
	return p.doPost(param)
}

func (p *PexPayStruct) c2CAdUpdate(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/update")
	return p.doPost(param)
}

func (p *PexPayStruct) c2CAdUpdateStatus(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/updateStatus")
	return p.doPost(param)
}

func (p *PexPayStruct) c2CAdSearch(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/ad/search")
	return p.doPost(param)
}

func (p *PexPayStruct) c2cOrderAdd(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/placeOrder")
	return p.doPost(param)
}

func (p *PexPayStruct) c2cOrderMark(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/markOrderAsPaid")
	return p.doPost(param)
}

func (p *PexPayStruct) c2cOrderPass(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/releaseCoin")
	return p.doPost(param)
}

func (p *PexPayStruct) c2cOrderCancel(param map[string]string) ([]byte, error) {

	p.setAddress("https://api.pexpay.com/sapi/v1/c2c/orderMatch/cancelOrder")
	return p.doPost(param)
}

func (p *PexPayStruct) setAddress(address string) {
	//p.urlAddress = "https://api.pexpay.com"
	p.address = address
}

func (p *PexPayStruct) doGet(param map[string]string) ([]byte, error) {

	if param != nil {
		values := utils.GenerateUrlValuesByMapString(param)
		values.Set("timestamp", strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))
		signs := utils.GenerateHmac(values.Encode(), p.ApiSecret)
		values.Set("signature", signs)
		p.address += "?" + values.Encode()
	} else {
		values := make(url.Values)
		values.Set("timestamp", strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))
		signs := utils.GenerateHmac(values.Encode(), p.ApiSecret)
		values.Set("signature", signs)
		p.address += "?" + values.Encode()
	}

	req, err := http.NewRequest("GET", p.address, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-MBX-APIKEY", p.ApiKey)

	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *PexPayStruct) doPost(param map[string]string) ([]byte, error) {

	var body io.Reader
	if param != nil {

		values := utils.GenerateUrlValuesByMapString(param)
		values.Set("timestamp", strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))
		signs := utils.GenerateHmac(values.Encode(), p.ApiSecret)
		values.Set("signature", signs)
		log.Println("values.Encode() : ", values.Encode())
		body = strings.NewReader(values.Encode())
	} else {
		values := make(url.Values)
		values.Set("timestamp", strconv.Itoa(int(time.Now().UnixNano()/int64(time.Millisecond))))
		signs := utils.GenerateHmac(values.Encode(), p.ApiSecret)
		values.Set("signature", signs)
		body = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest("POST", p.address, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-MBX-APIKEY", p.ApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))
	_, err = io.Copy(buf, res.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
