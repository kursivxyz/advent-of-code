// First column - opponent is gonna play: A - rock, B - paper, C - scissors
// Second column - what you should play?: X - rock, Y - paper, Z - scissors 
// Winner = highest score, total score is score for each round. 
// Score per round is score for shape, 1 - rock, 2 - paper, 3 - scissors + outcome (0 if loss, 3 if draw, 6 if win)
//
// Round 1: opponent -> A, you -> Y, = 8 score (paper + win)
// Round 2: opponent -> B, you -> X, = 1 score (rock + loss)
// Round 3: opponent -> C, you -> Z, = 6 score (scissor + draw)
// Total score: 15

#include <algorithm>
#include <cctype>
#include <fstream>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;


map<string, int> initPlayScoreMap() {
    map<string, int> playScoreMap;
    playScoreMap["rock"] = 1;
    playScoreMap["paper"] = 2;
    playScoreMap["scissors"] = 3;
    return playScoreMap;
}

map<string, int> initResultScoreMap() {
    map<string, int> resultScoreMap;
    resultScoreMap["LOSS"] = 0;
    resultScoreMap["DRAW"] = 3;
    resultScoreMap["WIN"] = 6;
    return resultScoreMap;
}

string trimAllWhitespace(const string& str) {
    string result = str;
    result.erase(remove_if(result.begin(), result.end(), ::isspace), result.end());
    return result;
}

int main() {
    vector<int> roundScores;
    ifstream input("../../input/day2_test_input.txt");
    if(!input.is_open()) {
        cerr << "failed to open file." << endl;
        return 1;
    }

    string line;

    map<string, int> playScores = initPlayScoreMap();
    map<string, int> resultScores = initResultScoreMap();

    while(getline(input, line)) {
        line = trimAllWhitespace(line);
        
        // TEMP TEST, PLS REMOVE 
        if (line.substr(0, 1) =="A" ) {
            cout << "hehe" << endl;
        }

        if (line.substr(1, 1) == "Y") {
            cout << "hihi" << endl;
        }
    }
}
