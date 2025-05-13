#include <cmath>
#include <iostream>

using namespace std;

template<typename T>
class ApproxNum {
private:
    T a;
    T k;
public:
    ApproxNum(T a, T k) : a(a), k(k) {}

    ApproxNum operator*=(T k) {
        this->a = k * this->a;
        this->k = k * this->k;
        return *this;
    }

    ApproxNum operator+=(ApproxNum& other) {
        this->a = this->a + other.a;
        this->k = this->k + other.k;
        return *this;
    }

    ApproxNum operator-=(ApproxNum& other) {
        this->a = this->a - other.a;
        this->k = this->k - other.k;
        return *this;
    }

    ApproxNum operator+(ApproxNum& other) {
        auto new_a = this->a + other.a;
        auto new_k = this->k + other.k;
        return ApproxNum(new_a, new_k);
    }

    ApproxNum operator-(ApproxNum& other) {
        auto new_a = this->a - other.a;
        auto new_k = this->k - other.k;
        return ApproxNum(new_a, new_k);
    }

    ApproxNum operator*(T k) {
        auto new_a = k * this->a;
        auto new_k = k * this->k;
        return ApproxNum(new_a, new_k);
    }

    bool operator==(ApproxNum& other) {
        if (this->a == other.a && this->k == other.k){
            return true;
        }
        return false;
    }

    bool operator!=(ApproxNum& other) {
        return !(this == other);
    }

    bool operator<(ApproxNum& other) {
        if (this->a < other.a || (this->a == other.a && this->k < other.k)){
            return true;
        }
        return false;
    }

    bool operator<=(ApproxNum& other) {
        return this < other || this == other;
    }

    bool operator>(ApproxNum& other) {
        return !(this < other) && !(this == other);
    }

    bool operator>=(ApproxNum& other) {
        return this > other || this == other;
    }

    void show(){ cout << this->a << "+" << this->k << "δ" << endl; }
};

int main(){
    ApproxNum<int> n1(10, 5);
    ApproxNum<int> n2(6, 2);

    n1 *= 2;
    n1.show();
    auto n3 = n1 + n2;
    n3.show();
}