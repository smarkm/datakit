package csv

var (
	configSample = `
#[[inputs.csv]]
#	path_env = ''    # 有效的 python 安装路径
#	file = "xxx.xlsx"
#	start_rows = 10  # 有效数据开始行编号, 编号从 0 开始
#	metric = "test"
#	[[inputs.csv.columns]]
#		index = 0
#		name = "t1"
#		na_action = "ignore"  # ignore/drop/abort
#		as_tag = true
#	[[inputs.csv.columns]]
#		index = 1
#		name = "f1"
#		na_action = "ignore"
#		type = "int"         # int/str/float
#		as_field = true
#	[[inputs.csv.columns]]
#		index = 2
#		name = "f2"
#		na_action = "drop"
#		type = "float"
#		as_field = true
#	[[inputs.csv.columns]]
#		index = 3
#		name = "f3"
#		na_action = "abort"
#		type = "str"
#		as_field = true
#	[[inputs.csv.columns]]
#		index = 4
#		name = ""
#		na_action = "ignore"
#		as_time = true
#		time_format = "15/08/27 10:20:06" # csv/excel 中时间格式
#		time_precision = "s"              # 时间戳单位: h/m/s/ms/us/ns
`)