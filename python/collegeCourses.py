"""
https://app.codesignal.com/arcade/python-arcade/fumbling-in-functional/hJqXodrjeBDPZPZRn

Given a list of courses, remove the courses with titles consisting of x letters and return the result.

Example

For x = 7 and
courses = ["Art", "Finance", "Business", "Speech", "History", "Writing", "Statistics"],
the output should be
collegeCourses(x, courses) = ["Art", "Business", "Speech", "Statistics"]
"""
def collegeCourses(x, courses):
    def shouldConsider(course):
        return len(course) != x

    return filter(shouldConsider, courses)
