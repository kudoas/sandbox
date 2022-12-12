# Djangoの検索機能の実装

> SearchFormに検索ワードを入れて該当している内容(title, text)の記事だけ抜き出す

## クラスビューを使用した場合

### HTML

> bootstrap4.3vのnavbarを使用

```html
<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="#">Navbar</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item active">
        <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="#">Link</a>
      </li>
      <li class="nav-item dropdown">
        <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
          Dropdown
        </a>
        <div class="dropdown-menu" aria-labelledby="navbarDropdown">
          <a class="dropdown-item" href="#">Action</a>
          <a class="dropdown-item" href="#">Another action</a>
          <div class="dropdown-divider"></div>
          <a class="dropdown-item" href="#">Something else here</a>
        </div>
      </li>
      <li class="nav-item">
        <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">Disabled</a>
      </li>
    </ul>
 
<!--
inputタグで検索ワードをname=keywordとして受け取る
acitonでURLを指定しているのは、このnavbarがどの画面でもindexの検索として機能してほしいため
(空の場合はそのページ内の記事検索機能になる)
-->
    <form class="form-inline my-2 my-lg-0" method='GET' action='{% url 'blog:index' %}'>
      <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search" name="keyword">
      <button class="btn btn-outline-success my-2 my-sm-0" type="submit">検索</button>
    </form>
  </div>
</nav>
```

### model

```python
# models.py
from django.db import models
from django.utils import timezone

class Post(models.Model):
    title = models.Charfield('title', max_length=200)
    text = models.Textfield('text')
    pub_date = models.DateTimeField('published', dafault=timezone.now)
```

### views.pyの編集

```python
from django.views import generic
from .model import Post

class IndexView(generic.ListView):
    model = Post
    
    def get_queryset(self):
        # model.objectを日付順にソート、記事ごと(タイトル、発行日、日付)取り出す
        queryset = Blog.objects.order_by('-create_at')
        
        # 検索フォーム(HTML)から値を受け取る
        keyword = self.request.GET.get('keyword')
        if keyword:
        
            # keywordとtitleが完全一致しているものをfilter
            queryset = queryset.filter(title=keyword)
        
        return queryset
```

#### 部分一致検索の場合

```python
# __icontains

# 大文字小文字を区別する場合は__containを使う
queryset = queryset.filter(title__icontains=keyword)
```

#### タイトルと本文でOR検索させる場合

```python
from djnago.db.models import Q
from django.views import generic
from .model import Post

class IndexView(generic.ListView):
    model = Post
    
    def get_queryset(self):
        queryset = model.objects.ordered_by('-create_at')
        keyword = self.request.GET.get('keyword')
        if keyword:
            queryset = queryset.filter(
                Q(title__icontains=keyword) | Q(text__icontains=keyword)
            )
        
        return queryset
```

## 関数ビューを使用した場合

> 意外に単純なコードで実装可能である

### model

```python
class Blog(models.Model):
    title = models.CharField(blank=False, null=False, max_length=150)
    text = models.TextField(blank=True)
    created_datetime = models.DateTimeField(auto_now_add=True)
    updated_datetime = models.DateTimeField(auto_now=True)
```



### form

```python
from django import forms

from .models import Blog

class BlogSearchForm(forms.Form):
    title = forms.CharField(label='検索', required=False)
```



### view

```python
from django.db.models import Q
from django.shortcuts import render

from .form import BlogSearchForm
from .models import Blog

def index(request):
    form = BlogSearchForm(request.POST)
    blogs = Blog.objects.order_by('-created_datetime')

    if request.method == 'POST':
        # formはrequired=Falseだが、cleaned_dataで受け取るにバリデーションが必要
        if form.is_valid():
            blogs = blogs.filter(
                Q(title__icontains=form.cleaned_data['title']) | Q(
                    text__icontains=form.cleaned_data['title'])
            )

    contexts = {
        'form': form,
        'blogs': blogs,
    }

    return render(request, 'blogs/index.html', contexts)
```



### HTML

> 　単純にformで受け取ってあげればOK

```python
<form action="" method="POST" style="text-align: center">
            {% csrf_token %}
            {{ form }}
            <button type="submit">検索</button>
</form>
```

