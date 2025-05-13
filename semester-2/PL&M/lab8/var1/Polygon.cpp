#include <cmath>

using namespace std;

template<int N, typename T>
class Polygon {
private:
    pair<T, T> data[N];
public:
    Polygon(pair<T, T> (&vertices)[N]) {
        for (int i = 0; i < N; ++i) {
            data[i] = vertices[i];
        }
    }

    double perimeter() {
        int ans = 0;
        for (int i = 0; i < N; i++) {
            double x = data[i].first - data[(i+1) % N].first;
            double y = data[i].second - data[(i+1) % N].second;
            ans += sqrt(x*x + y*y);
        }
        return ans;
    }

    void addNode(int i, pair<T, T> node) {
        pair<T, T> new_data[N + 1];
        for (int j = 0; j < N; j++) {
            if (j < i) {
                new_data[j] = data[j];
            }else if (j == i) {
                new_data[j] = node;
            }else {
                new_data[j] = data[j-1];
            }
        }
        data = new_data;
    }

    template <typename U = T>
    typename enable_if<is_same<U, double>::value>::type
    rotate(double angle) {
        const double x0 = data[0].first;
        const double y0 = data[0].second;
        const double cos_a = cos(angle);
        const double sin_a = sin(angle);

        for (size_t i = 1; i < N; ++i) {
            const double dx = data[i].first - x0;
            const double dy = data[i].second - y0;
            data[i].first  = x0 + dx * cos_a - dy * sin_a;
            data[i].second = y0 + dx * sin_a + dy * cos_a;
        }
    }

    const pair<T, T>& getVertex(int i) const {
        return data[i];
    }
};