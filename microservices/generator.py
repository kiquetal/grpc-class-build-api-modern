def countdown(n):
    while n > 0:
        print("Countdown:", n)
        yield
        print("paro")
        n -= 1

def launch():
    print("Launching...")
    yield
    print("Rocket launched!")

# Create generator objects
countdown_gen = countdown(5)

# Run the generators
while True:
    print("here")
    try:
        next(countdown_gen)
    except StopIteration:
        # Countdown generator is completed
        break

    print("waiting for")
