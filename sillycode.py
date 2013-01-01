#!/usr/bin/python
import re

vowels = re.compile('(a|e|i|o|u)', re.IGNORECASE)

def oddletters(word):
  even = True
  res = ""

  for letter in word:
    if (even == True):
      even = False
    else:
      even = True
      res = "".join([res, letter])

  return res

def evenletters(word):
  even = True
  res = ""

  for letter in word:
    if (even == True):
      even = False
      res = "".join([res, letter])
    else:
      even = True

  return res

def placename(name):
  
  # If the name is shorter than 4 chars, pad with X and don't strip vowels
  if (len(name) < 4):
    while len(name) < 4:
      name = ''.join([name, 'X'])
    return name.upper()

  # If name is exactly 4 chars, just return it
  if (len(name) == 4):
    return name.upper()
  
  newname = ""

  words = name.split()

  if (len(words) == 2):
    newname = ""

    name1 = words[0]
    name2 = words[1]
    
    # handle first word
    first = name1[0]
    rest = vowels.sub('', name1[1:])

    name = ''.join([first, rest])

    if (len(name) == 1):
      name = ''.join([word, 'X'])

    elif (len(name) == 2):
      pass
      
    else:
      name = ''.join([first, rest[0]])
    
    newname = name

    # handle second word
    first = name2[0]
    rest = vowels.sub('', name2[1:])

    name = ''.join([first, rest])

    if (len(name) == 1):
      name = ''.join([word, 'X'])

    elif (len(name) == 2):
      pass

    else:
      name = ''.join([first, rest[-1]])

    newname = ''.join([newname, name])

    return newname.upper()

  elif (len(words) > 2):
    return "Input too long"

  else: # only one word in name
    first = name[0]
    word = name[1:]
    word = vowels.sub('', word)

    word = ''.join([first, word])

    if (len(word) < 4): # word without vowels is too short so pad
      while len(word) < 4:
        word = ''.join([word, 'X'])

    elif (len(word) == 4): # len is perfect
      pass

    elif (len(word) == 5): # one too long so return first 3 and last
      first3 = word[:3]
      last = word[-1]
      option = word[:4]
      word = ''.join([first3, last])

      if (word.upper() == option.upper()):
        return word.upper()
      else:
        return [word.upper(), option.upper()]
    
    else: # too long so even letters only
      word = ''.join([word[0], evenletters(word[1:])])

    return word.upper()

if __name__ == '__main__':
  names = [
    "River Bend", 
    "Hartford", 
    "Watertown", 
    "Waterloo", 
    "Erin", 
    "Reeseville", 
    "Madison", 
    "North Lake",
    "Denver",
    "Ashburn",
    "New Berlin",
    "WAUKESHA",
    "Cedarburg",
    "Delafield",
    "Lake Geneva",
    "North Prairie",
    "Rio",
    "Re",
    "Oconomowoc",
    "Hubbelton",
    "Columbus",
    "Clyman",
  ]

  for name in names:
    print "\nName: %s" % (name)
    print "Silly: %s" % (placename(name))


