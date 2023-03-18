from typing import List


class Operation:
    # supported op types
    OP_TYPE_PROJECT_COST = '生成项目核算表格'
    OP_TYPE_EXPORT_ITEMS = '导出数据至新表格'
    
    # OP_TYPE_EXPORT_SALARY_INFO = '导出工资条'
    # OP_TYPE_IMPORT_INFO = '根据外部信息生成表格数据'

    # put add support op types in a list
    # OP_TYPE_LIST = [OP_TYPE_EXPORT_ITEMS, OP_TYPE_PROJECT_COST,
    #                 OP_TYPE_EXPORT_SALARY_INFO, OP_TYPE_IMPORT_INFO]
    OP_TYPE_LIST = [OP_TYPE_PROJECT_COST, OP_TYPE_EXPORT_ITEMS]


class RowFilterType:
    FILTER_PERSON_STRING = "姓名"
    FILTER_PERSON_VALUE = 1
    FILTER_DEPART_STRING = "部门"
    FILTER_DEPART_VALUE = 2
    FILTER_PROJECT_STRING = "项目"
    FILTER_PROJECT_VALUE = 3
    FILTER_CLASS_STRING = "类别"
    FILTER_CLASS_VALUE = 4
    FILTER_NO_STRING = "全选"
    FILTER_NO_VALUE = 5

    VALUE_TO_STRING = {FILTER_PERSON_VALUE: FILTER_PERSON_STRING,
                       FILTER_DEPART_VALUE: FILTER_DEPART_STRING,
                       FILTER_PROJECT_VALUE: FILTER_PROJECT_STRING,
                       FILTER_CLASS_VALUE: FILTER_CLASS_STRING,
                       FILTER_NO_VALUE: FILTER_NO_STRING}

class SheetColumns:
    def __init__(self, sheet=None, columns: List = None):
        """ A sheet object with a list of column object """
        self.sheet = sheet
        self.columns = columns
