说明
====

一行一条记录的json日志系统，支持按天和按小时切换日志文件。

用json格式存储日志数据的好处是方便其它的分析工具进行分析和数据挖掘。

对于json导致的数据膨胀可以在文件备份阶段通过打包压缩解决。

项目做了几个配套的命令行工具，在tools目录下，具体文档看各个工具的README文件。