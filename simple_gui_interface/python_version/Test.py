from SimpleGUI import SimpleGUI

def submit_recall_func(inputs):
    for i in range(len(inputs)):
        gui.write_output("idx: " + str(i) + "; value: " + inputs[i] + "\n")

gui = SimpleGUI("Test Gui Interface", ["input 1", "input 2"], submit_recall_func)

gui.create_window()


