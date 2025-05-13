#include <iostream>
#include <vector>

using namespace std;

template<int L, int H, size_t N>
class IntVector {
private:
    using StorageType = typename std::conditional<(H - L <= 256), char, int>::type;
    std::vector<StorageType> data;

public:
    IntVector() : data(N, static_cast<StorageType>(0 - L)) {}

    IntVector(initializer_list<int> values) : data(N) {
        size_t i = 0;
        for (int val : values) {
            data[i++] = static_cast<StorageType>(val - L);
        }
    }

    int operator[](size_t index) const {
        return static_cast<int>(data[index]) + L;
    }

    void set(size_t index, int value) {
        data[index] = static_cast<StorageType>(value - L);
    }

    template<int L2, int H2>
    auto operator+(const IntVector<L2, H2, N>& other) const {
        IntVector<L + L2, H + H2, N> result;
        for (size_t i = 0; i < N; ++i) {
            result.set(i, (*this)[i] + other[i]);
        }
        return result;
    }

    template<int L2, int H2>
    int dot(const IntVector<L2, H2, N>& other) const {
        int result = 0;
        for (size_t i = 0; i < N; ++i) {
            result += (*this)[i] * other[i];
        }
        return result;
    }

    void print() const {
        cout << "[ ";
        for (size_t i = 0; i < N; ++i) {
            cout << (*this)[i] << " ";
        }
        cout << "]" << std::endl;
    }

    constexpr size_t size() const { return N; }
};