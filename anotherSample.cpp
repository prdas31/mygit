#include <iostream>

using namespace std;

int main() {
  cout << "hello world!!\n";

  auto myAdder = [](auto x, auto y) { return x + y; };
  cout << "Adder By Pranabesh: " << myAdder(2, 5) << endl;
  return 0;
}
