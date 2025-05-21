#include <iostream>

using namespace std;

class NumSet {
private:
    vector<vector<int>> set;
public:
    NumSet(vector<int> s) {
        vector<vector<int>> new_set;
        for (int i = 0; i < s.size(); i++){
            for (int j = i+1; j < s.size(); j++){
                for (int k = j+1; k < s.size(); k++){
                    auto a = s[i] * s[i];
                    auto b = s[j] * s[j];
                    auto c = s[k] * s[k];
                    if (a == b + c || b == a + c || c == a + b) {
                        vector<int> tr;
                        tr.push_back(s[i]);
                        tr.push_back(s[j]);
                        tr.push_back(s[k]);
                        new_set.push_back(tr);
                    }
                }
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
        NumSet* numSet;
    public:
        Iterator(NumSet* ptr, int idx) : numSet(ptr), idx(idx){}

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
    vector<int> set1 = {3, 4, 5, 7, 12, 13, 43};
    NumSet numSet1(set1);

    for (auto a = numSet1.begin(); a != numSet1.end(); ++a) {
        std::cout << "(" << (*a)[0] << ", " << (*a)[1] << ", " << (*a)[2] << ")" << endl;
    }
}