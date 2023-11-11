
#include <iostream>
#include <fstream>
int main () {

    std::ifstream input("test_input.txt");
    if (!input.is_open()) {
        std::cerr << "failed to open file" << std::endl;

        return 1;
    }

    int numberGroupOne = 0;
    std::string line;

    while (std::getline(input, line)) {
        int num = 0;
        if (line.empty()) {
            num = 0;
        } else {
            num = std::stoi(line); 
        }
            numberGroupOne += num;
    }

    input.close();

    std::cout << numberGroupOne << std::endl;
    return 0;
}
