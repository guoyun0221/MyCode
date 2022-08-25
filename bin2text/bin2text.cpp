#include <iostream>
#include <fstream>
#include <string.h>
#include <stdio.h>
#include <string>
#include <regex>
#include <iomanip>

using namespace std;

int main(int argc, char *argv[]){

    if(argc < 3){
        cout << "usage: bin2text <src_file> <bytes_number_per_line>" << endl;
        return 0;
    }

    // some arguments and parameters 
    string input_name = argv[1];
    int bytes_per_line = stoi(argv[2]);
    
    // output file, convert .bin to .txt
    string output_name = input_name;
    output_name.replace(output_name.find("bin"), 3, "txt");

    // open src and dst file
    ifstream src(input_name ,ios::in|ios::binary);
    if(!src){
        cout<<"Open src file error!"<<endl;
        return 0;
    }

    ofstream dst;
    dst.open(output_name.c_str());
    if(!dst){
        cout << "Open dst file error!" << endl;
    }

    
    int count = 0;
    // handle data
    char buf = 0;
    while(src.read((char *)&buf, 1)){
        // write data
        // format is like: "0x01, 0x23, "
        dst << hex  << setfill('0') << setw(2) << "0x" << (unsigned int)(unsigned char)buf << ", ";
        
        // switch wo next line if count to bytes limit
        count++;
        if(count == bytes_per_line){
            dst << endl;
            count = 0;
        }
    }

    // close resources
    src.close();
    dst.close();
    return 0;
}
