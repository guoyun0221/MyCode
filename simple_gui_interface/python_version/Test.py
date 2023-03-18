from typing import List
from SimpleGUI import SimpleGUI

def submit_callback_func(inputs: List[str]):
    gui.clean_output()
    for i in range(len(inputs)):
        gui.write_output("idx: " + str(i) + "; value: " + inputs[i] + "\n")

gui = SimpleGUI("Test Gui Interface", ["input 1", "input 2"], submit_callback_func)

gui.create_window()


