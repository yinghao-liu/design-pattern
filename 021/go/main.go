package main

import (
	"fmt"
	"strategypattern/strategy"
)

var PaymentSingle strategy.Payment

func PaymentInit() {
	// 初始化
	PaymentSingle.Init()
	// 此处可以动态添加
	// TODO 使用工厂模式创建类

	PaymentSingle.SetPayment(strategy.PayAlipay, new(strategy.AlipayPayment))
	PaymentSingle.SetPayment(strategy.PayWechat, new(strategy.WechatPayment))
	PaymentSingle.SetPayment(strategy.PayBankCard, new(strategy.BankCardPayment))
}
func main() {
	PaymentInit()

	payType := strategy.PayAlipay

	// 使用支付
	pay, _ := PaymentSingle.GetPayment(payType)
	pay.Pay()

	payType = strategy.PayBankCard
	// 使用查询余额
	pay, _ = PaymentSingle.GetPayment(payType)
	if _, err := pay.CheckBalance(); nil != err {
		fmt.Printf("CheckBalance error:%s\n", err.Error())
	}
}
