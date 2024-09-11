import sys
import networkx as nx
import matplotlib.pyplot as plt
import re

##### usage #####
# 1. objdump -d test.elf > test.dis
# 2. python func_analyzer.py test.dis start_function

##### global const #####
# jump instructions we can refer from disassembly
static_jumps = ["jal"]
# jump address determined at runtime
dynamic_jumps = ["jalr"]

# function name line pattern
func_name_pattern = r'^[0-9a-fA-F]+\s+<([\w_\.]+)>:$'
# instruction line pattern
inst_pattern = r'^\s*[0-9a-fA-F]+:\s+([0-9a-fA-F]+\s+)+.*$'
# jal instruction jump label pattern
jal_label_pattern = r'jal\s+0x[0-9a-fA-F]+\s+<([\w_\.]+)>'

# instruction operation type start index in line string
operation_type_index = 24

#### global variables ####
dynamic_jump_insts = []

##### functions #####
def extract_function_names(filename):
    """ get all function names from disassembly """
    # function name line match pattern
    pattern = re.compile(func_name_pattern)
    
    function_names = []
    # read file by lines
    with open(filename, 'r') as file:
        for line in file:
            # find function name line
            match = pattern.match(line)
            if match:
                # add fucntion name to list
                function_names.append(match.group(1))
    
    return function_names

def get_inst_op_type(line):
    op_type = ""
    end_index = line.find('\t', operation_type_index)
    if end_index == -1:
        op_type = line[operation_type_index:]
    else:
        op_type = line[operation_type_index:end_index]
    return op_type

def get_jal_label(str, line_num):
    pattern = re.compile(jal_label_pattern)
    match = pattern.search(str)
    if match:
        return match.group(1)
    else:
        print("unexpected unmatched jal instruction: ", str.strip(), "; line number: ", line_num)
        sys.exit(1)

def add_edge(graph, from_node, to_node):
    """ check and add egde by function calling """
    if from_node == "":
        # thing went wrong
        print("from_node is empty")
        sys.exit(1)
    # we do not care recursion
    elif from_node == to_node:
        return
    # Check if adding the edge would create a cycle
    elif nx.has_path(graph, to_node, from_node):
        print(f"Skipping edge {from_node} -> {to_node} to avoid a cycle")
    else:
        # No cycle would be created, safe to add the edge
        graph.add_edge(from_node, to_node)

def extract_subgraph(graph, start_node):
    """ Extract the subgraph reachable from the start_node """
    sub_nodes = nx.dfs_tree(graph, start_node).nodes
    # Create a subgraph based on these nodes
    subgraph = graph.subgraph(sub_nodes)
    return subgraph

def analyze_call_graph(graph, start_node):
    """ get call depth information from graph"""
    # get Maximum call depth from start_node
    if nx.is_directed_acyclic_graph(graph):
        longest_path = nx.dag_longest_path(graph)
        print(f"Longest call chain of all graph: {len(longest_path)} : {longest_path}")
        print("")
        subgraph = extract_subgraph(graph, start_node)
        sub_longest_path = nx.dag_longest_path(subgraph)
        print(f"Longest call chain from {start_node}: {len(sub_longest_path)} : {sub_longest_path}")
    else:
        print("The graph contains cycles, so a longest path cannot be determined in a DAG.")
        sys.exit(1)

def draw_graph(call_graph):
    nx.draw(call_graph, with_labels=True)
    plt.show()

def build_call_graph(filename, functions):
    """ build function calling graph """
    # create directed graph
    call_graph = nx.DiGraph()
    
    # add all function names as nodes
    for func in functions:
        call_graph.add_node(func)
    
    # function name pattern
    pattern_func_name = re.compile(func_name_pattern)
    pattern_inst = re.compile(inst_pattern)

    # mark the function we are in
    current_func = ""
    # mark line number
    line_num = 0

    # read disassembly file by lines
    with open(filename, 'r') as file:
        for line in file:
            line_num += 1
            # try to match function name
            match_func_name = pattern_func_name.match(line)
            if match_func_name:
                current_func = match_func_name.group(1)
                continue
            # try to match instruction lines
            match_inst = pattern_inst.match(line)

            if match_inst:
                # get instruction operation
                op_type = get_inst_op_type(line)
                # print("op type: ", op_type)
                # check if it's jump instrucion
                if op_type in static_jumps:
                    # get jump target function name
                    jal_target = get_jal_label(line[operation_type_index:], line_num)
                    # create edge for this jal
                    add_edge(call_graph, current_func, jal_target)
                elif op_type in dynamic_jumps:
                    dynamic_jump_insts.append(
                        "line numer: " + str(line_num) + "; in function: " + current_func + "; content: " + line.strip())
    
    return call_graph

def main():
    # get disassembly file name
    if len(sys.argv) < 3:
        sys.exit('usage: python func_analyzer.py disassembly_file.dis start_function')
    filename = sys.argv[1]
    start_function = sys.argv[2]

    print("\n------------ build graph from disassembly ---------")

    # get all nodes (fucntions)
    functions = extract_function_names(filename)
    # print(functions)

    # build graph. 
    call_graph = build_call_graph(filename, functions)

    # analyze graph
    print("\n--------------- calling analysis -----------------")
    analyze_call_graph(call_graph, start_function)

    # get sub-graph reachable from the start_function
    call_graph_sub = extract_subgraph(call_graph, start_function)

    # print all dynamic_jump_insts
    print("\n-------------- dynamic jump instsructions -----------------")
    for inst in dynamic_jump_insts:
        print(inst)

    # draw graph
    draw_graph(call_graph_sub)
    
if __name__ == "__main__":
    main()
