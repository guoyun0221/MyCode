#!/bin/env python3

import os
import subprocess

#### params
path_to_nm = '/path/to/nm'
path_to_lib = '/path/to/lib'
####

class FuncInfo:
    def __init__(self, cnt, min_size, max_size):
        # count function appearing times
        self.cnt = cnt
        # min and max size of implement of function
        self.min_size = min_size
        self.max_size = max_size

# func_name to FuncInfo map
func_map = {}
        
def extract_functions_from_static_lib(lib_path):
    try:
        # run nm to get symbol list
        nm_output = subprocess.check_output([path_to_nm,  '-g', "-C", '--print-size', '--defined-only', lib_path]).decode('utf-8')
    except subprocess.CalledProcessError as e:
        print(f"Error processing {lib_path}: {e}")
        return set()

    functions = set()
    
    with open(os.path.basename(lib_path) + ".txt", 'w') as f:
        # analyse symbols
        for line in nm_output.splitlines():
            # only collect .text symbol
            if " T " in line:
                # add fucntion name to set
                func_name = line[line.index(" T ") + 3:]
                functions.add(func_name)
                # get func size
                func_size = int(line.split(' ')[1], 16)
                
                # print current function to file
                f.write(f"{func_name};  size:{func_size}\n")
                
                # modify func_info in global func_map
                func_info = func_map.get(func_name, None)
                if (func_info is None):
                    func_info = FuncInfo(1, func_size, func_size)
                else :
                    func_info.cnt += 1
                    func_info.min_size = min(func_info.min_size, func_size)
                    func_info.max_size = max(func_info.max_size, func_size)
                func_map[func_name] = func_info
        # print summary
        f.write(f"\nExtracted {len(functions)} functions from {lib_path}\n")
    return functions

def main():
    # lib include path
    lib_path = path_to_lib
    all_functions = set()
    summary=""
    
    for root, dirs, files in os.walk(lib_path):
        for file in files:
            if file.endswith('.a'):
                lib_path = os.path.join(root, file)
                functions = extract_functions_from_static_lib(lib_path)
                all_functions.update(functions)
                print(f"Extracted {len(functions)} functions from {file}")
                summary += f"Extracted {len(functions)} functions from {file}\n"
                
    print(f"Extracted all unique functions: {len(all_functions)}")
    summary += f"Extracted all unique functions: {len(all_functions)}\n\n"
    
    with open('all_lib_functions.txt', 'w') as f:
        # print summary
        f.write(summary)
        # print all functions
        for func in sorted(all_functions):
            func_info = func_map[func]
            f.write(f"{func};  count:{func_info.cnt};  min_size:{func_info.min_size};  max_size:{func_info.max_size}\n")
    
if __name__ == "__main__":
    main()
