#include "declaration.h"
#include <iostream>

void printPolyline(Polyline& pl) {
    std::cout << "Ломаная линия содержит " << pl.getSegmentCount() << " точек:" << std::endl;
    for (int i = 0; i < pl.getSegmentCount(); ++i) {
        Point& p = pl.getPoint(i);
        std::cout << "Точка " << i << ": (" << p.x << ", " << p.y << ")" << std::endl;
    }
    std::cout << std::endl;
}

int main() {
    Polyline pl;
    pl.addPoint(0.0, 0.0);
    pl.addPoint(1.0, 1.0);
    pl.addPoint(5.0, 0.0);
    pl.addPoint(1.0, -1.0);

    std::cout << "Было:" << std::endl;
    printPolyline(pl);

    double minLength = 1.5;
    pl.removeShortSegments(minLength);
    std::cout << "После удаления" << minLength << ":" << std::endl;
    printPolyline(pl);

    std::cout << "точка изменена:" << std::endl;
    pl.getPoint(0).x = 10.0;
    pl.getPoint(0).y = 10.0;
    printPolyline(pl);

    return 0;
}