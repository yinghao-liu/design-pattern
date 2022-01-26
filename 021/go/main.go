package main

import "fmt"

// 支付类型
const (
	PayAlipay   = "Alipay"
	PayWechat   = "Wechat"
	PayBankCard = "BankCard"
)

// 支付接口
type Paymenter interface {
	payType() string
	pay() error
}

// TODO: 工厂模式实现
var Payment map[string]Paymenter

// 支付宝支付
type AlipayPayment struct {
}

func (p *AlipayPayment) payType() string {
	return PayAlipay
}

func (p *AlipayPayment) pay() error {
	fmt.Printf("AlipayPayment doing pay\n")
	return nil
}

// 微信支付
type WechatPayment struct {
}

func (p *WechatPayment) payType() string {
	return PayWechat
}

func (p *WechatPayment) pay() error {
	fmt.Printf("WechatPayment doing pay\n")
	return nil
}

// 银行卡支付
type BankCardPayment struct {
}

func (p *BankCardPayment) payType() string {
	return PayBankCard
}

func (p *BankCardPayment) pay() error {
	fmt.Printf("BankCardPayment doing pay\n")
	return nil
}

func main() {
	// 初始化
	// TODO 初始化封装Payment提供接口实现不直接访问
	Payment = make(map[string]Paymenter)
	Payment[PayAlipay] = new(AlipayPayment)
	Payment[PayWechat] = new(WechatPayment)

	payType := PayAlipay
	// 使用
	// TODO   提供接口
	pay := Payment[payType]
	pay.pay()

}
