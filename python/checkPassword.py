"""
https://app.codesignal.com/arcade/python-arcade/yin-and-yang/NXz5g32Puypw3R97N

You're implementing a web application. Like most applications, yours also has the authorization function. Now you need to implement a server function that will check password attempts made by users. Since you expect heavy loads, the function should be able to accept a bunch of requests that will be sent simultaneously.

In order to validate your function, you want to test it locally. Given a list of attempts and the correct password, return the 1-based index of the first correct attempt, or -1 if there were none.

Example

For attempts = ["hello", "world", "I", "like", "coding"] and
password = "like", the output should be
checkPassword(attempts, password) = 4.
"""
def checkPassword(attempts, password):
    def check():
        while True:
            attemp = yield
            if attemp == password:
                yield True


    checker = check()
    for i, attempt in enumerate(attempts):
        next(checker)
        if checker.send(attempt):
            return i + 1

    return -1
