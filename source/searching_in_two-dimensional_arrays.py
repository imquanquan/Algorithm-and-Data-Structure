#!/usr/bin/env python3.6
# -*- coding:utf-8 -*-

"""
题目：二维数组中，每行从左到右递增，每列从上到下递增，给出一个数，判断它是否在数组中
"""


def find_num(matrix, num):
    """
    :param matrix:  [[]]
    :param num: int
    :return:  bool
    """
    if not matrix:
        return
    cols, rows = len(matrix[0]), len(matrix)
    col, row = cols - 1, 0

    while col >= 0 and row <= rows - 1:
        if matrix[col][row] == num:
            return True
        elif matrix[col][row] > num:
            col -= 1
        else:
            row +=1
    return False


if __name__ == '__main__':
    matrix = [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
    print(find_num(matrix, 8))
    print(find_num(matrix, 25))