# 告警记录写入Elasticsearch

**注意**:

- 支持 ES 7.x 和 8.x 版本。（ES 从 v7 到 v8 有一些变化，但主要的 API 设计上是保持兼容的。）
- 索引根据年月动态创建 `prometheusalert-YYMM`(如prometheusalert-202112)
- 支持 ES HTTPS。

<br/>
<br/>

## es相关配置

```conf
# 是否将告警记录写入es7，0为关闭，1为开启
alert_to_es=0

# es地址，是[]string，beego.AppConfig.Strings读取配置为[]string，使用";"而不是","
# 单个地址
to_es_url=http://localhost:9200
# 多个地址
# to_es_url=http://es1:9200;http://es2:9200;http://es3:9200
# 多个 https 地址，程序里会跳过 https 证书检查
# to_es_url=https://es1:9200;https://es2:9200;https://es3:9200

# 是否有认证，es用户和密码, 无认证则不需要填写。
# to_es_user=username
# to_es_pwd=password
```

<br/>
<br/>

## Kibana展示效果

可直接创建索引模式展示告警记录。还可以自行创建表格、柱状图等展示。

![kibana-index](../images/kibana-index.png)

<br/>

![kibana-table](../images/kibana-table.png)
