"""
https://app.codesignal.com/arcade/python-arcade/caravan-of-collections/3q55u2MWA2Rw5HvmM/

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
You're implementing a plugin for your favorite code editor. This plugin launches various scripts depending on the open file extension. Each script is associated with exactly one extension, and the information about which script should be launched for each extension is stored in a dictionary scriptByExtension.

You are planning to add more supported extensions for some scripts, so now you would also like to store information about the extensions which each script supports. As a starting point, you'd like to obtain the (extension, script) pairs from the dictionary, sorted lexicographically by the extensions.

Implement a function that will do the job.

Example

For

scriptByExtension = {
  "validate": "py",
  "getLimits": "md",
  "generateOutputs": "json"
}
the output should be

transposeDictionary(scriptByExtension) = [["json", "generateOutputs"], 
                                          ["md", "getLimits"], 
                                          ["py", "validate"]]
"""
def transposeDictionary(scriptByExtension):
    return [[x[1], x[0]] for x in sorted(scriptByExtension.items(), key = lambda x:x[1])]
