from tkinter import *
from tkinter import ttk, filedialog, scrolledtext, messagebox
from typing import List, Callable

from service.Model import SheetColumns, RowFilterType
from ui.AppView import AppView


class ExportDataView:
    def __init__(self):
        """ A window view for user to select export cols and rows """
        # arguments
        self.sheet_column_labels_list = None
        # ui weights
        self.window = None
        self.row_filter_content_area = None
        self.submit_button = None
        # use SheetColumn list to put cols tab weight
        self.cols_tab = []
        # some values
        # cols frame select all button last state,
        # to recognize which frame's select all button clicked
        self.chk_all_last_state = []
        self.row_type_selected = None
        # call export data function when submit
        self.do_export_data_function = None

    def set_do_export_data_function(self, do_export_data_function: Callable[[List[SheetColumns], int, str], None]):
        """ set callback function to be called when user submit
        :argument do_export_data_function like this -
         do_export(selected_cols: List[SheetColumns], row_filter_type: int, row_filter_content: str) -> None"""
        self.do_export_data_function = do_export_data_function

    def draw(self, sheet_columns: List[SheetColumns]):
        # clear fields.
        self.__reset_attributes()
        # update self.sheet_column_labels_list
        self.sheet_column_labels_list = sheet_columns
        # check callback function
        if self.do_export_data_function is None:
            raise Exception("No callback function to do export data")
        # check sheet_columns first
        if len(self.sheet_column_labels_list) == 0:
            messagebox.showerror("error", "empty sheet list")
        # draw window
        self.window = Tk()
        self.window.title("导出数据")
        self.window.geometry(f"{AppView.BASIC_SIZE * 16 * 9}x{AppView.BASIC_SIZE * 9 * 9}")

        # there are two part of this window
        # ------------------- part one to select excel columns ----------------------
        cols_panel = PanedWindow(self.window)
        # layout cols_panel
        cols_panel.pack(side='left', anchor='nw', padx=AppView.BASIC_SIZE * 4, pady=AppView.BASIC_SIZE * 4)
        # cols select label
        cols_label = Label(cols_panel, text="选择要导出的列",
                           font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        cols_label.pack(side='top')
        # get sheet number
        n_sheet = len(self.sheet_column_labels_list)
        # a notebook(tab) to hold all sheets
        sheets_tab = ttk.Notebook(cols_panel)
        sheets_tab.pack(fill=BOTH, expand=True, pady=AppView.BASIC_SIZE * 3)
        # tab header font size
        style = ttk.Style(sheets_tab)
        style.configure('TNotebook.Tab', font=("微软雅黑", int(AppView.BASIC_SIZE * 1.3)),
                        padding=[AppView.BASIC_SIZE / 2, AppView.BASIC_SIZE / 5])
        for i_sheet in range(n_sheet):
            # A new sheetColumns added to list to hold the tab frame
            cols_frame = SheetColumns(None, [])
            self.cols_tab.append(cols_frame)
            # default check state is False
            default_chk_state = False
            # to mark select all state change
            self.chk_all_last_state.append(default_chk_state)
            # add tab
            sheet_frame = Frame(sheets_tab)
            sheets_tab.add(sheet_frame, text=self.sheet_column_labels_list[i_sheet].sheet)
            # add current sheet frame name to SheetColumns.sheet
            cols_frame.sheet = self.sheet_column_labels_list[i_sheet].sheet
            # create tab content
            grid_col_num = 4
            # select all button
            chk_all_state = BooleanVar(sheet_frame)
            chk_all_state.set(default_chk_state)
            chk_all = Checkbutton(sheet_frame, text="全选", variable=chk_all_state, pady=AppView.BASIC_SIZE,
                                  font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)),
                                  command=self.__cols_check_all_selected)
            chk_all.grid(column=0, row=0, columnspan=grid_col_num, sticky=W)
            chk_all.var = chk_all_state
            # put select all button to cols_frame.columns[0],
            # so specific columns index starts with 1,
            # same as it at xlwings
            cols_frame.columns.append(chk_all)
            # get all cols in current sheet
            cols = self.sheet_column_labels_list[i_sheet].columns
            for i_col in range(len(cols)):
                chk_state = BooleanVar(sheet_frame)
                chk_state.set(default_chk_state)
                chk = Checkbutton(sheet_frame, text=cols[i_col], variable=chk_state,
                                  font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)))
                chk.grid(row=int(i_col // grid_col_num) + 1, column=i_col % grid_col_num,
                         padx=AppView.BASIC_SIZE, sticky=W)
                chk.var = chk_state
                # add column chk button to list
                cols_frame.columns.append(chk)
        # ------------------------ part two to filter excel rows -----------------
        rows_panel = PanedWindow(self.window)
        # layout rows_panel
        rows_panel.pack(side='right', ipadx=AppView.BASIC_SIZE * 2, anchor='ne', pady=AppView.BASIC_SIZE * 2)
        # label for row filter type
        row_filter_type_label = Label(rows_panel, text="选择筛选行的方式",
                                      font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        row_filter_type_label.grid(column=0, row=0, columnspan=5, sticky=W, pady=AppView.BASIC_SIZE)
        self.row_type_selected = IntVar(rows_panel)
        # row filter type
        rad_person = Radiobutton(rows_panel, text=RowFilterType.FILTER_PERSON_STRING,
                                 value=RowFilterType.FILTER_PERSON_VALUE,
                                 variable=self.row_type_selected, font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        # default choose person
        rad_person.select()
        rad_depart = Radiobutton(rows_panel, text=RowFilterType.FILTER_DEPART_STRING,
                                 value=RowFilterType.FILTER_DEPART_VALUE,
                                 variable=self.row_type_selected, font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        rad_proj = Radiobutton(rows_panel, text=RowFilterType.FILTER_PROJECT_STRING,
                               value=RowFilterType.FILTER_PROJECT_VALUE,
                               variable=self.row_type_selected, font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        rad_class = Radiobutton(rows_panel, text=RowFilterType.FILTER_CLASS_STRING,
                                value=RowFilterType.FILTER_CLASS_VALUE,
                                variable=self.row_type_selected, font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        rad_no = Radiobutton(rows_panel, text=RowFilterType.FILTER_NO_STRING,
                             value=RowFilterType.FILTER_NO_VALUE,
                             variable=self.row_type_selected, font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        rad_person.grid(column=0, row=1)
        rad_depart.grid(column=1, row=1)
        rad_proj.grid(column=2, row=1)
        rad_class.grid(column=3, row=1)
        rad_no.grid(column=4, row=1)
        # a blank label to separate row filter type and content
        row_filter_content_label = Label(rows_panel, text="  ",
                                         font=("微软雅黑", int(AppView.BASIC_SIZE * 2)))
        row_filter_content_label.grid(column=0, row=2, columnspan=5, pady=AppView.BASIC_SIZE * 3)
        # row filter content
        row_filter_content_label = Label(rows_panel, text="选择筛选的内容（多个项目之间用中文逗号分隔）",
                                         font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        row_filter_content_label.grid(column=0, row=3, columnspan=5, sticky=W, pady=AppView.BASIC_SIZE)
        # row filter content text area
        self.row_filter_content_area = scrolledtext.ScrolledText(rows_panel, width=AppView.BASIC_SIZE * 4,
                                                                 height=AppView.BASIC_SIZE,
                                                                 font=("微软雅黑", int(AppView.BASIC_SIZE * 1.6)))
        self.row_filter_content_area.grid(column=0, row=4, columnspan=5, sticky=W)
        # submit button
        self.submit_button = Button(self.window, text="确定", command=self.__submit_button_clicked,
                                    font=("微软雅黑", int(AppView.BASIC_SIZE * 1.2)),
                                    padx=AppView.BASIC_SIZE)
        # self.submit_button.config(state=DISABLED)
        # layout
        self.submit_button.place(relx=0.5, rely=0.92, anchor=CENTER)
        # ------------------------------------------------
        # show window
        self.window.mainloop()

    def __submit_button_clicked(self):
        # get cols selected
        selected_cols = []
        for i in range(len(self.cols_tab)):
            sht_sel_cols = SheetColumns(None, [])
            selected_cols.append(sht_sel_cols)
            # sheet name
            sht_sel_cols.sheet = self.sheet_column_labels_list[i].sheet
            # selected cols index
            cols = self.cols_tab[i]
            # cols.columns[0] is check_all, so idx starts with 1 here
            for j in range(1, len(cols.columns)):
                if cols.columns[j].var.get() is True:
                    # minus 1 to get correct index of cols
                    sht_sel_cols.columns.append(j - 1)
        # get row filter type
        row_filter_type = self.row_type_selected.get()
        # get row filter content
        row_filter_content = self.row_filter_content_area.get(1.0, END)
        # ----------- check valid ----------
        cols_empty = True
        for sht in selected_cols[:]:
            if len(sht.columns) != 0:
                cols_empty = False
            else:
                # remove sheet with no cols selected
                selected_cols.remove(sht)
        if cols_empty is True:
            messagebox.showinfo("错误", "请至少选出一列导出")
        elif len(row_filter_content) == 1:
            # if user input nothing, row_filter_content will be a '\n',
            # so the empty input length is 1
            messagebox.showinfo("错误", "请输入行筛选内容")
        else:
            try:
                # check pass, call function to export data
                self.do_export_data_function(selected_cols, row_filter_type, row_filter_content)
                # close this window
                self.window.destroy()
            except Exception as e:
                messagebox.showerror("error", "exception in do_export_data_function: " + repr(e))

    def __cols_check_all_selected(self):
        # do not care which check all clicked,
        # just update all button state
        for i in range(len(self.cols_tab)):
            cols_frame = self.cols_tab[i]
            # if status change
            if cols_frame.columns[0].var.get() is not self.chk_all_last_state[i]:
                if cols_frame.columns[0].var.get() is True:
                    # select all
                    for j in range(len(cols_frame.columns)):
                        cols_frame.columns[j].select()
                else:
                    # deselect all
                    for j in range(len(cols_frame.columns)):
                        cols_frame.columns[j].deselect()
                # update last state
                self.chk_all_last_state[i] = cols_frame.columns[0].var.get()

    def __reset_attributes(self):
        """ clear some self attributes. avoid opening this window multi times leading errors"""
        self.sheet_column_labels_list = None
        self.cols_tab = []
        self.chk_all_last_state = []
        self.row_type_selected = None

# -------------------test--------------------
# ex_view = ExportDataView()
# ex_view.draw([SheetColumns("sheet1中文", ["abc", "dsf", "ssss", "yyyy", "aa", "ff", "af", "wef", "adg"]),
#               SheetColumns("sh2", ["abc", "d阿道夫", "ssss", "yy", "阿斯顿", "一个", "中文", "aad"])])
