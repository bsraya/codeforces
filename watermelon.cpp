#include <iostream>

using std::cout;
using std::endl;
using std::cin;

int main(int argc, char const *argv[])
{
    int weight = 0;
    cin >> weight;
    if(weight <= 2 or (weight%2 != 0))
        cout << "NO";
    else 
        cout << "YES";
    return 0;
}
