#include<iostream>
#include<string>
//#include<memory.h>
using namespace std;

int main()
{
	//use hash
	string j = "aA";
	string s = "aAAbbbb";
	bool a[150];
	memset(a, 0, sizeof(a));
	int count = 0;
	for (int i = 0; i < j.length(); i++)
	{
		a[j[i]] = true;
	}
	for (int i = 0; i < s.length(); i++)
	{
		if (a[s[i]])
			count++;
	}
	cout << count << endl;
	return 0;
}