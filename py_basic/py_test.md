#　doctest

本格的なテストというよりは、例を示して他のプログラマーに使い方を教えるついでのものという感じ

```python
import unitest


class Cal(object):
    def add_num_and_double(self, x, y):
        # doctestの記述方法
        """Add and double

        >>> c = Cal()
        >>> c.add_num_and_double(1, 1)
        4
        
        # 対策と書き方
        >>> c.add_num_and_double('1', '1')
        Traceback (most recent call last):
        ...
        ValueError
        """
        # int以外をキャッチアップできるように！
        if type(x) is not int or type(y) is not int:
            raise ValueError
        result = x + y
        result *= 2
        return result

# import時に動かないようにする
if __name__ == "__main__":
    import doctest
    doctest.testmod()
```

#　Unitest

他のファイルからimportしたモジュールをテストする

```python
# calcuration.py(今回はこのファイルをテストする)
class Cal(object):
    def add_num_and_double(self, x, y):
        if type(x) is not int or type(y) is not int:
            raise ValueError
        result = x + y
        result *= 2
        return result
```

```python
import unittest

import calcuration

# unitestからTestCaseを継承させる
class CalTest(unittest.TestCase):
    def test_add_num_and_double(self):
        cal = calcuration.Cal()
        self.assertEqual(
            cal.add_num_and_double(1, 1),
            5)
        
if __name__ == '__main__':
    unitest.main()

>>>
FAIL: test_add_num_and_double (__main__.CalTest)
----------------------------------------------------------------------
Traceback (most recent call last):
...
AssertionError: 4 != 5

----------------------------------------------------------------------
Ran 1 test in 0.000s

FAILED (failures=1)
```

##　例外テスト

```python
import unittest

import calcuration


class CalTest(unittest.TestCase):
    def test_add_num_and_double(self):
        cal = calcuration.Cal()
        self.assertEqual(
            cal.add_num_and_double(1, 1),
            4)
	# withステートメントで記述する
    def test_add_num_and_double_raise(self):
        cal = calcuration.Cal()
        with self.assertRaises(ValueError):
            cal.add_num_and_double('1', '1')


if __name__ == "__main__":
    unittest.main()
>>
# ちゃんとValueErrorが出たのでテストはパスされた！
----------------------------------------------------------------------
Ran 2 tests in 0.001s

OK
```

## setUp, tearDown

```python
import unittest

import calcuration


class CalTest(unittest.TestCase):
    def setUp(self):
        print('setup')
        self.cal = calcuration.Cal()

    def tearDown(self):
        print('clean up')
        del self.cal

    def test_add_num_and_double(self):
        self.assertEqual(
            self.cal.add_num_and_double(1, 1),
            4)

    def test_add_num_and_double_raise(self):
        with self.assertRaises(ValueError):
            self.cal.add_num_and_double('1', '1')


if __name__ == "__main__":
    unittest.main()
    
>>>
setup
clean up
.setup
clean up
.
----------------------------------------------------------------------
Ran 2 tests in 0.001s

OK
```

## skip

> 特定のテストをコードを消さずにスキップできるデコレーターがある

```python
import unittest

import calcuration


class CalTest(unittest.TestCase):
    def setUp(self):
        print('setup')
        self.cal = calcuration.Cal()

    def tearDown(self):
        print('clean up')
        del self.cal
    
    # テストをスキップして"skip"と表示
    @unitest.skip("skip")
    def test_add_num_and_double(self):
        self.assertEqual(
            self.cal.add_num_and_double(1, 1),
            4)

    def test_add_num_and_double_raise(self):
        with self.assertRaises(ValueError):
            self.cal.add_num_and_double('1', '1')


if __name__ == "__main__":
    unittest.main()
```

```python
# 特定の条件下でのスキップもできる
@unitest.skipIf(release_name="lesson", "skip!") 
def test_add_num_and_double(self):
        self.assertEqual(
            self.cal.add_num_and_double(1, 1),
            4)
```

#　pytest

pytestを事前にインストールしておく

```python
# calcuration.py
class Cal(object):
    def add_num_and_double(self, x, y):
        if type(x) is not int or type(y) is not int:
            raise ValueError
        result = x + y
        result *= 2
        return result
```

```python
# pytest_p.py
import calcuration

# 簡単にテストできる
def test_add_num_and_double():
    cal = calcuration.Cal()
    assert cal.add_num_and_double(1, 1) == 4
    
# 実行はpytest File名    
>>> pytest pytest_p.py
...
collected 1 item  
...
1 passed in 0.03 seconds
```

```python
# classでもテストできる
import calcuration

# Testで始める、継承する必要はない
class TestCal(object):
    def test_add_num_and_double(self):
        cal = calcuration.Cal()
        assert cal.add_num_and_double(1, 1) == 4
```

## 例外テスト

ちゃんと例外処理がされるかのテスト

```python
import pytest

import calcuration


class TestCal(object):
    def test_add_num_and_double(self):
        cal = calcuration.Cal()
        assert cal.add_num_and_double(1, 1) == 4
        
    # 例外テスト
    def test_add_num_and_double_raise(self):
        with pytest.raises(ValueError):
            cal = calcuration.Cal()
            cal.add_num_and_double('1', '1')
```

