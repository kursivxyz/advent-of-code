
#include <iostream>
#include <fstream>
int main () {
	
	std::ifstream input("test_input.txt");
	if (input.is_open()) {
		std::cerr << "failed to open file" << std::endl;

		return 1;
	}

	int sum = 0;

	return 0;
}
