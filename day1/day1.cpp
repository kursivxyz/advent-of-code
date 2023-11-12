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
    std::vector<int> numberVector = {0};

    while (std::getline(input, line)) {
        if (line.empty()) {
            if (numberVector.at(0) < num) {
                numberVector.push_back(num);
                num = 0;
            }
        } else {  
            num += std::stoi(line);
        }
    }

    for (int element : numberVector) {
        std::cout << element << std::endl;
    }
    input.close();
    return 0;
}
