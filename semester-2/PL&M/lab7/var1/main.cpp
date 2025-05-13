#include "declaration.h"
#include <iostream>

void printMatrix(RationalMatrix& matrix) {
    for (int i = 0; i < matrix.getRowCount(); ++i) {
        for (int j = 0; j < matrix.getColCount(); ++j) {
            Rational& r = matrix.getElement(i, j);
            std::cout << r.getNumerator() << "/" << r.getDenominator() << "\t";
        }
        std::cout << std::endl;
    }
    std::cout << std::endl;
}

int main() {
    RationalMatrix mat(2, 3);
    
    mat.getElement(0, 0) = Rational(1, 2);
    mat.getElement(0, 1) = Rational(3, 4);
    mat.getElement(0, 2) = Rational(5, 6);
    mat.getElement(1, 0) = Rational(7, 8);
    mat.getElement(1, 1) = Rational(9, 10);
    mat.getElement(1, 2) = Rational(11, 12);

    std::cout << "Исходная матрица:" << std::endl;
    printMatrix(mat);

    mat.multiplyRow(0, Rational(2, 1));
    std::cout << "После умножения строки 0 на 2:" << std::endl;
    printMatrix(mat);

    mat.addRow(1, 0, Rational(1, 2));
    std::cout << "После добавления половины строки 0 к строке 1:" << std::endl;
    printMatrix(mat);

    return 0;
}