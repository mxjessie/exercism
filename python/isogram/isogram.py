""" isogram: test isogramicity """

def is_isogram(string):
    """ True if string is an isogram """
    if string == "":
        return True
    return max([string.lower().count(x) for x in string.lower() if x.isalpha()]) <= 1
