#include <iostream>
#include <string>
#include <limits>
#include <cstdlib>

void displayDisclaimer() {
    std::cout << "\nDisclaimer:\n";
    std::cout << "This software is an open-source project, prohibited for any commercial use!!!\n";
    std::cout << "This tool is for learning, exchange, and technical testing purposes only, aimed at exploring web video parsing technology. It does not provide any content storage or dissemination services. Do not use it for any illegal purposes. Users should bear all legal responsibilities arising from the misuse of this tool. Please respect film and television copyrights and support genuine content. If your rights are infringed, please contact us for timely removal.\n";
    std::cout << "The developer assumes no direct or indirect liability for losses.\n";
}

int main() {
    std::string movieLink = "";
    int choice;

    const std::string IQIYI_URL = "https://www.iqiyi.com";
    const std::string TENCENT_URL = "https://v.qq.com";
    const std::string YOUKU_URL = "https://www.youku.com/";
    const std::string HITMUX_URL = "https://hitmux.top";
    const std::string PARSE_PREFIX = "https://jx.xmflv.cc/?url=";

    while (true) {
        std::cout << "\n--- VIP Video Parser ---\n";
        std::cout << "Current Video Link: " << (movieLink.empty() ? "None" : movieLink) << "\n\n";
        std::cout << "1. Enter Video Link\n";
        std::cout << "2. Clear Link\n";
        std::cout << "3. Open iQiyi\n";
        std::cout << "4. Open Tencent Video\n";
        std::cout << "5. Open Youku Video\n";
        std::cout << "6. Open Hitmux Official Site\n";
        std::cout << "7. Play VIP Video Now\n";
        std::cout << "8. Show Disclaimer\n";
        std::cout << "0. Exit\n";
        std::cout << "Enter your choice: ";

        std::cin >> choice;

        std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');

        switch (choice) {
            case 1:
                std::cout << "Please paste the video link: ";
                std::getline(std::cin, movieLink);
                std::cout << "Video link updated.\n";
                break;
            case 2:
                movieLink = "";
                std::cout << "Video link cleared.\n";
                break;
            case 3:
                std::cout << "Opening iQiyi: " << IQIYI_URL << "\n";
                break;
            case 4:
                std::cout << "Opening Tencent Video: " << TENCENT_URL << "\n";
                break;
            case 5:
                std::cout << "Opening Youku Video: " << YOUKU_URL << "\n";
                break;
            case 6:
                std::cout << "Opening Hitmux Official Site: " << HITMUX_URL << "\n";
                break;
            case 7:
                if (!movieLink.empty()) {
                    std::cout << "Playing VIP Video: " << PARSE_PREFIX << movieLink << "\n";
                } else {
                    std::cout << "Please enter a valid video link first!\n";
                }
                break;
            case 8:
                displayDisclaimer();
                break;
            case 0:
                std::cout << "Exiting application. Goodbye!\n";
                return 0;
            default:
                std::cout << "Invalid choice. Please try again.\n";
                break;
        }
        std::cout << "\nPress Enter to continue...";
        std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
    }

    return 0;
}
