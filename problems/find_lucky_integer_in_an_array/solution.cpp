using namespace std;

class Solution {
public:
    int findLucky(vector<int>& arr) {
        map<int, int> maps;
        for (vector<int>::iterator i = arr.begin(); i != arr.end(); i++)
        {
            int num = *i;
            pair<int, int> temp(num, 1);
            if (maps.find(num) != maps.end())
            {
                temp.second = (*maps.find(num)).second + 1;
            }
            maps[temp.first] = temp.second;
        }

        int res = -1;
        for (map<int, int>::iterator iter = maps.begin(); iter != maps.end(); ++iter)
        {
            pair<int, int> temp = *iter;
            if (temp.first == temp.second && temp.first > res)
            {
                res = temp.first;
            }
        }

        return res;
    }
};