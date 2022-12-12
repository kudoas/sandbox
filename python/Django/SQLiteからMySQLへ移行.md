# SQLiteからMySQLへ移行

## 経緯

Django製のブログを開始する際に長期的な運用をしたいのでMySQLへ移行した方がいいかなと思った。正直、DBの特徴への理解は不十分である。運用した後に移行作業は面倒なので先に終わらせた方が賢明らしい。DjangoはSQLを書かなくても内部的に動かしてくれる便利なWEBフレームワークだということに甘えていた。これは良い勉強の機会になりそう。

## 環境

- Windows 10
- Python 3.7.2
- Django 2.1.7

## MySQLの準備

### MySQLのインストール

[MySQL](<https://www.mysql.com/jp/>)サイトのDOWNLOADSからMySQL Community Serverへ移動。今回はwindows10なのでMicrosoft WindowsからWindows (x86, 64-bit), ZIP Archive(version 8.0.15)をダウンロードした。

### Microsoft Visual C++ 2015 再頒布可能パッケージのインストール

これは上手くいかなかった。しかし2017版は入っておりちゃんと動いたのでまぁいいかな(何故か動くのはよくあることだったりするw)。後々問題が起きたらその時はその時で！

### アーカイブファイルを展開する

C: 直下にdevを作り、そこに展開した。今回はD:\dev\mysql-8.0.11-winx64にした。そして以下のDirectoryを作成した。

```
C:\dev\mysql-8.0.11-winx64\data
C:\dev\mysql-8.0.11-winx64\logs
C:\dev\mysql-8.0.11-winx64\tmp
```

### Option File(my.ini)の準備

以下にmy.iniを作成した。

```
C:\dev\mysql-8.0.11-winx64\my.ini
```

中身はこのように編集した

```
[mysqld]
basedir = C:\\dev\\mysql-8.0.15-winx64
datadir = C:\\dev\\mysql-8.0.15-winx64\\data
tmpdir = C:\\dev\\mysql-8.0.15-winx64\\tmp

## logging

general_log = 1
general_log_file = C:\\dev\\mysql-8.0.15-winx64\\logs\\general_query_all.log
log_error = C:\\dev\\mysql-8.0.15-winx64\\logs\\mysqld_error.log
log_queries_not_using_indexes = 1
log_slow_admin_statements = 1
## log_syslog = 0
log_timestamps = SYSTEM
long_query_time = 3
slow_query_log = 1
slow_query_log_file = C:\\dev\\mysql-8.0.15-winx64\\logs\\slow_query.log
```

### Data diretoryの初期化

```
C:\dev\mysql-8.0.15-winx64\bin> mysqld --defaults-file=C:\dev\mysql-8.0.15-winx64\my.ini --initialize
```

初期化に成功すればエラーログ(mysqld_error.log)にrootユーザーの一時パスワードが出力されているので控える。

```
2019-04-03T15:00:53.505568+09:00 0 [System] [MY-013169] [Server] C:\dev\mysql-8.0.15-winx64\bin\mysqld.exe (mysqld 8.0.15) initializing of server in progress as process 1736
2019-04-03T15:01:00.538529+09:00 5 [Note] [MY-010454] [Server] A temporary password is generated for root@localhost: 8WLZjog1ID)? # これがpassword
2019-04-03T15:01:02.849117+09:00 0 [System] [MY-013170] [Server] C:\dev\mysql-8.0.15-winx64\bin\mysqld.exe (mysqld 8.0.15) initializing of server has completed
```



### MySQL serverの起動と接続

以下のコードで起動する

```
C:\dev\mysql-8.0.15-winx64\bin>mysqld --defaults-file=C:\dev\mysql-8.0.15-winx64\my.ini --console
```

別のコマンドプロンプトを管理者権限で立ち上げる

```
C:\dev\mysql-8.0.15-winx64> bin\mysql -u root -p
Enter password: ************ # さっきのPasswordを入力
Welcome to the MySQL monitor. Commands end with ; or \g.
Your MySQL connection id is 11
Server version: 8.0.15

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

# mysqlが起動した！取り敢えず安心！
mysql>
```

### root のpassword変更

```
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY 'new password';
Query OK, 0 rows affected (0.02 sec)
```

### インストールされているUserを確認

```
mysql> SELECT user, host, password_expired, password_lifetime FROM mysql.user;
+------------------+-----------+------------------+-------------------+
| user | host | password_expired | password_lifetime |
+------------------+-----------+------------------+-------------------+
| mysql.infoschema | localhost | N | NULL |
| mysql.session | localhost | N | NULL |
| mysql.sys | localhost | N | NULL |
| root | localhost | N | NULL |
+------------------+-----------+------------------+-------------------+
4 rows in set (0.00 sec)
```

以下が予約済みアカウント

- mysql.infoschema
- mysql.session
- mysql.sys

### 停止

```
C:\dev\mysql-8.0.15-winx64> bin\mysqladmin -u root -p shutdown
```

## DBの作成

DjangoでDBを切り替える際に事前に色々作成しておく必要がある。

### アカウントの作成

kudoaで設定する

```
mysql> CREATE USER IF NOT EXISTS 'kudoa'@'localhost'
-> IDENTIFIED BY 'kudoa'
-> PASSWORD EXPIRE NEVER
-> ;
Query OK, 0 rows affected (0.02 sec)
```

### Passwordの設定

```
mysql> ALTER USER 'kudoa'@'localhost' IDENTIFIED WITH mysql_native_password BY 'New Password';
Query OK, 0 rows affected (0.01 sec)
```

### DBの作成

今回はBLOGというDBを作成する

```
mysql> create database BLOG;
```

## DjangoにDBを登録する

DjangoとMySQLの接続には専用のパッケージが必要なのでインストールする。

```
pip install PyMySQL
```

### manage.pyが接続できるように登録しておく

```python
# manage.py
import pymysql

pymysql.install_as_MySQLdb()
```

### settings.pyにDBを登録する

```python
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'BLOG',  # DB名を設定
        'USER': 'kudoa',  # DBへ接続するユーザIDを設定
        'PASSWORD': 'NewPassword',  # DBへ接続するユーザIDのパスワードを設定
        'HOST': '',
        'PORT': '',
        'OPTIONS': {
            'init_command': "SET sql_mode='STRICT_TRANS_TABLES'",
        },
        'TEST': {
            'NAME': 'test_sample'
        }
    }
}
```

ここから先はモデルを作成するところまでの手順を割愛

今回は簡易的なブログを作るため、以下のようにmodelを設計した

```python
# models.py
from django.db import models
from django.utils import timezone


class Category(models.Model):
    name = models.CharField('カテゴリー名', max_length=50)
    pub_date = models.DateTimeField('作成日', default=timezone.now)

    def __str__(self):
        return self.name


class Post(models.Model):
    title = models.CharField('題名', max_length=50)
    text = models.TextField('本文')
    pub_date = models.DateTimeField('作成日', default=timezone.now)
    category = models.ForeignKey(
        Category, verbose_name='カテゴリー', on_delete=models.PROTECT
    )

    def __str__(self):
        return self.title
```

### make migrationsとmigrateを行う

```
(env) C:\..\myblog>py manage.py makemigrations
C:\..\env\lib\site-packages\pymysql\cursors.py:170: Warning: (3135, "'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in
a future release.")
  result = self._query(query)
Migrations for 'blogs':
  blogs\migrations\0001_initial.py
    - Create model Category
    - Create model Post
```

```
(env) C:\..\myblog>py manage.py migrate
C:\..\env\lib\site-packages\pymysql\cursors.py:170: Warning: (3135, "'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in
a future release.")
  result = self._query(query)
Operations to perform:
  Apply all migrations: admin, auth, blogs, contenttypes, sessions
Running migrations:
  Applying contenttypes.0001_initial... OK
  Applying auth.0001_initial... OK
  Applying admin.0001_initial... OK
  Applying admin.0002_logentry_remove_auto_add... OK
  Applying admin.0003_logentry_add_action_flag_choices... OK
  Applying contenttypes.0002_remove_content_type_name... OK
  Applying auth.0002_alter_permission_name_max_length... OK
  Applying auth.0003_alter_user_email_max_length... OK
  Applying auth.0004_alter_user_username_opts... OK
  Applying auth.0005_alter_user_last_login_null... OK
  Applying auth.0006_require_contenttypes_0002... OK
  Applying auth.0007_alter_validators_add_error_messages... OK
  Applying auth.0008_alter_user_username_max_length... OK
  Applying auth.0009_alter_user_last_name_max_length... OK
  Applying blogs.0001_initial... OK
  Applying sessions.0001_initial... OK
```

なんか謎のエラー文があるけど取り敢えずOKっぽい感じ。試しに管理画面からDataを挿入してみたが特に不具合はなかったので大丈夫そう。でも後々不具合が出るとなると少し怖い。

```python
# admin.py
from djangp.contrib import admin
from .models import Category, Post

admin.site.register(Category)
admin.site.register(Post)
```

## 追記: 次回以降の起動20190410

**サーバーを自分で起動させないとエラーを吐く**ことを知った。これのせいで大苦戦した。

```
ERROR 2003 (HY000): Can't connect to MySQL server on localhost (10061)
```

サーバーを起動してからpy manage.py runserverで問題なく動作した

### Herokuにデプロイするとき

いちいち動かすときにいちいちサーバーを起動させないといけないとなると結構困る。どうすればよいのやら、、、

