#ifndef DECLARATION_H
#define DECLARATION_H

struct Point {
    double x;
    double y;
};

class Polyline {
private:
    Point* points;
    int segmentCount;
    int capacity;
public:
    Polyline();
    Polyline(const Polyline& other);
    virtual ~Polyline();
    Polyline& operator=(const Polyline& other);
    int getSegmentCount() const;
    Point& getPoint(int index);
    const Point& getPoint(int index) const;
    void addPoint(double x, double y);
    void removeShortSegments(double minLength);
};

#endif