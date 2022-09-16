import traceback
from typing import List, Dict
import xlwings as xw


class ExcelWorker:

    def __init__(self):
        # instantiate an Excel app
        self.app = None
        # all opened workbooks
        self.workbooks = None
        # init state
        self.init_state = False

    def init(self):
        # instantiate an Excel app
        self.app = xw.App(visible=False, add_book=False)
        # set workbooks type to list
        self.workbooks = []
        # set inited flag
        self.init_state = True

    def read_excel(self, file_path: str) -> xw.main.Sheets:
        """read excel with given file_path, return the sheets list"""
        # open workbook
        workbook = self.get_workbook(file_path)
        # return all sheets
        sheets = workbook.sheets
        return sheets

    def get_workbook(self, file_path: str = None) -> xw.main.Book:
        """ Open a workbook with given filename, \n
        or create a new one if the filename is None"""
        if file_path is None:
            workbook = self.app.books.add()
        else:
            workbook = self.app.books.open(file_path)
        # add workbook to list if not in it
        if workbook not in self.workbooks:
            self.workbooks.append(workbook)
        # return the workbook
        return workbook

    def save_workbooks(self):
        """ saved opened workbooks """
        for workbook in self.workbooks:
            workbook.save()

    def close_all(self):
        """ Close and release resources. \n
        This function should always be called after processing excel"""
        # close workbooks
        for workbook in self.workbooks:
            workbook.save()
            workbook.close()
        # quit Excel app
        self.app.quit()

    def get_one_row_with_column_name(self, sheet: xw.main.Sheet, row_index: int) -> List[str]:
        """get one row of sheet. \n
        returns a list for the row selected, blank cells dropped\n
        the element format will be like: "colname: cell_value" \n """
        self.__check_okay()
        # get sheet cols number
        cols_num = sheet.used_range.last_cell.column
        result = []
        for i in range(1, cols_num + 1):
            # get cell
            cell = sheet.range((row_index, i))
            # get value and check not empty
            value = cell.value
            if (value is not None) and (value != ""):
                # get address and leave only column name
                address = cell.get_address()
                # $A$1 -> A
                address = address.split("$")[1]
                result.append(address + ": " + value)
        return result

    def copy_selected_data(self, src_sheet: xw.main.Sheet, dst_sheet: xw.main.Sheet, rows: List[int], cols: List[int]):
        """export data from original Excel sheet (src_sheet) to a new one (dst_sheet)\n
        argument rows and cols is to specify which cells will be exported"""
        self.__check_okay()
        for i in range(len(rows)):
            for j in range(len(cols)):
                # get original cell value
                value = src_sheet.range(rows[i], cols[j]).value
                # print("rows[i]: ", rows[i], " cols[j]: ", cols[j], " value: ", value)
                # write to dst sheet
                # xlwings index starts with 1
                dst_sheet.range(i + 1, j + 1).value = value

    def create_workbook_with_sheet_names(self, file_path: str, sheet_names: List[str]) -> xw.main.Sheets:
        """ create a new workbook with given file_path. \n
        and add a list of sheets with given sheet_names \n
        return sheets list created"""
        # create workbook
        workbook = self.get_workbook()
        # create first sheet
        workbook.sheets.add(sheet_names[0])
        for i in range(1, len(sheet_names)):
            # add sheet to workbook
            workbook.sheets.add(sheet_names[i], after=sheet_names[i - 1])
        # delete default created sheet1
        workbook.sheets["sheet1"].delete()
        # save the workbook created
        workbook.save(file_path)
        return workbook.sheets

    def filter_columns(self, sheet: xw.main.Sheet, key_row: int, select_content_list: List[str]) -> List[int]:
        """ For a specific row(ket_row), scan all columns, select column with whose value is in select_content_list.\n
        returns a list of selected columns indexes"""
        self.__check_okay()
        cell_values = sheet.range(f"{key_row}:{key_row}").value
        result = []
        for i in range(len(cell_values)):
            if cell_values[i] in select_content_list:
                # xlwings index starts with 1
                result.append(i + 1)
        return result

    def filter_rows(self, sheet: xw.main.Sheet, key_column: int, select_content_list: List[str]) -> List[int]:
        """ For a specific column(ket_column), scan all rows, select rows with whose value is in select_content_list.\n
        returns a list of selected rows indexes"""
        self.__check_okay()
        # convert number to column letter title
        col = self.__number_to_column_title(key_column)
        cell_values = sheet.range(col + ":" + col)[:sheet.used_range.last_cell.row].value
        result = []
        for i in range(len(cell_values)):
            if cell_values[i] in select_content_list:
                # xlwings index starts with 1
                result.append(i + 1)
        return result

    def all_rows_indexes(self, sheet: xw.main.Sheet) -> List[int]:
        """ get a list of specific sheet rows indexes """
        self.__check_okay()
        max_row = sheet.used_range.last_cell.row
        result = []
        for i in range(1, max_row):
            result.append(i)
        return result

    def write_dict_of_list_to_sheet(self, sheet: xw.main.Sheet, list_dict: Dict[str, List],
                                    start_row: int, start_col: int):
        """ write a dict to sheet. the value type of the dict is list.\n
         every element in the dict takes one row. the key will be written first, \n
         then the list follows to right in order. the dict will be sort by key before write \n
         the first key of the dict will be written to (start_row, start_col)"""
        self.__check_okay()
        order_key = sorted(list_dict.keys())
        for i in range(len(order_key)):
            # write key
            sheet.range(start_row + i, start_col).value = order_key[i]
            # write value
            value_list = list_dict[order_key[i]]
            sheet.range(start_row + i, start_col + 1).value = value_list

    def __check_okay(self):
        if self.init_state is False:
            raise Exception("Excel worker haven't init yet")
        if len(self.workbooks) == 0:
            raise Exception("No workbooks found")

    def __number_to_column_title(self, column_number: int) -> str:
        """convert number index to excel column letter title"""
        self.__check_okay()
        # reference: leetcode 168
        ans = list()
        while column_number > 0:
            column_number -= 1
            ans.append(chr(column_number % 26 + ord("A")))
            column_number //= 26
        return "".join(ans[::-1])


# -------------- test -------------
worker = ExcelWorker()
try:
    worker.init()
    dst_sheets = worker.create_workbook_with_sheet_names(r'C:\Users\guoyu\Desktop\test\dst.xlsx',
                                                         ["aaa", "bbb", "cc"])
    # worker.copy_selected_data(src_sheets[0], dst_sheets[0], [5, 6, 7], [2, 7, 8, 9])
    # ret = worker.filter_columns(src_sheets[0], 2, ["id"])
    # print(ret)
    l1 = [12, 34, 56, 78]
    l2 = [78, 56, 34, 12]
    l3 = [11, 11, 11, 11]
    d1 = {"list2": l2, "list1": l1, "list3": l3}
    worker.write_dict_of_list_to_sheet(dst_sheets[0], d1, 2, 3)
except Exception as e:
    print("exception happened in excel worker: ", repr(e))
    traceback.print_exc()
finally:
    worker.close_all()
