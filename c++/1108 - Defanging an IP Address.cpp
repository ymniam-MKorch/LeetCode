#include<iostream>
#include<string>
#include<vector>
//#include<memory.h>
using namespace std;

int main()
{
	string a = "1.1.1.1";
	int index = a.find(".");
	while (index != a.npos) {
		a = a.replace(index, 1, "[.]");
		index = a.find(".", index+=2);
	}
	cout << a << endl;
	return 0;
}