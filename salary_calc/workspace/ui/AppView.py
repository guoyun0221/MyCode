from tkinter import *
from tkinter import ttk, filedialog, messagebox
import os.path
from typing import Callable, List

from service.Model import Operation


class AppView:
    # ---------------- constant -----------------
    BASIC_SIZE = 10

    # ---------------- variable -----------------
    def __init__(self, operation_function: Callable[[str, str, str], None]):
        """ App overview window. user select operation, src_path, dst_path in this window. \n
        After user submit, corresponding service will be executed \n
        :argument operation_function function to be called when user submit.
            this function requires three params - op_type: str, src_path: str, dst_path: str"""
        # UI window and weights
        self.window = None
        self.op_type_combo = None
        self.src_text_entry = None
        self.dst_text_entry = None
        self.ret_text = None
        self.submit_button = None
        self.operation_function = operation_function

    def __src_path_browser_button_clicked(self):
        # operations whose src is a single file
        if self.op_type_combo.get() == Operation.OP_TYPE_EXPORT_ITEMS \
                or self.op_type_combo.get() == Operation.OP_TYPE_PROJECT_COST:
            src_path = filedialog.askopenfilename()
        else:
            src_path = ""
        self.src_text_entry.delete(0, END)
        self.src_text_entry.insert(0, src_path)
        # print("file: ", self.src_text_entry.get())

    def __dst_path_browser_button_clicked(self):
        # export data dst is a directory
        if self.op_type_combo.get() == Operation.OP_TYPE_EXPORT_ITEMS or \
             self.op_type_combo.get() == Operation.OP_TYPE_PROJECT_COST:
            dst_path = filedialog.askdirectory()
        else:
            dst_path = ""
        self.dst_text_entry.delete(0, END)
        self.dst_text_entry.insert(0, dst_path)

    def __submit_button_clicked(self):
        op_type = self.op_type_combo.get()
        src_path = self.src_text_entry.get()
        dst_path = self.dst_text_entry.get()
        # ------- debug ---------
        # print("submit op: ", op_type)
        # print("src path: ", src_path)
        # print("dst path:", dst_path)
        # ---------------------
        # check all fields legal
        # check op type
        if op_type not in Operation.OP_TYPE_LIST:
            messagebox.showerror("错误", "无效的操作类型")
        elif src_path is None or src_path == "":
            messagebox.showinfo("选择文件", "请选择源文件路径")
        elif (dst_path is None or dst_path == "" or not os.path.exists(dst_path)):
            messagebox.showerror("错误", "无效的目标文件夹路径")
        else:
            # check pass
            try:
                self.operation_function(op_type, src_path, dst_path)
                # if the operation do not open a new window, a message box is needed
                # to tell user operation is done
                if op_type == Operation.OP_TYPE_PROJECT_COST:
                    messagebox.showinfo("成功", op_type + "完成。")
            except Exception as e:
                messagebox.showerror("error", "exception in operation_function: " + repr(e))

    def draw(self):
        # init window
        self.window = Tk()
        self.window.title("人工成本分析核算")
        self.window.geometry(f"{AppView.BASIC_SIZE * 16 * 4}x{AppView.BASIC_SIZE * 9 * 4}")
        # label for operation type
        op_type_label = Label(self.window, text="选择操作类型", font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        # layout for it
        # op_type_label.place(x=AppView.BASIC_SIZE * 4, y=AppView.BASIC_SIZE * 4)
        op_type_label.place(relx=0.2, rely=0.15, anchor=CENTER)
        # op_type_label.grid(column=0, row=0, columnspan=2)

        # operation type combo
        self.op_type_combo = ttk.Combobox(self.window, width=AppView.BASIC_SIZE * 3,
                                          font=("Helvetica", int(AppView.BASIC_SIZE * 1.4)))
        # import op types to combo options
        self.op_type_combo['values'] = Operation.OP_TYPE_LIST
        # set default operation
        self.op_type_combo.current(0)
        # combo select listener -> not need anymore
        # self.op_type_combo.bind('<<ComboboxSelected>>', self.get_selected_operation)
        # layout for it
        # self.op_type_combo.place(x=AppView.BASIC_SIZE * 20, y=AppView.BASIC_SIZE * 4.2)
        self.op_type_combo.place(relx=0.6, rely=0.15, anchor=CENTER)
        # self.op_type_combo.grid(column=3, row=0, columnspan=2)

        # src file path and dst file path
        # label for src and dst path
        src_path_label = Label(self.window, text="源文件路径", font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        dst_path_label = Label(self.window, text="目标文件路径", font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        # layout for label
        # src_path_label.place(x=AppView.BASIC_SIZE * 4, y=AppView.BASIC_SIZE * 12)
        # dst_path_label.place(x=AppView.BASIC_SIZE * 4, y=AppView.BASIC_SIZE * 17)
        src_path_label.place(relx=0.2, rely=0.4, anchor=CENTER)
        dst_path_label.place(relx=0.2, rely=0.55, anchor=CENTER)
        # src_path_label.grid(column=0, row=1)
        # dst_path_label.grid(column=1, row=1)

        # text
        self.src_text_entry = Entry(self.window, width=int(AppView.BASIC_SIZE * 3.3),
                                    font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)))
        self.dst_text_entry = Entry(self.window, width=int(AppView.BASIC_SIZE * 3.3),
                                    font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)))
        # layout
        # self.src_text_entry.place(x=AppView.BASIC_SIZE * 20, y=AppView.BASIC_SIZE * 12.5)
        # self.dst_text_entry.place(x=AppView.BASIC_SIZE * 20, y=AppView.BASIC_SIZE * 17.5)
        self.src_text_entry.place(relx=0.55, rely=0.4, anchor=CENTER)
        self.dst_text_entry.place(relx=0.55, rely=0.55, anchor=CENTER)
        # file browser button
        src_path_browser_button = Button(self.window, text="选择", command=self.__src_path_browser_button_clicked,
                                         font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)),
                                         padx=AppView.BASIC_SIZE)
        dst_path_browser_button = Button(self.window, text="选择", command=self.__dst_path_browser_button_clicked,
                                         font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)),
                                         padx=AppView.BASIC_SIZE)
        # layout for it
        # src_path_browser_button.place(x=AppView.BASIC_SIZE * 50, y=AppView.BASIC_SIZE * 12)
        # dst_path_browser_button.place(x=AppView.BASIC_SIZE * 50, y=AppView.BASIC_SIZE * 17)
        src_path_browser_button.place(relx=0.88, rely=0.4, anchor=CENTER)
        dst_path_browser_button.place(relx=0.88, rely=0.55, anchor=CENTER)

        # submit button
        self.submit_button = Button(self.window, text="确定", command=self.__submit_button_clicked,
                                    font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)),
                                    padx=AppView.BASIC_SIZE)
        # self.submit_button.config(state=DISABLED)
        # layout
        self.submit_button.place(relx=0.5, rely=0.8, anchor=CENTER)

        # show window
        self.window.mainloop()
