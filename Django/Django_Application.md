#　日記あぷり作るぞー！おぉー！

##　初期設定

> プロジェクトの作成

```
> py -m django startproject second
```

>  アプリを作成する

```
> cd second
> py manage.py startapp diary
```

### それぞれのファイル構造

```
second/
    manage.py
    second/
        __init__.py
        settings.py
        urls.py
        wsgi.py
        
# secondフォルダ内        
diary/
    __init__.py
    admin.py
    apps.py
    migrations/
        __init__.py
    models.py
    tests.py
    views.py
```

###　appliの登録

```python
# setting.py
INSTALLED_APPS = [
    # 書き方は大体全部同じ。diaryフォルダ内のapps.pyに自動登録
    "diary.apps.DiaryConfig",
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
]

# app.py
from django.apps import AppConfig

# 自動登録されている
class DiaryConfig(AppConfig):
    name = 'diary'
```

###　urlの登録

> 　viewで表示させるものを登録する

```python
# second/urls.py
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    # 管理画面、これだけ特別
    path('admin/', admin.site.urls),
    # なんちゃら/diaryのURLを認識。diary/urls.pyへ
    path('diary/', include('diary.urls'))
]
```

> 　diaryフォルダ内にurls.pyを作成する

```python
# diary/urls.py
from django.urls import path
from . import views

app_name = "diary"

urlspattern = [
    # /diaryの時に呼び出される
    path('', views.index, name='index')
]
```

#### path()関数とは

> 　以下の四つの引数を取る。引数のうち routeとviewの2つは必須で、kwargs、name の2つは省略可能

```python
path(route, view, kwargs, name)
# route: URL
# view: Django がマッチする正規表現を見つけると、 Django は指定されたビュー関数を呼び出す
# kwargs: 任意のキーワード引数を辞書として対象のビューに渡せる
# name: 名前付けをしておくとどこからでもURLを参照し呼び出せる(逆引き)
```



###　views.pyの編集

```python
from django.shortcuts import render


def index(request):
    return render(request, 'diary/day_list.html')
```

### templatesの作成

> 同じようなhtmlを使いまわす場合は

```html
# templates/diary/base.html
<!doctype html>
<html lang="ja">

<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
        integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

    <title>日記あぷり</title>
</head>

<body>
    {% block content %}
    {% endblock %}

    <!-- Optional JavaScript -->
    <!-- jQuery first, then Popper.js, then Bootstrap JS -->
    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
        crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"
        integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1"
        crossorigin="anonymous"></script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"
        integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    {% block extracjs %}{% endblock %}
</body>

</html>
```

> 　これを呼び出す

```html
# day_list.html
# ファイルを呼び出す
{% extends 'diary/base.html' %}

# 特定の領域を書き換える
{% block content %}
<h1>日記あぷりですよね！わかります！わかります！</h1>
{% endblock %}
```

## モデルを作成する

> models.Modelの書き方は決まっている。
>
> いずれもdjango.db.models.Modelのサブクラスとなっている

```python
from django.db import models
# django専用(pythonだとdatetime,now)
from django.utils import timezone

# models.Modelを継承する
# それぞれのFieldで受け取れるものが異なる
class Day(models.Model):
    # CharField: そこまで長くないテキストを受け取る(タイトルとか)
    title = models.CharField('タイトル', max_length=200)
    # TextField: テキストを受け取る
    text = models.TextField('本文')
    # DateTimeField: 日時を指定する
    date = models.DateTimeField('日付', default=timezone.now)
```

> 　モデルを有効にする

```python
# modelを作成
> py manage.py makemigrations diary
Migrations for 'diary':
  diary\migrations\0001_initial.py
    - Create model Day
"""
makemigrationsを実行することで、Djangoにモデルに変更があったこと(この場合、新しいものを作成しました)を伝え、そして変更をマイグレーションの形で保存
"""

# modelのdbの作成(DBの反映)
> py manage.py migrate
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, diary, sessions
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
  Applying diary.0001_initial... OK
  Applying sessions.0001_initial... OK
```

##　データの追加機能

> 　日記を書くたびにDBに保存されるシステムを作ろう

```python
# diary/urls.py
from django.urls import path
from . import views

app_name = "diary"

urlpatterns = [
    path('', views.index, name='index'),
    # addを登録する
    path('add', views.add, name='add'),
]
```

## model(User)

> setting.py

```
AUTH_USER_MODEL = 'diaries.User'
```

