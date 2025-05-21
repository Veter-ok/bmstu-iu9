#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <iostream>
#include <dirent.h>
#include <fstream>

using namespace std;

string getLinkFromStr(string str) {
    bool linkStart = false;
    string ans = "";
    for (int i = 0; i < str.size(); i++) {
        if (str[i] == '"' || str[i] == '\'' ) {
            if (linkStart) {
                break;
            }else{
                linkStart = true;
                continue;
            }
        }
        if (linkStart) {
            ans += str[i];
        }
    }
    return ans;
}

void getLinks(vector<string> *links, string name) {
    ifstream f;
    f.open(name);
    string line;
	while (getline(f,line)){
        if (line.find("href") != string::npos) {
            links->push_back(getLinkFromStr(line));
        }
	}
}


void writeInFile(vector<string> data) {
    ofstream f;
    f.open("testVar2/links.txt");
    for (int i = 0; i < data.size(); i++) {
        f << data[i] << endl;
    }
}

int main(int argc, char** argv) {
    DIR* mainFolder = opendir(argv[1]);
    if (mainFolder == NULL) {
        cout << "wrong folder" << endl;
        return 0;
    }

    dirent* curDir;
    vector<string> links;
    while ((curDir = readdir(mainFolder)) != NULL) {
        string nameFile = curDir->d_name;
        if (nameFile.size() > 5 && nameFile.substr(nameFile.size() - 4, nameFile.size() - 1) == "html") {
            getLinks(&links, string(argv[1]) + "/" + nameFile);
        }
    }

    sort(links.begin(), links.end());
    writeInFile(links);

    return 0;
}