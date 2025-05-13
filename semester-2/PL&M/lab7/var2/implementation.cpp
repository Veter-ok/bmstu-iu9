#include "declaration.h"
#include <cmath>
#include <algorithm>

Polyline::Polyline() : points(nullptr), segmentCount(0), capacity(0) {}

Polyline::Polyline(const Polyline& other) : segmentCount(other.segmentCount), capacity(other.capacity) {
    points = new Point[capacity];
    for (int i = 0; i < segmentCount; ++i) {
        points[i] = other.points[i];
    }
}

Polyline::~Polyline() {
    delete[] points;
}

Polyline& Polyline::operator=(const Polyline& other) {
    if (this != &other) {
        delete[] points;
        segmentCount = other.segmentCount;
        capacity = other.capacity;
        points = new Point[capacity];
        for (int i = 0; i < segmentCount; ++i) {
            points[i] = other.points[i];
        }
    }
    return *this;
}

int Polyline::getSegmentCount() const {
    return segmentCount;
}

Point& Polyline::getPoint(int index) {
    return points[index];
}

const Point& Polyline::getPoint(int index) const {
    return points[index];
}

void Polyline::addPoint(double x, double y) {
    if (segmentCount >= capacity) {
        int newCapacity = (capacity == 0) ? 2 : capacity * 2;
        Point* newPoints = new Point[newCapacity];
        for (int i = 0; i < segmentCount; ++i) {
            newPoints[i] = points[i];
        }
        delete[] points;
        points = newPoints;
        capacity = newCapacity;
    }
    points[segmentCount] = {x, y};
    segmentCount++;
}

void Polyline::removeShortSegments(double minLength) {
    if (segmentCount < 2) return;

    int newCount = 0;
    for (int i = 0; i < segmentCount; ++i) {
        int next = (i + 1) % segmentCount;
        double dx = points[next].x - points[i].x;
        double dy = points[next].y - points[i].y;
        double length = sqrt(dx*dx + dy*dy);
        
        if (length >= minLength) {
            points[newCount++] = points[i];
        }
    }

    segmentCount = newCount;
}