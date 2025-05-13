#include <cmath>
#include <iostream>

using namespace std;

template<typename T>
class Curve {
private:
    T sinRatio;
    T cosRatio;
public:
    Curve(bool flag) {
        if (flag) {
            sinRatio = 1;
            cosRatio = 0;
        }else{
            cosRatio = 1;
            sinRatio = 0;
        }
    }

    Curve(T sinK, T cosK) : sinRatio(sinK), cosRatio(cosK) {}

    Curve operator+(Curve& other){
        auto new_sinRatio = this->sinRatio + other.sinRatio;
        auto new_cosRatio = this->cosRatio + other.cosRatio;
        return Curve(new_sinRatio, new_cosRatio);
    }

    Curve operator-(Curve& other){
        auto new_sinRatio = this->sinRatio - other.sinRatio;
        auto new_cosRatio = this->cosRatio - other.cosRatio;
        return Curve(new_sinRatio, new_cosRatio);
    }

    Curve operator*(T k){
        auto new_sinRatio = k*this->sinRatio;
        auto new_cosRatio = k*this->cosRatio;
        return Curve(new_sinRatio, new_cosRatio);
    }

    Curve operator-(){
        auto new_sinRatio = -this->sinRatio;
        auto new_cosRatio = -this->cosRatio;
        return Curve(new_sinRatio, new_cosRatio);
    }

    T operator()(T x){
        return sinRatio * sin(x) + cosRatio * cos(x);
    }

    Curve operator!(){
        auto new_sinRatio = -this->cosRatio;
        auto new_cosRatio = this->sinRatio;
        return Curve(new_sinRatio, new_cosRatio);
    }

    void show() {
        cout << sinRatio << "six(x) + " << cosRatio << "cos(x)" << endl;
    }
};

int main(){
    Curve<float> curve1(true);
    Curve<float> curve2(false);

    auto curve3 = (curve1 * 2) + curve2;

    curve3.show();
    curve3 = !curve3;
    curve3.show();

    cout << curve3(1.5);
}