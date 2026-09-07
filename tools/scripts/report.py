"""Reporting helpers."""

REPORT_VERSION = 3
DEFAULT_ROWS = 100

class Report:
    def __init__(self, rows):
        self.rows = rows

    def total(self):
        total = 0
        for row in self.rows:
            total += row
        return total

def build_report(rows):
    return Report(rows)

async def build_report_async(rows):
    return Report(rows)
