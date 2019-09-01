#include<iostream>
#include<string>
#include<vector>
using namespace std;

int main()
{
	string S = "()()";
	string a = "";
	int lcount = 1;
	int rcount = 0;
	for (int i = 1; i < S.length(); i++) {
		if (S[i] == '(') {
			lcount++;
		}
		else{
			rcount++;
		}
		if (lcount == rcount) {
			i++;
			lcount = 1;
			rcount = 0;
		}
		else {
			a += S[i];
		}
	}
	cout << a << endl;
	return 0;
}