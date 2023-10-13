import re, sys

input = ""

if len(sys.argv) < 2:
    print("No input file specified")
    exit(1)
with open(sys.argv[1]) as file:
    input = file.read()
print(input,"\nResult")
pattern = r'''
        (?:                 # Формат дня
            (?:0?[1-9])     # 01-09 1-9 
            |               # ИЛИ
            (?:[1-3][0-9])  # 10-31
        )
        ([.\/\\\-\s])       # Разделитель между днем и месяцем
        (?:                 # Формат месяца
            (?:0?[1-9])     # 01-09 1-9
            |               # ИЛИ    
            (?:1[0-2])      # 10-12
        )
        \1                  # Такой же разделитель
        (?:                 # Год
            (?:[12]\d{3})   # 1984 2007
            |               # ИЛИ
            (?:\d\d)        # 09 84 22 
        )
        '''
r = re.finditer(pattern, input, re.VERBOSE)
for i in r:
    print(i.group())
    