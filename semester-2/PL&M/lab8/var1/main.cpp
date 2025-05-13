#include <iostream>
#include "Polygon.cpp"

int main() {
    std::pair<double, double> square[4] = {
        {0.0, 0.0},
        {1.0, 0.0},
        {1.0, 1.0},
        {0.0, 1.0}
    };
    Polygon<4, double> poly(square);

    std::cout << "Perimeter: " << poly.perimeter() << std::endl;

    double pi = 2 * acos(0.0); 
    poly.rotate(pi / 2);

    for (int i = 0; i < 4; ++i) {
        const auto& vertex = poly.getVertex(i);
        std::cout << "Vertex " << i << ": (" << vertex.first << ", " << vertex.second << ")" << std::endl;
    }
    return 0;
}