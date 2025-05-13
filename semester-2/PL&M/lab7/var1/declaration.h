#ifndef DECLARATION_H
#define DECLARATION_H

class Rational {
private:
    int numerator; 
    int denominator;
    void normalize();
public:
    Rational(int num, int denom);
    Rational(const Rational& other);
    int getNumerator();
    int getDenominator();
    Rational operator+(Rational other);
    Rational operator*(Rational other);
    Rational& operator=(const Rational& other);
};


class RationalMatrix {
private:
    Rational** data;
    int rows;
    int cols;
public:
    RationalMatrix(int m, int n);
    RationalMatrix(const RationalMatrix& other);
    virtual ~RationalMatrix();
    int getRowCount();
    int getColCount();
    Rational& getElement(int i, int j);
    void multiplyRow(int row, const Rational& scalar);
    void addRow(int targetRow, int sourceRow, const Rational& scalar);
    RationalMatrix& operator=(const RationalMatrix& other);
};

#endif