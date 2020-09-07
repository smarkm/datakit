package inputs

var (
	AllInputs = map[string]bool{
		"activemq":               true,
		"aliyunactiontrail":      false,
		"aliyuncdn":              true, // bad doc
		"aliyuncms":              true,
		"aliyuncost":             false,
		"aliyunddos":             false,
		"aliyunfc":               false,
		"aliyunlog":              false,
		"aliyunobject":           true,
		"aliyunprice":            true,
		"aliyunrdsslowlog":       false,
		"aliyunsecurity":         true,
		"amqp_consumer":          false,
		"apache":                 true,
		"aws_billing":            false,
		"azure_monitor":          false,
		"baiduIndex":             true,
		"binlog":                 true,
		"cassandra":              true,
		"ceph":                   true,
		"clickhouse":             true,
		"cloudflare":             true,
		"cloudwatch":             false,
		"collectd":               true,
		"confluence":             true,
		"consul":                 true,
		"containerd":             true,
		"coredns":                true,
		"cpu":                    true,
		"csv":                    false,
		"dataclean":              false,
		"disk":                   true,
		"diskio":                 false,
		"dns_query":              true,
		"docker":                 true,
		"docker_log":             true,
		"druid":                  true,
		"elasticsearch":          true,
		"envoy":                  true,
		"etcd":                   true,
		"exec":                   true,
		"expressjs":              false, // TODO: impl by http-input
		"external":               false,
		"flink":                  true,
		"fluentd":                true,
		"github":                 true,
		"gitlab":                 false,
		"goruntime":              true,
		"hadoop_hdfs":            true,
		"haproxy":                false,
		"harborMonitor":          false, // bad doc
		"hostobject":             false,
		"http":                   false,
		"http_response":          false, // doc missing
		"httpjson":               false, // doc missing
		"httpstat":               true,
		"influxdb":               false,
		"iptables":               true,
		"jenkins":                true,
		"jira":                   true,
		"jolokia2_agent":         true,
		"jvm":                    true,
		"kafka":                  true,
		"kafka_consumer":         true,
		"kapacitor":              true,
		"kernel":                 true,
		"kibana":                 true,
		"kong":                   false, // TODO: impl by prom
		"kube_inventory":         false,
		"kubernetes":             false,
		"lighttpd":               true,
		"mem":                    true,
		"memcached":              true,
		"mock":                   false,
		"modbus":                 false,
		"mongodb":                true,
		"mongodb_oplog":          true,
		"mqtt_consumer":          true,
		"mysql":                  false,
		"mysqlMonitor":           false, // bad doc
		"nats":                   false,
		"neo4j":                  false, // TODO: impl by prom
		"net":                    false, // bad doc
		"net_response":           true,
		"netstat":                false, // doc missing
		"nfsstat":                true,
		"nginx":                  false,
		"nsq":                    true,
		"nsq_consumer":           true,
		"ntpq":                   false,
		"nvidia_smi":             true,
		"openldap":               true,
		"openntpd":               true,
		"oraclemonitor":          true,
		"phpfpm":                 true,
		"ping":                   true,
		"postgresql":             true,
		"postgresql_replication": true,
		"processes":              false,
		"procstat":               true,
		"prom":                   false,
		"puppetagent":            true,
		"rabbitmq":               true,
		"redis":                  false,
		"scanport":               false,
		"self":                   false,
		"snmp":                   true,
		"socket_listener":        false,
		"solr":                   true,
		"sqlserver":              false,
		"squid":                  false,
		"ssh":                    true,
		"statsd":                 true,
		"swap":                   true,
		"syslog":                 true,
		"system":                 true,
		"systemd":                true,
		"systemd_units":          false,
		"tailf":                  false,
		"tcpdump":                false,
		"tencentcms":             true,
		"tencentcost":            true,
		"tencentobject":          false, // doc miss
		"tengine":                true,
		"timezone":               true,
		"traceJaeger":            false,
		"traceSkywalking":        false,
		"traceZipkin":            false,
		"tracerouter":            true,
		"traefik":                true,
		"ucloud_monitor":         false,
		"uwsgi":                  false,
		"varnish":                false,
		"vsphere":                false,
		"weblogic":               true,
		"wechatminiprogram":      false, // bad doc
		"win_services":           true,
		"x509_cert":              true,
		"yarn":                   true,
		"zabbix":                 false,
		"zookeeper":              false,
	}
)
