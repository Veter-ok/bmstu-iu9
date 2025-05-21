#include <iostream>

using namespace std;

class NumberSet {
private:
    vector<vector<int>> set;
public:
    NumberSet(vector<int> a) {
        vector<vector<int>> new_set;
        for (int i = 0; i < a.size(); i++){
            vector<int> subsq;
            int sum = 0;
            for (int j = i; j < a.size(); j++){
                if (sum + a[j] <= 21) {
                    subsq.push_back(a[j]);
                    sum += a[j];
                }else{
                    break;
                }
            }
            if (subsq.size() > 0) {
                new_set.push_back(subsq);
            }
        }
        set = new_set;
    }

    vector<int> operator[](int idx) {
        return set[idx];
    } 

    class Iterator {
    private:
        int idx;
        NumberSet* numSet;
    public:
        Iterator(NumberSet* ptr, int idx) : numSet(ptr), idx(idx){}

        vector<int> operator*() {
            return (*numSet)[idx];
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
    Iterator end(){ return Iterator(this, set.size()); }
};

int main() {
    vector<int> set1 = {3, 4, 5, 43, 7, 12, 13};
    NumberSet numSet1(set1);

    for (auto a = numSet1.begin(); a != numSet1.end(); ++a) {
        std::cout << "{";
        for (int i = 0; i < (*a).size() - 1; i++){
            std::cout << (*a)[i] << ", ";
        }
        std::cout << (*a)[(*a).size() - 1] << "}" << endl;
    }
}