def main():
    a = int(input("1: "))
    b = int(input("2: "))
    c = int(input("3: "))
    d = int(input("4: "))
    e = int(input("5: "))

    a = int(a/3)
    b += a
    e += a

    b = int(b/3)
    c += b
    a += b

    c = int(c/3)
    d += c
    b += c

    d = int(d/3)
    e += d
    c += d

    e = int(e/3)
    a += e
    d += e

    print(a, b, c, d, e)
    return ""


def advance(n):
    # 进阶写法：n 为 n 个小朋友
    l = []
    for i in range(n):
        l.append(int(input(f"第{i+1}个小朋友: ")))

    for i in range(n):
        share = int(l[i]/3)
        l[i] = share
        if i - 1 >= 0:
            l[i-1] += share
        else:
            l[n-1] += share
        if i + 1 >= n:
            l[0] += share
        else:
            l[i+1] += share
    print(l)
    return ""


if __name__ == "__main__":
    print(advance(5))
