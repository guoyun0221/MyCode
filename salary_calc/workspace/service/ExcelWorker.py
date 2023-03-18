from typing import Dict, Iterable, List, Tuple
import xlrd
import xlwt


def read_excel(filename: str) -> List[xlrd.sheet.Sheet]:
    """read excel with given name, return the sheets list"""
    # open excel
    xl = xlrd.open_workbook(filename)
    # return all sheets
    return xl.sheets()


def get_one_row_with_colname(sheet: xlrd.sheet.Sheet, row_index: int) -> List[str]:
    """get one row of sheet. \n
    returns a list for the row selected \n
    the element format will be like: "colname: cell_value" \n """
    # get sheet cols number
    n_cols = sheet.ncols
    ret = []
    for i in range(n_cols):
        ret.append(xlrd.colname(i) + ":" + sheet.cell_value(row_index, i))
    return ret


def export_data(src_sheet: xlrd.sheet.Sheet, dst_sheet:  xlwt.Worksheet, rows: List[int], cols: List[int]):
    """export data from original Excel sheet (src_sheet) to a new one (dst_sheet)\n
    argument rows and cols is to specify which cells will be exported"""
    for i in range(len(rows)):
        for j in range(len(cols)):
            # get original cell value
            value = src_sheet.cell_value(rows[i], cols[j])
            # print("rows[i]: ", rows[i], " cols[j]: ", cols[j], " value: ", value)
            # write to dst sheet
            dst_sheet.write(i, j, value)


def create_workbook_with_sheet_names(sheet_names: List[str]) -> Tuple[xlwt.Workbook, List[xlwt.Worksheet]]:
    # create workbook
    workbook = xlwt.Workbook(encoding="UTF-8")
    # sheet list
    sheets = []
    for sheet_name in sheet_names:
        # add sheet to workbook
        sheet = workbook.add_sheet(sheet_name)
        # add sheet to return sheet list
        sheets.append(sheet)
    return workbook, sheets

def save_workbook(workbook:xlwt.Workbook, filename: str):
    workbook.save(filename)

def filter_columns(sheet: xlrd.sheet.Sheet, key_row: int, select_content_list: List[str]) -> List[int]:
    """ For a specific row(ket_row), scan all columns, select column with whose value is in select_content_list.\n
        returns a list of selected columns indexes"""
    res = []
    cell_values = sheet.row_values(key_row)
    for i in range(len(cell_values)):
        if cell_values[i] in select_content_list:
            res.append(i)
    return res

def filter_rows(sheet: xlrd.sheet.Sheet, key_col: int, select_content_list: List[str]) -> List[int]:
    """ For a specific column(ket_column), scan all rows, select rows with whose value is in select_content_list.\n
        returns a list of selected rows indexes"""
    res = []
    cell_values = sheet.col_values(key_col)
    for i in range(len(cell_values)):
        if cell_values[i] in select_content_list:
            res.append(i)
    return res

def all_rows_indexes(sheet: xlrd.sheet.Sheet) -> List[int]:
    nrows = sheet.nrows
    result = []
    for i in range(0, nrows):
        result.append(i)
    return result

def write_dict_of_list_to_sheet(sheet: xlwt.Worksheet,list_dict: Dict[str, List], start_row: int, start_col: int):
    """ write a dict to sheet. the value type of the dict is list.\n
    every element in the dict takes one row. the key will be written first, \n
    then the list follows to right in order. the dict will be sort by key before write \n
    the first key of the dict will be written to (start_row, start_col)"""
    order_key = sorted(list_dict.keys())
    for i in range(len(order_key)):
        # write key
        sheet.write(start_row + i, start_col, order_key[i])
        # write value
        value_list = list_dict[order_key[i]]
        for j in range(len(value_list)):
            sheet.write(start_row + i, start_col + j + 1, value_list[j])

def get_sheet_by_name(sheets: List[xlrd.sheet.Sheet], name: str) -> xlrd.sheet.Sheet:
    for i in range(len(sheets)):
        if sheets[i].name == name:
            return sheets[i]

def get_sheet_index_by_name(sheets: List[xlrd.sheet.Sheet], name: str) -> xlrd.sheet.Sheet:
    for i in range(len(sheets)):
        if sheets[i].name == name:
            return i
# -------------- test -------------
# sheets = read_excel(r'C:\Users\guoyu\Desktop\工资核算2022（青岛）.xlsx')
# print(sheets[0].name)
# filter_columns(sheets[0], 1, [])
# create_ret = create_workbook_with_sheet_names(["aaa", "bbb"])
# wb = create_ret[0]
# dst_sheets = create_ret[1]
# export_data(sheets[0], dst_sheets[0], [5,6,7], [1,7,8,9])
# save_workbook(wb, r'C:\Users\guoyu\Desktop\salary_calc\workspace\excel\dst.xls')