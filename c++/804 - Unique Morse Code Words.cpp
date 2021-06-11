#include<iostream>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;

int main()
{
    string Morse[] = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
	vector<string> words;
	vector<string> morses;
	words.push_back("gin"); words.push_back("zen"); words.push_back("gig"); words.push_back("msg");
	for (auto iter = words.cbegin(); iter < words.cend(); iter++) {
		string morse = "";
		for (int i = 0; i < (*iter).size(); i++) {
			morse += Morse[(*iter)[i] - 97];
		}
		vector<string>::iterator iVector = find(morses.begin(), morses.end(), morse);
		if (iVector == morses.end()) {
			morses.push_back(morse);
		}
	}
	cout << morses.size() << endl;
	return 0;
}