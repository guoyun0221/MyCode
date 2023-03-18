import traceback

from service.TopService import TopService
from ui.AppView import AppView
from ui.ExportDataView import ExportDataView


class ExcelHelper:

    def __init__(self):
        self.top_service = None
        self.app_view = None
        self.export_view = None
        # self.excel_worker = None

    def start(self):
        try:
            # init field
            # self.excel_worker = ExcelWorker()
            self.export_view = ExportDataView()
            # self.top_service = TopService(self.excel_worker, self.export_view)
            self.top_service = TopService(self.export_view)
            self.app_view = AppView(self.top_service.dispatch_operation)
            # init excel worker
            # self.excel_worker.init()
            # show app window
            self.app_view.draw()
        except Exception as e:
            print("Exception happened in ExcelHelper: ", str(e))
            traceback.print_exc()
        # finally:
        #     self.excel_worker.close_all()
