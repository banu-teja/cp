from typing import List


class UnionFind:
    def __init__(self, n):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x):
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        px, py = self.parent[x], self.parent[y]
        if px == py:
            return False

        if self.rank[px] < self.rank[py]:
            self.parent[px] = py
        elif self.rank[px] > self.rank[py]:
            self.parent[py] = px
        else:
            self.rank[px] += 1
            self.parent[py] = px
        return True


class Solution:
    def minCostConnectPoints(self, points: List[List[int]]) -> int:
        n = len(points)
        edges = []

        for i in range(n):
            for j in range(i + 1, n):
                dist = abs(points[i][0] - points[j][0]) + abs(
                    points[i][1] - points[j][1]
                )
                edges.append((dist, i, j))

        edges.sort()

        uf = UnionFind(n)
        total_cost = 0
        num_edges = 0

        for weight, u, v in edges:
            if uf.union(u, v):
                total_cost += weight
                num_edges += 1
                if num_edges == n - 1:
                    break

        return total_cost


def test_minCostConnectPoints():
    solution = Solution()

    points1 = [[0, 0], [2, 2], [3, 10], [5, 2], [7, 0]]
    assert solution.minCostConnectPoints(points1) == 20

    points2 = [[3, 12], [-2, 5], [-4, 1]]
    assert solution.minCostConnectPoints(points2) == 18

    points3 = [[0, 0], [1, 1], [1, 0], [-1, 1]]
    assert solution.minCostConnectPoints(points3) == 4

    points4 = [[0, 0]]
    assert solution.minCostConnectPoints(points4) == 0

    print("All test cases passed!")


test_minCostConnectPoints()
