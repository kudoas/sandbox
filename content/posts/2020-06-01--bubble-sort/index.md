---
title: バブルソート
category: "algorithm"
cover: ../cover-images/web-image.jpeg
author: Kudoa
---

面接でのコーディング試験で配列をソートしてくれという課題がでた。全然出来なくて悔しかったので**バブルソート**の勉強をしたのでその時のメモ。

## バブルソートとは

配列の一番左の値を取ってきて右隣のものと比較し、右隣の方が小さければ交換するアルゴリズム。
この作業を配列の長さ分行えば、全てソートされるので計算量はO(n^2)かかる。

## Pythonでの実装

実際に書くとこんな感じ

```python
def bubble_sort(A):
  	n = len(A)
		for _ in range(n):
      	for i in range(n-1):
          	if A[i] > A[i+1]:
              	A[i], A[i+1] = A[i+1], A[i]
    return A

# example
A = [2,3,4,1,5]
sorted_A = bubble_sort(A) # [1,2,3,4,5]
```

他にも挿入ソート、バケットソート、クイックソートなど様々なソートがある模様。

次、同じ課題がでたら絶対解いたるぞ！！