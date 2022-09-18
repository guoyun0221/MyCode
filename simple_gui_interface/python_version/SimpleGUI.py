import sys
from tkinter import *
from tkinter import scrolledtext
from typing import Callable, List


class SimpleGUI:

    def __init__(self, title: str, input_prompts: List[str], callback_func: Callable[[List[str]], None], button_text: str = "submit", basic_size: int = 10) -> None:
        '''To create simpleGui.\n
        Args: \n
            tille.  title of this window to be shown. \n
            input_prompts.  list of string prompts for user inputs.
                            each element in the list corresponds to an input \n
            callback_func.    A func to be called. when user click submit button, 
                            the function will be called. this func must require one
                            string list argument to get user inputs. \n
            button_text.    text of submit button. default is "submit".\n
            basic_size.     basic size for component. default is 10 '''
        # Parameters
        self.title = title
        self.input_prompts = input_prompts
        self.button_text = button_text
        self.basic_size = basic_size
        self.callback_func = callback_func
        # Gui window and weights
        self.window = None
        self.input_text_entry = []
        self.output_text_area = None

    def create_window(self):
        '''create a gui window'''
        # init window
        self.window = Tk()
        self.window.title(self.title)
        self.window.geometry(f"{self.basic_size * 16 * 10}x{self.basic_size * 9 * 10}")
        # create and put components to window
        self.__create_title_label()
        self.__create_input_panel()
        self.__create_output_text_area()
        self.__create_submit_button()
        # window closed handler
        self.window.protocol("WM_DELETE_WINDOW", self.__window_closed)
        # show window
        self.window.mainloop()

    def write_output(self, s: str):
        '''Write output to output text area'''
        self.output_text_area.insert(INSERT, s)
        self.window.update()

    def __window_closed(self):
        # close the window 
        self.window.destroy()
        # shutdown the app after window closed
        sys.exit(0)

    def __create_title_label(self):
        title_label = Label(self.window, text=self.title, font=('consolas', int(self.basic_size * 7)))
        # layout for it
        title_label.place(relx=0.5, rely=0.1, anchor=CENTER)
    
    def __create_input_panel(self):
        input_panel = PanedWindow(self.window)
        # add input prompts and text entry pairs
        for i in range(len(self.input_prompts)):
            # input prompt label
            label = Label(input_panel, text=self.input_prompts[i], font=('consolas', int(self.basic_size * 2)))
            label.grid(column=0, row=i, pady=self.basic_size * 2)
            # input text entry
            text = Entry(input_panel, width=int(self.basic_size * 2), font=("consolas", int(self.basic_size * 2)))
            text.grid(column=1, row=i, pady=self.basic_size * 2)
            # add text entry to self
            self.input_text_entry.append(text)
        # layout for the panel
        input_panel.place(relx=0.08, rely=0.55, anchor=W)

    def __create_output_text_area(self):
        self.output_text_area = scrolledtext.ScrolledText(self.window, width=self.basic_size * 4, height=self.basic_size, font=("consolas", int(self.basic_size * 2)))
        self.output_text_area.place(relx=0.95, rely=0.55, anchor=E)

    def __create_submit_button(self):
        button = Button(text=self.button_text, command=self.__submit_button_clicked, font=("consolas", int(self.basic_size * 2)))
        button.place(relx=0.5, rely=0.98, anchor=S)

    def __submit_button_clicked(self):
        # get usr inputs
        inputs = []
        for i in range(len(self.input_text_entry)):
            inputs.append(self.input_text_entry[i].get())
        # call the callback function, and inputs is the argument
        self.callback_func(inputs)
        