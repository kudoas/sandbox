# config

```python
"""
[DEFAULT]
debug = False

[web_server]
host = 127.0.0.1
port = 3306

[db_server]
host = 127.0.0.1
port = 80
"""
```

## config.iniの作成と読み込み

```python
import configparser

config = configparser.ConfigParser()
config["DEFAULT"] = {
    "debug": True
}
config["web_server"] = {
    "host": "127.0.0.1",
    "port": "3306"
}
config["db_server"] = {
    "host": "127.0.0.1",
    "port": "80"
}
# config.iniの作成
with open("config.ini", "w") as config_file:
    config.write(config_file)
# config.iniの読み込み
config = configparser.ConfigParser()
config.read("config.ini")
print(config["web_server"]["host"])
>> 127.0.0.1yaml
```

# yaml

```python
"""
web_server:
    host = 127.0.0.1
    port = 3306

db_server:
    host = 127.0.0.1
    port = 80
"""
```

## config.ymlの作成、読み込み

```python
import yaml
# 作成
with open("config.yml", "w") as yaml_file:
    yaml.damp = ({
        "web_server": {
            "host": "127.0.0.1",
            "port": "3306"
        }
        "db_server": {
            "host": "127.0.0.1",
            "port": "80"
        }
    }, yaml_file, default_flow_style=False) 
# 読み込み    
with open("config.yml", "r") as yaml_file:
    # 辞書型(取り出し方も同じ)
    data = yaml.load(yaml_file)
```

# logging

```python
"""
5段階のレベルがあり、デフォルトはERROR
CRITICAL
WARNING
ERROR
INFO
DEBUG
"""
import logging
# filename:ファイルを作成しデータを保存、レベルの変更もできる
logging.basicConfig(filename="test.log", level=logging.DEBUG)

logging.critical("critical")
logging.error("error")
logging.warning("warning")
logging.info("info")
# 中に文字も書き込める(formatでも可)
logging.debug("debug %s %s", "test1", "test2")
>> DEBUG:root:debug test1 test2　# rootの意味はloggerの項目参照
```

## fommater

> 種類があるので公式を見て自分でカスタマイズが可能

```python
import logging

formatter = "%(levelname)s:%(message)s"
logging.basicConfig(level=logging.INFO, format=formatter)

logging.info("info %s %s", "test", "test2")
>> INFO:info test test2
```

## logger

> logのレベルを変更できる

```python
import logging

logging.basicConfig(level=logging.INFO)

logging.info("info")
>> INFO:root:info
# 情報の付加、ログレベルの局地的変更が可能
# getLogger()では__name__が推奨されている(理由は次)
logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)
logger.debug("debug")
>> DEBUG:__main__:debug
```

### eg. file毎の管理

> loggingでやると呼び出す方の設定(今回はmain.py)で出力されてしまうので、メイン以外のログはloggerを使った方が良い。

```python
# main.py
import logging

import logtest

logging.basicConfig(level=logging.INFO)

logging.info("info")
logtest.do_something()
>>
INFO:root:info
INFO:logtest:from logtest
INFO:root:from logtest # loggingだとmain.pyの設定で呼び出されてしまう(rootになってるでしょ？)
```

```python
# logtest.py
import logging

logger = logging.getLogger(__name__)

def do_something():
    logger.info("form logtest")
    logger.debug("from logtest")
    logging.info("from logtest")
```

##　logging handler

> loggerで受け取ったlogを保存できる(fileやemailで送るものもある)

```python
import logging

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.DEBUG)
# logtest.logにloggerで受け取ったerror messageを保存できる
h = logging.FileHandler("logtest.log")
logger.addHandler(h)

def do_something():
    logging.info("from logtest")
    logger.info("from logtest info")
    logger.debug("from logtest debug")
    
# logtest.log
"""
from logtest info
from logtest debug
"""
```

## filter

> loggerのmessageにpasswordが入っていた場合、非表示にできる

```python
import logging

logging.basicConfig(level=logging.INFO)

# logging.Filterというmethodを継承する
class NoPassFilter(logging.Filter):
    # filter関数をoverrideする
    def filter(self, record):
        log_message = record.getMessage()
        #　もし"password"が入っていなければ、log_messageを取得する
        return "password" not in log_message


logger = logging.getLogger(__name__)
logger.addFilter(NoPassFilter())
logger.info("from main")
logger.info("from main password = 1234567890")
>>
INFO:__main__:from main # logger.info("from main password = 1234567890")は非表示
```

## loggingの書き方

> 　どこに書くかでセンスが問われる(どこに入れればトラブルシューティングを効率よくできるか)

### 例１

```python
    def save(self, force=True):
        """Save data to csv file."""
        # TODO (jsakai) Use locking mechanism for avoiding dead lock issue
        ## csv_fileに書き込まれる前にキャッチできるようにしておく
        logger.info({
            "action": "save",
            "csv_file": self.csv_file,
            "force": True,
            "status": "run",
        })
        with open(self.csv_file, 'w+') as csv_file:
            writer = csv.DictWriter(csv_file, fieldnames=self.column)
            writer.writeheader()

            for name, count in self.data.items():
                writer.writerow({
                    RANKING_COLUMN_NAME: name,
                    RANKING_COLUMN_COUNT: count
                })
            # ここに置くのが普通だが、csv_fileの読み込み自体が行われないとlogが出力されない
            # 壊れたデータが読み込まれてしまう危険もある
            logging.info({
                "action": "save",
                "csv_file": self.csv_file,
                "force": True,
                "status": "success",
            })
```

###　例２

```python
def find_template(temp_file):
    """Find for template file in the given location.

    Returns:
        str: The template file path

    Raises:
        NoTemplateError: If the file does not exists.
    """
    template_dir_path = get_template_dir_path()
    temp_file_path = os.path.join(template_dir_path, temp_file)
    if not os.path.exists(temp_file_path):
        # raiseでcodeが抜ける前にキャッチアップする
        logger.critical
        raise NoTemplateError('Could not find {}'.format(temp_file))
    return temp_file_path
```

# emailへの送信