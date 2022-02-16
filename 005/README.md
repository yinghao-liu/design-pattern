# 005

Singleton Pattern

>  所有单例的实现都包含以下两个相同的步骤：
> - 将默认构造函数设为私有， 防止其他对象使用单例类的 `new`运算符。
> - 新建一个静态构建方法作为构造函数。 该函数会 “偷偷” 调用私有构造函数来创建对象， 并将其保存在一个静态成员变量中。 此后所有对于该函数的调用都将返回这一缓存对象。
> 如果你的代码能够访问单例类， 那它就能调用单例类的静态方法。 无论何时调用该方法， 它总是会返回相同的对象。



单例模式有两种形式

## 懒汉式

类加载时没有生成单例，只有当第一次调用 getlnstance 方法时才去创建这个单例

## 饿汉式

类一旦加载就创建一个单例，保证在调用 getInstance 方法之前单例已经存在了。

![A](singleton.drawio.svg)

## reference

1. [单例模式](http://c.biancheng.net/view/1338.html)
2. [单例模式](https://refactoringguru.cn/design-patterns/singleton)
