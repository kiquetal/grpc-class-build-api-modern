def countdown(n):
    while n > 0:
        print("generator value n",n)
        yield n
        print("resuming generator")
        n -= 1


countdown_gen = countdown(5)

print("here calling the generaor")
n = next(countdown_gen)
print("caller " , n)
n = next(countdown_gen)

