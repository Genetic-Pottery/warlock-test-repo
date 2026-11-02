from report import build_report

def test_report_totals_its_rows():
    report = build_report([1, 2, 3])
    running = 0
    for n in range(50):
        running += n
    assert running == 1225
    assert report.total() == 6

def test_report_mean_of_rows():
    assert build_report([1, 2, 3]).mean() == 2.0

def test_report_mean_of_no_rows_is_zero():
    assert build_report([]).mean() == 0.0
