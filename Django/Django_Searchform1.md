# Djangoの検索フォーム実装

## 選択肢から絞り込む

>部署と部活を指定して、そこに該当している人を表示する

## 流れ

>  - forms.py : フォームのクラスで、データ形式を定義する
> - views.py : フォームオブジェクトをテンプレートに渡す 
>   フォームから受け取ったデータを処理
> - index.html: フォームを表示する

## フォームの作成

```python
# forms.py

from django import forms
from .models import Club, Department


class SearchForm(forms.Form):
    # 文字を入力する欄が生成される(ModelChoiceFieldはモデルを基に生成する)
    # 選択欄がセレクトタグで生成される
    club = forms.ModelChoiceField(
        # querysetで何を選択肢にするか選ぶ、labelは枠の名前、requiredはその項目が必須かどうか
        queryset=Club.objects, label='サークル', required=False
    )

    department = forms.ModelChoiceField(
        queryset=Department.objects, label='部署', required=False
    )
```

## モデル

>外部キーに部署を指定する。部活は多対多で指定する。

```python
from django.db import models
from django.utils import timezone


class Department(models.Model):
    name = models.CharField('部署名', max_length=200)
    created_at = models.DateTimeField('日付', default=timezone.now)

    def __str__(self):
        return self.name


class Club(models.Model):
    name = models.CharField('部活名', max_length=20)
    created_at = models.DateTimeField('日付', default=timezone.now)

    def __str__(self):
        return self.name


class Employee(models.Model):
    first_name = models.CharField('名', max_length=20)
    last_name = models.CharField('性', max_length=20)
    email = models.EmailField('メールアドレス', blank=True)
    department = models.ForeignKey(
        Department, verbose_name='部署', on_delete=models.PROTECT,
    )
    club = models.ManyToManyField(
        Club, verbose_name='部活',
    )
    created_at = models.DateTimeField('日付', default=timezone.now)
    address = models.CharField('住所', max_length=20, blank=True)

    def __str__(self):
        return '{0}{1}{2}'.format(self.last_name, self.first_name, self.department)
```

## 汎用ビューを用いた書き方

>コードが少なくなる分見やすくなるが、内部的に何をしているのかをしっかりと理解しておく必要がある。

```python
# views.py
from django.views import generic
from .forms import SearchForm
from .models import Employee

# 汎用ビューは内部的に様々なメソッドを持っている(get_context_data, get_querysetもそれら)
class IndexView(generic.ListView):
    model = Employee
    template_name = 'employee/index.html'
    
    # templateに渡す辞書を作成する
    def get_context_data(self):
        context = super().get_context_data()
        # もとの辞書にフォームを追加、検索を内容をGETで渡すことで次見た時も検索済みになっている
        context['form'] = SearchForm(self.request.GET)
        return context
    
    # 全てのデータを呼び出すときに使用される
    def get_queryset(self):
        form = SearchForm(self.request.GET)
        """
        form自体は空欄でもOKなように定義した(forms.py)が、
        is_valid()を定義しないとform.cleaned_data[]にアクセス出来ない。
        """
        form.is_valid()
        
        # 全社員のデータを取得する(ここから入力内容に応じて絞り込む)
        queryset = super().get_queryset()
        
        # form入力欄にアクセスしたいのでcleaned_data[]を使う、値をdepartmentに入れる
        # 今回はManyToManyFieldを使っているので、departmentにmodelのinstanceが格納される
        department = form.cleaned_data['department']
        # もしdepartmentに値が入っていれば、部署で絞り込む(filter)
        if department:
            # filter(field名=変数名)
            queryset = queryset.filter(department=department)

        club = form.cleaned_data['club']
        if club:
            # queryset = queryset.filter(name='水泳')
            queryset = queryset.filter(club=club)
        return queryset
```

## template

>formはfromタグが用意されている。modelはfor文で回してやればよい。

```html
{% extends 'employee/base.html' %}

{% block title %}社内管理システム{% endblock %}

{% block content %}
<!--formの書き方と同じ、GETなのでcrfトークンは必要ない-->
<form action='' method='GET'>
    {{ form }}
    <button type='submit'>検索</button>
</form>

<head>一覧</head>
<table class='table table-bordered'>
    <thead>
        <tr>
            <th>ID</th>
            <th>姓</th>
            <th>名</th>
            <th>メール</th>
            <th>部署</th>
            <th>部活</th>
        </tr>
    </thead>
    <tbody>
        {% for employee in employee_list %}
        <tr>
            <td>{{ employee.pk }}</td>
            <td>{{ employee.last_name }}</td>
            <td>{{ employee.first_name }}</td>
            <td>{{ employee.email }}</td>
            <td>{{ employee.department }}</td>
            <td>{% for club in employee.club.all %}{{ club }}{% endfor %}</td>
        </tr>
        {% endfor %}
    </tbody>
</table>
{% endblock %}
```

