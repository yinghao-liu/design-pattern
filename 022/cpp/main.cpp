#include <iostream>
#include <memory>

// 父类 定义模板
class Base {
public:
  void Template() {
    Step1();
    Step2();
    Step3();
  }

private:
  virtual void Step1() { std::cout << "default step1" << std::endl; }
  virtual void Step2() { std::cout << "default step2" << std::endl; }
  virtual void Step3() { std::cout << "default step3" << std::endl; }
};

// 具体子类
class Concrete : public Base {
private:
  virtual void Step1() { std::cout << "Concrete step1" << std::endl; }
};

int main() {
  // std::unique_ptr<Base> a = std::make_unique<Concrete>();
  auto a = std::make_unique<Concrete>();
  a->Template();

  std::cout << "---------------------------------" << std::endl;
  std::unique_ptr<Base> b = std::make_unique<Concrete>();
  b->Template();
}