# 001

工厂模式

以制作椅子为例，以client为我方视角，说明工厂模式

## Stone Age

没有工厂的时代，制作什么只能自己动手，需要掌握如何制造椅子（例如下例的`new`）和材料。

```go
// concrete product
type ChairBlue struct {
}

func (c ChairBlue) sit() {
	fmt.Printf("sit in ChairBlue\n")
}

// client
func main() {
	a := new(ChairBlue)
	a.sit()
}
```

## Simple Factory

有了简单的工厂，我知道能坐（有sit方法）的就是椅子

```go
// abstract product
type Chair interface {
	sit()
}
```

有了椅子制造工厂，有两种类型可供你选择，只需要告诉他要红椅子还是绿椅子

```go
// concrete factory
func ChairFactory(style string) Chair {
	switch style {
	case "blue":
		return new(ChairBlue)
	case "red":
		return new(ChairRed)
	}

	return nil
}

// client
func main() {

	a := ChairFactory("blue")
	a.sit()
}
```

类图

![simpleFactory](simpleFactory.drawio.png)

另外，简单工厂还不在设计模式之列

简单工厂出现了经典的、不满足开闭原则的switch-case逻辑，而这种情况正是[策略模式](../021)要解决的

## Factory Method

将简单工厂的工厂使用策略模式实现，就变成了工厂方法。



闭包和工厂模式

包的可见性和工厂模式

工厂方法(Factory Method)





![factoryMethod](factoryMethod.drawio.png)

工厂方法为每个 Product 的子类提供了一个生产类，在生产 Product 子类比较复杂的情况下使用，创建和使用隔离。如果产品只有一个，就变成了简单工厂模式。

到此为止，还是只能生产一类产品，如果想同时生产椅子和床，这个方式就满足不了了。



## reference

- [工厂方法](http://c.biancheng.net/view/1348.html)
- [工厂方法](https://refactoringguru.cn/design-patterns/factory-method)

