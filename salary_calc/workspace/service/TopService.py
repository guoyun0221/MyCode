import datetime
import traceback
from turtle import st
from typing import List

# from service.ExcelWorker import ExcelWorker
import service.ExcelWorker as excel_worker
from service.Model import Operation, SheetColumns, RowFilterType
from ui.ExportDataView import ExportDataView
import xlrd
import xlwt

class TopService:
    # def __init__(self, excel_worker: ExcelWorker, export_view: ExportDataView):
    def __init__(self, export_view: ExportDataView):
        self.op_type = None
        self.src_path = None
        self.dst_path = None
        # self.excel_worker = excel_worker
        self.export_view = export_view
        # for export excel. this operation need to read excel twice,
        # so cache the sheets to reduce one read
        self.export_src_sheets = None
        # flag mark it's beijing(0) or qingdao(1)
        self.location = -1

    def dispatch_operation(self, op_type: str, src_path: str, dst_path: str):
        # update self field
        self.op_type = op_type
        self.src_path = src_path
        self.dst_path = dst_path
        try:
            # check it's beijing or qingdao 
            # if "北京" in src_path:
            #     self.location = 0
            # elif "青岛" in src_path:
            #     self.location = 1
            # else:
            #     raise Exception("location needed in src path")
            # dispatch operation function
            if op_type == Operation.OP_TYPE_EXPORT_ITEMS:
                self.__export_data()
            elif op_type == Operation.OP_TYPE_PROJECT_COST:
                self.__generate_project_cost_sheet()
            else:
                print("other ops")
        except Exception as e:
            traceback.print_exc()
            raise Exception("exception happened in top service: " + repr(e))

    def __export_data(self):
        # get src columns label to show
        # sheets = self.excel_worker.read_excel(self.src_path)
        sheets = excel_worker.read_excel(self.src_path)
        # cache sheets
        self.export_src_sheets = sheets
        # get cols
        sht_cols_list = []
        for sht in sheets:
            if (sht.name in ["技术岗位", "其他类岗", "工资转青岛", "实习生+劳务"] and self.location == 1) \
                    or (sht.name in ["市场", "技术", "管理"] and self.location == 0):
                sht_cols = SheetColumns()
                sht_cols.sheet = sht.name
                sht_cols.columns = excel_worker.get_one_row_with_colname(sht, 1)
                sht_cols_list.append(sht_cols)
        # set export view callback function
        self.export_view.set_do_export_data_function(self.do_export_data)
        # show export view, wait user to select cols and rows filter
        self.export_view.draw(sht_cols_list)
        # export data

    def do_export_data(self, selected_cols: List[SheetColumns], row_filter_type: int, row_filter_content: str):
        """ exec export data from one excel to another. \n
         :argument selected_cols A SheetColumns list, to get which cols to export of each sheet
         :argument row_filter_type specific filter row by what column label
         :argument row_filter_content keep row with the content and drop others"""
        try:
            # check op_type not changed
            if self.op_type != Operation.OP_TYPE_EXPORT_ITEMS:
                raise Exception("operation type changed")

            # convert row_filter_content from str to str list
            row_filter_content_list = row_filter_content.split("，")
            for i in range(len(row_filter_content_list)):
                row_filter_content_list[i] = row_filter_content_list[i].strip().replace("\n", "")
            # delete empty elems
            def not_empty(s):
                return s and s.strip()
            filter_empty = filter(not_empty, row_filter_content_list)
            row_filter_content_list = list(filter_empty)
            # create dst workbook
            # get names
            sheet_names = []
            for sheet_cols in selected_cols:
                sheet_names.append(sheet_cols.sheet)
            create_sheet_result = excel_worker.create_workbook_with_sheet_names(sheet_names)
            dst_wb = create_sheet_result[0]
            dst_sheets = create_sheet_result[1]
            # read src excel
            # src_sheets = self.excel_worker.read_excel(self.src_path)
            # use cached src sheets
            src_sheets = self.export_src_sheets
            for i in range(len(selected_cols)):
                # get sheet index in src sheets
                sheet_name = selected_cols[i].sheet
                sheet_index = -1
                for j in range(len(src_sheets)):
                    if src_sheets[j].name == sheet_name:
                        sheet_index = j
                        break
                # filter rows
                if row_filter_type != RowFilterType.FILTER_NO_VALUE:
                    cols_index = excel_worker.filter_columns(src_sheets[sheet_index], 1,
                                                                  [RowFilterType.VALUE_TO_STRING[row_filter_type]])
                    select_rows = excel_worker.filter_rows(src_sheets[sheet_index], cols_index[0],
                                                                row_filter_content_list)
                    # copy excel header
                    select_rows.insert(0, 1)
                else:
                    select_rows = excel_worker.all_rows_indexes(src_sheets[sheet_index])
                # do copy data
                if len(select_rows) != 0:
                    excel_worker.export_data(src_sheets[sheet_index], dst_sheets[i],
                                                         select_rows, selected_cols[i].columns)

            # save workbook
            # handle dst file name
            t_now = datetime.datetime.now()
            t_now = t_now.strftime("%m%d%H%M")
            dst_file = self.dst_path + "/导出文件" + t_now + ".xls"
            excel_worker.save_workbook(dst_wb, dst_file)
        except Exception as e:
            traceback.print_exc()
            raise Exception("exception happened in top service:", repr(e))

    def __generate_project_cost_sheet(self):
        if self.op_type != Operation.OP_TYPE_PROJECT_COST:
            raise Exception("operation changed")
        # # create dst workbook and sheet
        create_dst_wb_result = excel_worker.create_workbook_with_sheet_names(['工资报会计表'])
        dst_wb = create_dst_wb_result[0]
        report_sheet = create_dst_wb_result[1][0]
        
        # get src sheets to operation
        src_sheet = []
        sheets = excel_worker.read_excel(self.src_path)
        # get copy workbook to write
        # old_workbook = xlrd.open_workbook(self.src_path)
        # new_workbook = copy(old_workbook)

        # if self.location == 0:
        #     src_sheet.append(excel_worker.get_sheet_by_name(sheets, '项目'))
        #     # report_sheet = new_workbook.get_sheet(excel_worker.get_sheet_index_by_name(src_sheet, '工资报会计表'))
        #     # original_report_sheet = excel_worker.get_sheet_by_name(sheets, '工资报会计表')
        #     # report_sheet = new_workbook.get_sheet('工资报会计表')
        #     # dst_sheet_index = excel_worker.get_sheet_index_by_name(sheets, "工资报会计表")
        #     # dst_nrows = excel_worker.get_sheet_by_name(sheets, '工资报会计表').nrows
        #     # dst_ncols = excel_worker.get_sheet_by_name(sheets, '工资报会计表').ncols
        # else:
        #     src_sheet.append(excel_worker.get_sheet_by_name(sheets, '项目'))
        #     src_sheet.append(excel_worker.get_sheet_by_name(sheets, '其他'))
        #     # report_sheet = new_workbook.get_sheet(excel_worker.get_sheet_index_by_name(src_sheet, '报会计2'))
        #     # original_report_sheet = excel_worker.get_sheet_by_name(sheets, '报会计2')
        #     # report_sheet = new_workbook.get_sheet('报会计2')
        #     # dst_sheet_index = excel_worker.get_sheet_index_by_name(sheets, "报会计2")
        #     # dst_nrows = excel_worker.get_sheet_by_name(sheets, '报会计2').nrows
        #     # dst_ncols = excel_worker.get_sheet_by_name(sheets, '报会计2').ncols

        src_sheet.append(excel_worker.get_sheet_by_name(sheets, 'Sheet1'))

        # a dict to save information.
        # the dict is - [project_name_str: project_info_list]
        # project_info_list is - List[int], every element in
        # the list stands for one column need to be sum.
        proj_info_map = dict()
        # get info sum from src_sheets
        for sheet in src_sheet:
            # get '项目' location column form header
            header_row = 1
            project_key_col = excel_worker.filter_columns(sheet, header_row, ['项目'])[0]
            # get last column and row index
            # last_col = project_key_col
            # while (sheet.range(header_row, last_col + 1).value is not None
            #        and sheet.range(header_row, last_col + 1).value != ""
            #        and sheet.range(header_row, last_col + 1).value != "None"):
            #     last_col += 1
            # last_row = sheet.used_range.last_cell.row
            nrows = sheet.nrows
            ncols = sheet.ncols
            # mark if it's in project row
            flag = False
            for i in range(1, nrows):
                # sheet_value = sheet.range(i, project_key_col).value
                sheet_value = sheet.cell_value(i, project_key_col)
                # some values are parsed to float lise 1901.0,
                # so convert value to int first and then convert to str
                if sheet_value is not None and type(sheet_value) is not str:
                    sheet_value = str(int(sheet_value))
                # it's header row, start projects rows
                if sheet_value == '项目':
                    # set flag and process next row
                    flag = True
                    continue
                # it's the end of project rows
                if sheet_value is None or sheet_value == '' \
                        or sheet_value == "None" or sheet_value == '合计':
                    # set flag and process next row
                    flag = False
                    continue
                # process project row
                if flag:
                    # add project info to dict
                    if sheet_value not in proj_info_map:
                        # init, info_list[0] is count
                        info_list = [0]
                        for j in range(project_key_col + 1, ncols):
                            info_list.append(0)
                        # add to dict
                        proj_info_map[sheet_value] = info_list
                    # write info to dict
                    proj_info = proj_info_map[sheet_value]
                    # proj_info[0] is count
                    proj_info[0] += 1
                    for k in range(1, len(proj_info)):
                        # handle blank cells
                        if sheet.cell_value(i, project_key_col + k) is not None:
                            current_sheet_value = sheet.cell_value(i, project_key_col + k)
                            if current_sheet_value == '':
                                current_sheet_value = 0
                            proj_info[k] = proj_info[k] + current_sheet_value
                    proj_info_map[sheet_value] = proj_info
        # write info to dst_sheet
        # nrows = dst_nrows
        # ncols = dst_ncols

        # the format is expected, okay to clear stuff and rewrite some
        # -> too much style and format stuff to set, do not want to do that then
        # if dst_sheet.range(3, 1).value == "序号" and dst_sheet.range(3, 2).value == "项目" \
        #         and max_col == 17 and dst_sheet.range(3, 17).value == "工资实发":
        #     dst_sheet.range(f'A4:Q{max_row}').clear()
        #     dst_sheet.range(f'A4:Q{max_row}').color = (255, 255, 255)
        #     for i in range(len(proj_info_map)):
        #         dst_sheet.range(4 + i, 1).value = i + 1
        #     sum_row = 4 + len(proj_info_map)
        #     dst_sheet.range(sum_row, 2).value = "合计"
        #     tail_row = sum_row + 1
        #     dst_sheet.range(tail_row, 2).value = "制表："
        #     dst_sheet.range(tail_row, 3).value = "赵丹"
        #     dst_sheet.range(tail_row, 8).value = "领导审核"
        # else:
        #     print("not expected")

        # hearder row
        if self.location == 0:
            headers = ['序号','项目','人数','应发合计','扣养老保险','扣失业保',',扣医疗保险','扣住房公积金','扣保险小计',
            '补缴保险','补缴公积金','预扣保险费','保密补贴（2022）','应发工资','个人所得税','工资实发']
            title = '2022年 月中科本原科技（北京）有限公司'
        else: 
            headers = ['序号','项目','人数','应发合计','扣养老保险','扣失业保',',扣医疗保险','扣住房公积金','扣保险小计',
            '补缴保险','补缴公积金','预扣保险费','保密补贴（2022）','应发工资','个人所得税','工会会费','工资实发']
            title = '2022年 月青岛本原微电子有限公司工资表---报会计'
        ncols = len(headers)

        # write title
        style_t = xlwt.XFStyle()
        al = xlwt.Alignment()
        al.horz = 0x02
        al.vert = 0x01
        style_t.alignment = al

        report_sheet.write_merge(0, 0, 0, ncols - 1, title, style=style_t)

        # write header row
        for j in range(ncols):
            report_sheet.write(1, j, headers[j])

        # write number col
        for j in range(len(proj_info_map)):
            report_sheet.write(j + 2, 0, j + 1)

        # write data
        excel_worker.write_dict_of_list_to_sheet(report_sheet, proj_info_map, 2, 1)
        
        # save
        t_now = datetime.datetime.now()
        t_now = t_now.strftime("%m%d%H%M")
        dst_file = self.dst_path + "/报会计表" + t_now + ".xls"
        dst_wb.save(dst_file)
