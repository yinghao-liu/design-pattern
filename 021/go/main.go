package main

import "strategypattern/strategy"

var PaymentSingle strategy.Payment

func PaymentInit() {
	// 初始化
	PaymentSingle.Init()
	// 此处可以动态添加
	PaymentSingle.SetPayment(strategy.PayAlipay, new(strategy.AlipayPayment))
	PaymentSingle.SetPayment(strategy.PayWechat, new(strategy.WechatPayment))
}
func main() {
	PaymentInit()

	payType := strategy.PayAlipay

	// 使用
	pay, _ := PaymentSingle.GetPayment(payType)
	pay.Pay()
}
