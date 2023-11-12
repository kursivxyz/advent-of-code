#include <iostream>
#include <fstream>
#include <vector>



int main () {
    std::ifstream input("test_input.txt");
    if (!input.is_open()) {
        std::cerr << "failed to open file." << std::endl;
        return 1;
    }

    std::string line;
    
    int num = 0;
    int highestNumber = 0;

    while (std::getline(input, line)) {
        if (line.empty()) {
            if (num > highestNumber) {
                highestNumber = num;
                num = 0;
            }
        } else {  
            num += std::stoi(line);
        }
    }

    std::cout << highestNumber << std::endl;
    
    input.close();
    return 0;
}
