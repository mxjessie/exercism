def is_pangram(inpan=""):
    """ what a disaster lmao """
    return False not in [(y in inpan.upper()) for y in
                         [chr(x) for x in range(65, 65+26)]]
