
t = int(input())

for _ in range(t):
    ships = list(map(int, input().split()))
    fours = ships.count(4)
    threes = ships.count(3)
    twos = ships.count(2)
    ones = ships.count(1)

    if fours == 1 and threes == 2 and twos == 3 and ones == 4:
        print('YES')
    else:
        print('NO')