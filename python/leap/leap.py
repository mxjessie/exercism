""" leap: leap year testing! """

def is_leap_year(year):
    """ Returns true if year is a leap year """
    return year % 4 == 0 if year % 100 != 0 else year % 400 == 0
