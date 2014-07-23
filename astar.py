# http://www.redblobgames.com/pathfinding/a-star/introduction.html

from Queue import Queue
from Queue import PriorityQueue

MATRIX = [
        [],
        [],
        [],
        [],
        [],
        ]

class Graph:
    def __init__(self, matrix):
        self.m_ = matrix
        self.row_ = len(matrix)
        self.col_ = len(matrix[0])
        for r in self.m_:
            if self.col_ > len(r):
                self.col_ = len(r)
  
    def valid(self, node):
        return node[0] < self.row_ and node[0] >= 0 \
                and node[1] < self.col_ and node[1] >= 0
  
    def neighborhoods(self, node):
        d = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        res = []
        for n in d:
            if self.valid([node[0]+d[0], node[1]+d[1]]):
              res.append([node[0]+d[0], node[1]+d[1]])
        return res

def breadth_first_search(start, target, graph):
    return path

