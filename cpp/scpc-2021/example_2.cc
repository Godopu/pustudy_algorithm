/*
You should use the standard input/output

in order to receive a score properly.

Do not use file input and output

Please be very careful.
*/

#include <iostream>
#include <string>
#include <vector>

using namespace std;

int Answer;

void solve(vector<int> &A, vector<int> &B, int N, int t) {
  for (int i = 0; i < N - t; i++) {
    if (i - t < 0) {
      if (B.at(i) == 1)
        A[i + t] = 1;
      if (B.at(N - 1 - i) == 1)
        A[N - 1 - i - t] = 1;
    } else {
      if (B.at(i) == 1) {
        if (A[i - t] == 1 || A[i + t] == 1)
          continue;

        if (B[i + 2 * t] == 1 || i + 2 * t >= N)
          A[i + t] = 1;
        else
          A[i - t] = 1;
      }
    }
  }
}

int main(int argc, char **argv) {
  int T, test_case;
  int N, t;
  string strB;

  /*
     The freopen function below opens input.txt file in read only mode, and
     afterward, the program will read from input.txt file instead of
     standard(keyboard) input. To test your program, you may save input data in
     input.txt file, and use freopen function to read from the file when using
     cin function. You may remove the comment symbols(//) in the below statement
     and use it. Use #include<cstdio> or #include <stdio.h> to use the function
     in your program. But before submission, you must remove the freopen
     function or rewrite comment symbols(//).
   */

  // freopen("input.txt", "r", stdin);

  cin >> T;
  for (test_case = 0; test_case < T; test_case++) {
    vector<int> A, B;

    cin >> N;
    cin >> t;
    cin >> strB;
    for (int i = 0; i < N; i++) {
      B.push_back(strB.at(i) - '0');
      A.push_back(0);
    }

    solve(A, B, N, t);
    /////////////////////////////////////////////////////////////////////////////////////////////
    /*
       Implement your algorithm here.
       The answer to the case will be stored in variable Answer.
     */
    /////////////////////////////////////////////////////////////////////////////////////////////

    // Print the answer to standard output(screen).
    cout << "Case #" << test_case + 1 << endl;
    for (int i = 0; i < N; i++) {
      cout << A[i];
    }

    cout << endl;
  }

  return 0; // Your program should return 0 on normal termination.
}