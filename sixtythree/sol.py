sum = 0

for i in range(0, 300):
    for j in range(1, 300):
        if len(str(i ** j)) == j:
            print('{}^{} is a {}-digit number: {}'.format(
                i, j, j, i ** j))
            sum += 1

print(sum)