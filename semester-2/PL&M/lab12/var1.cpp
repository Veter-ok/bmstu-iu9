#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <iostream>
#include <dirent.h>
#include <fstream>

using namespace std;

int sumInFile(string name) {
    ifstream f;
    int sum = 0;
    f.open(name);
    string line;
	while (getline(f,line)){
        sum += stoi(line);
	}
    return sum;
}

void writeInFile(vector<pair<string, int>> data) {
    ofstream f;
    f.open("testVar1/sum.txt");
    for (int i = 0; i < data.size(); i++) {
        f << data[i].first << " " << data[i].second << endl;
    }
}

int main(int argc, char** argv) {
    vector<pair<string, int>> ans;

    DIR* mainFolder = opendir(argv[1]);
    if (mainFolder == NULL) {
        cout << "wrong folder" << endl;
        return 0;
    }

    dirent* curDir;
    while ((curDir = readdir(mainFolder)) != NULL) {
        string nameFile = curDir->d_name;
        if (nameFile.size() > 4 && nameFile.substr(nameFile.size() - 3, nameFile.size() - 1) == "txt") {
            pair<string, int> p = {nameFile, sumInFile(string(argv[1]) + "/" + nameFile)};
            ans.push_back(p);
        }
    }

    sort(ans.begin(), ans.end());
    writeInFile(ans);

    return 0;
}