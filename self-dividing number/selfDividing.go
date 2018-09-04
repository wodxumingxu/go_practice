class Solution {
public:
    vector<int> selfDividingNumbers(int left, int right) {
        vector<int> ret;
        for(int i = left; i <= right; ++i){
            
            int next = i / 10;
            int divided = i % 10;
            bool qualify = (divided == 0) ? false : (0 == (i % divided));
            
            while((0 != divided) && (0 != next) && qualify){
                divided = next % 10;
                qualify = qualify = (divided == 0) ? false : (0 == (i % divided));
                next = next / 10;
            }
            
            if(qualify){
                ret.push_back(i);
            }
        }
        return ret;
    }
};
