#include <iostream>
#include "IntVector.cpp"

int main() {
    IntVector<-10, 10, 3> v1 = {1, 2, 3};
    IntVector<0, 1000, 3> v2 = {100, 200, 300};

    auto v3 = v1 + v2;
    std::cout << "Sum: ";
    v3.print();
    
    std::cout << "product: " << v1.dot(v2) << std::endl; 
    return 0;
}