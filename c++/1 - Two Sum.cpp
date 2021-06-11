#include<iostream>
#include<string>
#include<vector>
#include <unordered_map>
//#include<memory.h>
using namespace std;

class Solution {
public:
	vector<int> twoSum(vector<int>& nums, int target) {
		unordered_map<int, int> mymap;
		vector<int> a;
		for (int i = 0; i < nums.size(); i++) {
			int x = target - nums[i];
			if (mymap.find(x) != mymap.end()) {
				a.push_back(mymap[x]);
				a.push_back(i);
				return a;
			}
			mymap.insert(pair<int, int>{nums[i], i});
		}
		return a;
	}
};