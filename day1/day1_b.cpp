#include <algorithm>
#include <iostream>
#include <fstream>
#include <numeric>
#include <vector>



int main () {
    std::ifstream input("input.txt");
    if (!input.is_open()) {
        std::cerr << "failed to open file." << std::endl;
        return 1;
    }

    std::string line;
    
    int num = 0;
    int highestNumber = 0;
    std::vector<int> elfs;

    while (std::getline(input, line)) {
        if (line.empty()) {
            highestNumber = std::max(highestNumber, num);
            // match num against current entries in list, if not the smallest, remove the smallest?
            auto minEntry = std::min_element(elfs.begin(), elfs.end());
            if (elfs.size() < 3 || num > *std::min_element(elfs.begin(), elfs.end())) {
                if (elfs.size() == 3) {
                    elfs.erase(minEntry);
                }
                elfs.push_back(num);
            } 
            num = 0;
        } else {  
            num += std::stoi(line);
        }
    }

    std::cout << elfs.size() << std::endl;
    int sum = std::accumulate(elfs.begin(), elfs.end(), 0);
    std::cout << sum << std::endl;
    
    input.close();
    return 0;
}
