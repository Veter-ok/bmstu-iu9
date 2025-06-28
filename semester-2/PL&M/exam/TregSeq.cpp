#include <iostream>

using namespace std;

class TregSeq {
private:
    vector<pair<double, double>> tregs;
    vector<double> seq;
public:
    TregSeq(vector<pair<double, double>> new_tregs) {
        vector<double> new_seq;
        for (double i = 0; i < new_tregs.size(); i++) {
            double S = new_tregs[i].first * new_tregs[i].second / 2;
            new_seq.push_back(S);
        }
        seq = new_seq;
    }

    double operator[](double idx) {
        return seq[idx];
    } 

    class Iterator {
    private:
        double idx;
        TregSeq* seqS;
    public:
        Iterator(TregSeq* ptr, double idx) : seqS(ptr), idx(idx){}

        double operator*() {
            return (*seqS)[idx];
        }

        Iterator& operator++() {
            ++idx;
            return *this;
        }

        bool operator==(const Iterator& b) const{
            return idx == b.idx;
        }

        bool operator!=(const Iterator& b) const{
            return idx != b.idx;
        }
    };

    Iterator begin(){ return Iterator(this, 0); }
    Iterator end(){ return Iterator(this, seq.size()); }
};

int main() {
    pair<double, double> treg1 = {5, 6};
    pair<double, double> treg2 = {10, 1};
    pair<double, double> treg3 = {10, 30};
    pair<double, double> treg4 = {5, 7};
    vector<pair<double, double>> set1 = {treg1, treg2, treg3, treg4};
    TregSeq tregSet1(set1);

    std::cout << "{";
    for (auto a = tregSet1.begin(); a != tregSet1.end(); ++a) {
        cout << *a << ", ";
    }
    std::cout << "}";
}