# ファイル操作とシステム

## ファイルの作成

> 　今いるコマンドでファイルが作成されることに注意

```python
# "w"は書き込みモード, test.txtが作成される(ファイルがあったら上書きされるので注意)
f = open("test.txt", "w")
# test.txtにTestと書き込まれる
# print("hogehoge", file=f)でもかけるが f.writeの方が分かりやすいので好まれる
f.write("Test")
# fileを使い終わったらcloseしましょう
f.close()
```

```python
# 追加で文字列を入れたい場合は"a"
with open("new_test.csv", "a", newline="") as f:
    writer = csv.writer(f)
    writer.writerow([1, 'スパム', '500円'])
    writer.writerow([2, '卵', '168円'])
    writer.writerow([3, 'ベーコン', '1,250円'])
    
# 実行するたびにデータが足される＝データの保持
```

## withステートメント

```python
# f.closeがなくてもcloseできる
with open("test.txt", "w") as f:
    f.write("hoge")
```

## ファイルの読み込み

```python
s = """\
AAA
BBB
CCC
DDD
"""
# 書き込みモード
# with open("test.txt", "w") as f:
    # f.write(s)

# 読み込みモード(書き込まれているデータをコマンド上に表示)
with open("test.txt", "r") as f:
    # まとめて一気に読み込み
    print(f.read())
```

```python
# 一列ずつ読み込む
with open("test.txt", "r") as f:
    while True:
        line = f.readline()
        # print()は改行を入れる
        print(line)
        if not line:
            break
```

> 　出力結果

```
AAA

BBB

CCC

DDD
```

```python
# chunk毎に読み込む
with open("test.txt", "r") as f:
    while True:
        chunk = 2
        line = f.read(chunk)
        print(line)
        if not line:
            break
```

## seek

```python
with open("test.txt", "r") as f:
    # 今いる位置が返る
    print(f.tell())
    >> 0
    print(f.read(1))
    >> A
    f.seek(5)
    print(f.read(1))
    >> B
    f.seek(14)
    print(f.read(1))
    >> D
    f.seek(15)
    print(f.read(1))
    >> 
```

## 書き込み読み込み

> w+は新しいもの書き込む状態になるので、書き込む前に読み込むと何も入ってない状態になる

```python
# 書き込んだ後に読み込みも出来る
with open("test.txt", "w+") as f:
    f.write(s)
    # 書き込んだ後は一番先頭に戻る必要がある
    f.seek(0)
    print(f.read())
```

```python
# 事前にそのファイルがないとエラーになる
with open("test.txt", "r+") as f:
    print(f.read())
    f.seek(0)
    f.write(s)
```

## テンプレート

> $に文字を代入できる
>
> 元の文字列を扱うとそれが壊される可能性もあるので読み込み専用として置いておく

```python
import string
s = """\

Hi, $name

$contents

Have a good day
"""

t = string.Template(s)
contents = t.substitute(name="Mike", contents="How are you?")
print(contents)
```

> デザイナーチームが作ったデザインのファイルを読み込み時にも使える
>
> プログラマーチームとデザイナチームを完全に分ける時
>
 ```
# design\email_template_text
Hi, $name

$contents

Have a good day
 ```
```python
# テンプレートで読み込むものを上のテキストから読み込む
with open("design\email_template_text") as f:
    # tはwithの外でも使える
    t = string.Template(f.read())

contents = t.substitute(name="Mike", contents="How are you?")
print(contents)
```

## CSVファイルへの書き込み、読み込み

### 書き込み

```python
import csv
with open("test.csv", "w", newline="") as csv_file:
    fieldnames = ["Name", "Count"]
    writer = csv.DictWriter(csv_file, fieldnames=fieldnames)
    writer.writeheader()
    writer.writerow({"Name": "A", "Count": "1"})
    writer.writerow({"Name": "B", "Count": "2"})
```

### 読み込み

```python
with open("test.csv", "r") as csv_file:
    reader = csv.DictReader(csv_file)
    for row in reader:
        print(row)
# listにtupleのセットになっている
>> 
OrderedDict([('Name', 'A'), ('Count', '1')])
OrderedDict([('Name', 'B'), ('Count', '2')])
        print(row["Name"], row["count"])
```

## ファイル操作

> 四つのライブラリさえ覚えておけばやりたいことはほとんど出てくる

```python
import os
import pathlib
import glob
import shutil

print(os.path.exists("test.txt"))
>> True
print(os.path.isfile("test.txt"))
>> True
print(os.path.isdir("test.txt"))
# file名の変更
os.rename("test.txt", "renamed.txt")
# shortcut copy
os.symlink("renamed.txt", "symlink.txt")
# dir作成、削除
os.mkdir("test_dir")
os.rmdir("test_dir")
# empty.txt作成、削除
pathlib.Path("empty.txt").touch()
os.remove("empty.txt")
# 所定のファイルを表示
print(glob.glob("test_dir\test_dir2"\*))

# 階層dirの作成、確認、削除
os.mkdir("test_dir")
os.mkdir("test_dir\test_dir2")
print(os.listdir("test_dir"))
>> test_dir2

# 中身ごとdir全部削除(注意)
shutil.rmtree("test_dir")
# 今実行しているファイルの位置
print(os.getcwd())
```

