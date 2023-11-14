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
#include <numeric>
#include <string>
#include <tuple>
#include <vector>

using namespace std;


map<string, int> initPlayScoreMap() {
    map<string, int> playScoreMap;
    playScoreMap["rock"] = 1;
    playScoreMap["paper"] = 2;
    playScoreMap["scissors"] = 3;
    return playScoreMap;
}

int retrieveScores(map<string, int> map, string key) {
    int point;
    auto score = map.find(key);

    if (score != map.end()) {
        return score->second;
    } 
    return 0;
}

map<string, int> initResultScoreMap() {
    map<string, int> resultScoreMap;
    resultScoreMap["loss"] = 0;
    resultScoreMap["draw"] = 3;
    resultScoreMap["win"] = 6;
    return resultScoreMap;
}

string played(string playCode) {
    map<string, string> map;
    map["AX"] = "rock";
    map["BY"] = "paper";
    map["CZ"] = "scissors";

    for (const auto& pair: map) {
        if (pair.first.find(playCode) != string::npos) {
            return pair.second;
        }
    }
    return "";
}

string trimAllWhitespace(const string& str) {
    string result = str;
    result.erase(remove_if(result.begin(), result.end(), ::isspace), result.end());
    return result;
}

string determineResult(string opponent, string player) {
    if (opponent == player ) return "draw";
    if ((opponent== "rock" && player == "scissors") ||
            (opponent == "paper" && player == "rock") ||
            (opponent == "scissors" && player == "paper")) {
        return "loss";
    } 
    return "win";
}

int main() {
    int totalScore = 0;
    int roundNumber = 0;
    ifstream input("../../input/day2_input.txt");
    if(!input.is_open()) {
        cerr << "failed to open file." << endl;
        return 1;
    }

    string line;
    map<string, int> playScores = initPlayScoreMap();
    map<string, int> resultScores = initResultScoreMap();

    while(getline(input, line)) {
        roundNumber++;
        line = trimAllWhitespace(line);
        int roundScore = 0;

        string playedPlayed = played(line.substr(1,1));
        roundScore += retrieveScores(playScores, playedPlayed);

        string result = determineResult(played(line.substr(0,1)), playedPlayed);
        roundScore += retrieveScores(resultScores, result);        

        totalScore += roundScore;
        cout << "Round " << roundNumber << " score: " << roundScore << endl;
    }
    cout << "Total score: " << totalScore<< endl;
}
