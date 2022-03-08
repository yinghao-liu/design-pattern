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

/****************************支付的接口**********************************/
type Paymenter interface {
	PayType() string
	Pay() error
	CheckBalance() (float32, error)
}

/****************************支付的管理context**********************************/
// 扩展了UML图中context的功能，可添加Pay()方法，实现具体Pay操作前后的通用操作
// TODO 单例模式实现 -> 不用单例也行,可维护多个context,但每个都需要独立初始化
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

/****************************支付的具体实现**********************************/
// 支付宝支付 - 可单独拆分文件实现
type AlipayPayment struct {
}

func (p *AlipayPayment) PayType() string {
	return PayAlipay
}

func (p *AlipayPayment) Pay() error {
	fmt.Printf("AlipayPayment doing pay\n")
	return nil
}
func (p *AlipayPayment) CheckBalance() (float32, error) {
	fmt.Printf("AlipayPayment doing CheckBalance\n")
	return float32(0), nil
}

// 微信支付 - 可单独拆分文件实现
type WechatPayment struct {
}

func (p *WechatPayment) PayType() string {
	return PayWechat
}

func (p *WechatPayment) Pay() error {
	fmt.Printf("WechatPayment doing pay\n")
	return nil
}
func (p *WechatPayment) CheckBalance() (float32, error) {
	fmt.Printf("WechatPayment doing CheckBalance\n")
	return float32(0), nil
}

// 银行卡支付 - 可单独拆分文件实现
type BankCardPayment struct {
}

func (p *BankCardPayment) PayType() string {
	return PayBankCard
}

func (p *BankCardPayment) Pay() error {
	fmt.Printf("BankCardPayment doing pay\n")
	return nil
}
func (p *BankCardPayment) CheckBalance() (float32, error) {
	return float32(0), errors.New("not support CheckBalance")
}
