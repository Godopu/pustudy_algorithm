#include <iostream>
#include <vector>

using namespace std;
void changeData(vector<int> &v) { v[0] = 100; }

int main() {
  vector<int> v;
  v.push_back(10);

  cout << v.at(0) << endl;
  changeData(v);
  cout << v.at(0) << endl;

  cout << v.size() << endl;

  return 0;
}