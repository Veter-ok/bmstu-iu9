#include "declaration.h"
#include <algorithm>

int GCD(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

void Rational::normalize() {
    if (denominator < 0) {
        numerator *= -1;
        denominator *= -1;
    }
    int gcd = GCD(abs(numerator), abs(denominator));
    if (gcd != 0) {
        numerator /= gcd;
        denominator /= gcd;
    }
}

Rational::Rational(int num, int denom) : numerator(num), denominator(denom) {
    if (denominator == 0) {
        denominator = 1;
    }
    normalize();
}

Rational::Rational(const Rational& other) 
    : numerator(other.numerator), denominator(other.denominator) {}

int Rational::getNumerator() { return numerator; }
int Rational::getDenominator() { return denominator; }

Rational Rational::operator+(Rational other) {
    int newNum = numerator * other.denominator + other.numerator * denominator;
    int newDenom = denominator * other.denominator;
    return Rational(newNum, newDenom);
}

Rational Rational::operator*(Rational other) {
    int newNum = numerator * other.numerator;
    int newDenom = denominator * other.denominator;
    return Rational(newNum, newDenom);
}

Rational& Rational::operator=(const Rational& other) {
    if (this != &other) {
        numerator = other.numerator;
        denominator = other.denominator;
    }
    return *this;
}

RationalMatrix::RationalMatrix(int m, int n) : rows(m), cols(n) {
    data = new Rational*[rows];
    for (int i = 0; i < rows; ++i) {
        data[i] = new Rational[cols];
    }
}

RationalMatrix::RationalMatrix(const RationalMatrix& other) : rows(other.rows), cols(other.cols) {
    data = new Rational*[rows];
    for (int i = 0; i < rows; ++i) {
        data[i] = new Rational[cols];
        for (int j = 0; j < cols; ++j) {
            data[i][j] = other.data[i][j];
        }
    }
}

RationalMatrix::~RationalMatrix() {
    for (int i = 0; i < rows; ++i) {
        delete[] data[i];
    }
    delete[] data;
}

int RationalMatrix::getRowCount() { return rows; }
int RationalMatrix::getColCount() { return cols; }

Rational& RationalMatrix::getElement(int i, int j) {
    return data[i][j];
}

void RationalMatrix::multiplyRow(int row, const Rational& scalar) {
    for (int j = 0; j < cols; ++j) {
        data[row][j] = data[row][j] * scalar;
    }
}

void RationalMatrix::addRow(int targetRow, int sourceRow, const Rational& scalar) {
    for (int j = 0; j < cols; ++j) {
        data[targetRow][j] = data[targetRow][j] + (data[sourceRow][j] * scalar);
    }
}

RationalMatrix& RationalMatrix::operator=(const RationalMatrix& other) {
    if (this != &other) {
        for (int i = 0; i < rows; ++i) {
            delete[] data[i];
        }
        delete[] data;

        rows = other.rows;
        cols = other.cols;
        data = new Rational*[rows];
        for (int i = 0; i < rows; ++i) {
            data[i] = new Rational[cols];
            for (int j = 0; j < cols; ++j) {
                data[i][j] = other.data[i][j];
            }
        }
    }
    return *this;
}