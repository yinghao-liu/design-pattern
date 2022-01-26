package strategy

import (
	"errors"
	"fmt"
)

// 支付类型
const (
	PayAlipay   = "Alipay"
	PayWechat   = "Wechat"
	PayBankCard = "BankCard"
)

// 支付接口
type Paymenter interface {
	PayType() string
	Pay() error
}

// TODO 单例模式实现
type Payment struct {
	payment map[string]Paymenter
}

func (p *Payment) Init() {
	if nil == p.payment {
		p.payment = make(map[string]Paymenter)
	}
}
func (p *Payment) SetPayment(payType string, pay Paymenter) {
	p.payment[payType] = pay
}
func (p *Payment) GetPayment(payType string) (Paymenter, error) {
	if pay, exist := p.payment[payType]; exist {
		return pay, nil
	}
	return nil, errors.New("not exist")
}

// 支付宝支付
type AlipayPayment struct {
}

func (p *AlipayPayment) PayType() string {
	return PayAlipay
}

func (p *AlipayPayment) Pay() error {
	fmt.Printf("AlipayPayment doing pay\n")
	return nil
}

// 微信支付
type WechatPayment struct {
}

func (p *WechatPayment) PayType() string {
	return PayWechat
}

func (p *WechatPayment) Pay() error {
	fmt.Printf("WechatPayment doing pay\n")
	return nil
}

// 银行卡支付
type BankCardPayment struct {
}

func (p *BankCardPayment) PayType() string {
	return PayBankCard
}

func (p *BankCardPayment) Pay() error {
	fmt.Printf("BankCardPayment doing pay\n")
	return nil
}
