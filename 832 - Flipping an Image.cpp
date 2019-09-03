#include<iostream>
#include<string>
#include<vector>
#include<algorithm>
using namespace std;

class Solution {
public:
	vector<vector<int>> flipAndInvertImage(vector<vector<int>>& A) {
		vector<vector<int>> B;
		for (auto iter = A.cbegin(); iter < A.cend(); iter++) {
			vector<int> C;
			for (int i = (*iter).size() - 1; i >= 0; i--) {
				C.push_back(1 - (*iter)[i]);
			}
			B.push_back(C);
		}
		return B;
	}
};