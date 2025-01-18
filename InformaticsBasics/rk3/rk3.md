% Рубежный контроль № 3: конспект по скриптовому языку
% 15 января 2025 г.
% Родион Лавров, ИУ9-12Б

# Язык программирования Python

## 1. Типизация и система типов языка
Типизация в Python динамическая, неявная, строгая.

| Тип данных |Неизменяемый|Описание|Пример|
|:----:|:----:|:----------:|:------:|
|   None     | + | Имеет единственное значение | `None`{.python} |
|   Bool     | + | Булевы значения (`True`, `False`) | `True`{.python}, `False`{.python} |
|  Bytes     | + | Последовательность отдельных байтов | `bytes(10)`{.python} |
|  bytearray | - | Последовательность отдельных байтов | `bytes(10)`{.python} |
|    Int     | + | Представление целых чисел | `-45`{.python}, `52`{.python} |
|   Float    | + | Числа с плавающей точкой | `14.676`{.python}, `5.43`{.python} |
|  Complex   | + | Комплексные числа | `1.23+0j`{.python}, `-1.23+4.5j`{.python} |
|     Str    | + | Последовательность кодовых точек Unicode | `"python"`{.python}, `"Hello World"`{.python} |
|   Tuple    | + | Хранение множества разнородных данных | `(32, 456)`{.python} |
|   List     | - | Хранение множества однородных  данных | `[32, 456]`{.python} |
|   Range    | + | Последовательность чисел, обычно используется в циклах for | `list(range(5))`{.python}  |
|    Dict    | - | Ассоциативный массив | `{"one": 1, "two": 2}`{.python} |
|     Set    | - | Хранение множества различных объектов с возможностью хэширования | `{'car', 32}`{.python} |
|  Frozenset | + | Хранение множества различных объектов с возможностью хэширования | `{'car', 32}`{.python} |

## 2. Основные управляющие конструкции

- Ветвление: `<if>::= if<condition>: <body> elif<condition>: <body> else: <body>`
- Цикл с условие: `<while-loop>::= while<condition>: <body>`
- Цикл с переменной: `<for-loop>::= for <var> in <iterable>: <body>`
- Прерывние: `break`
- Переход к следующей итерации: `continue`
- Ничего не делает: `pass`
- Аналог switch-case: `<match-statement>::= match <var>: case <pattern>: <body>`
- Объявление функций: `<define-func>::= def <name>(<param>*): <body>`
- Обработка ошибок: `<try>::= try: <body> except <exception>: <body> else: <body> finally: <body>`


## 3. Подмножество языка для функционального программирования
## Cпособы обеспечить иммутабельность данных
Такие типы данных, как `Set`{.js}, `Dict`{.bash}, `List`{.js}, `bytearray`{.python} поумолчанию являются 
мутабельными, в отличии от отстальных. Создать иммутабельный класс, можно различными способами. 
Вот несколько из них:

1) С помощью класса `property`{.python}, который позволяет задать определение, 
получение и изменение полей класса

```python
class Immutable():
   b = property(lambda s: "Hello World")

a = Immutable()
a.b = "mutated"  # AttributeError: property 'piece' of 'Immutable' object has no setter
```

2) C помощью метода `__setattr__`{.python}, который определяет каким образом будут изменяться поля класса
```python
class Immutable():
    def __init__(self):
        self.b = 2

    def __setattr__(self, name, value):
        raise Exception(f"Cannot change value of {name}.")

a = Immutable()
a.b = 1    # Exception: Cannot change value of b.
```

3) C помощью метода `__slots__`{.python}, который не позволяет добавлять новые атрибутов после инициализации. 
```python
class Immutable:
    __slots__ = ('x')  # Ограничение атрибутов

    def __init__(self, x):
        self.x = x

a = Immutable(1)
a.z = 4  # AttributeError: 'Immutable' object has no attribute 'z'
```

Тем не менее, иммутабельность пользовательских классов в Python является лишь формальной, 
так как её можно обойти. Например в примере 2 это можно сделать следующем образом
```python
a.__dict__["b"] = 1
```

## Функции, как объекты 1-го класса в Python
Все данные в Python представлены объектами или отношениями между объектами. 
Поэтому функции могут быть присвоены другим переменным,
```python
def func(string):
    return string.lower()

var = func
print(var("Hello World")) # hello world
```
переданы другим функциям в качестве аргуметов
```python
def greet(func):
    greeting = func('Hi, I am a Python program')
    print(greeting)
    
greet(func) # hi, i am a python program
```
и возвращены из функции
```python
def func2(func):
    return func
    
a = func2(func)("Hi, I am a Python program")
print(a) # hi, i am a python program
```

## Функции высших порядков
Функция высшего порядка — это функция, принимающая в качестве аргументов другие функции или возвращающая 
другую функцию в качестве результата. Реализация такой функции была приведена в примере выше. 
Рассмотрим встроенные функции высших порядков `map`{.python} и `filter`{.python}:
```python
numbers = [5, 6, 7, 8, 9, 10]

new_numbers = map(lambda x: x ** 2, numbers)
print(list(new_numbers))   # [5, 8, 13, 21, 34, 55]

odd_numbers = filter(lambda x: x % 2 , numbers)
print(list(odd_numbers))   # [5, 7, 9]
```

## 4. Если выбранный язык — объектно-ориентированный, синтаксис определения простых классов, 
какие-то примечательные особенности ООП. В Python класс определяется с помощью ключевого слова 
`class`{.python}, за которым следует имя и двоеточие.

## Опеределение
```python
class Professor():
    pass
```

## Жизненый цикл класса
При создании экземпляра класса, первым срабатывает метод `__new__`{.python}, 
после него – `__init__`{.python}. Метод `__new__`{.python} редко переопределяют, но он 
может быть полезен, например, для реализации синглтонов. В `__init__`{.python} обычно 
инициализируются атрибуты экземпляра.
```python
class Professor():
    def __new__(cls): # cls - ссылка на класс
        print("Calling __new__")
        return super().__new__(cls)

    def __init__(self, name, surname):
        print("Calling __init__")
        self.name = name
        self.surname = surname

prof = Professor("Vladimir", "Konovalov")
# Вывод:
# Calling __new__
# Calling __init__
```
## Инкапсуляция
Python поддерживает базовые механизмы инкапсуляции, хотя строгих модификаторов доступа 
(как private или protected в других языках) нет. Атрибуты и методы могут быть:

- Публичными (доступны отовсюду),
- Защищёнными (начинаются с `_`{.python}, это соглашение о том, что 
они предназначены для внутреннего использования),
- Приватными (начинаются с `__`{.python}, создают обфускацию имени, что затрудняет доступ извне).

```python
class Professor:
    def __init__(self, name, surname):
        self.name = name  
        self.surname = surname
        self._department = "IU9"
        self.__salary = 1000000000

    def teach(self):
        print(f"{self.surname} starts the lesson")

    def __estimate_rk(self):
        print(f"{self.surname} estimates RK by maximum points")

prof = Professor("Vladimir", "Konovalov")
print(prof.name)              # Публичный: работает
print(prof._department)       # Защищённый: работает, но использовать не рекомендуется
print(prof.__salary)          # AttributeError
```
Однако в Python даже к приватными атрибутам можно получить доступ.
```python
print(prof._Professor__salary) # Вывод: 1000000000
```

## Наследование
```python
class Person:
    def __init__(self, name, surname):
        self.name = name
        self.surname = surname

    def introduce(self):
        print(f"My name is {self.name} {self.surname}.")

class Professor(Person):
    def __init__(self, name, surname, university):
        super().__init__(name, surname)
        self.university = university

    def introduce(self):
        print(f"My name is Prof. {self.surname}, I work in {self.university}.")

prof = Professor("Vladimir", "Konovalov", "BMSTU")
prof.introduce() # My name is Prof. Konovalov, I work in BMSTU.
```
Функция `super()`{.python} используется для вызова методов родительского класса. 
Обычно применяется в переопределённых методах для вызова реализации родителя.

## Полиморфизм
```python
class Person:
    def introduce(self):
        print("I am a person.")

class Professor(Person):
    def introduce(self):
        print("I am a professor.")

class Student(Person):
    def introduce(self):
        print("I am a student.")

people = [Person(), Professor(), Student()]

for person in people:
    person.introduce()

# Вывод:
# I am a person.
# I am a professor.
# I am a student.
```

## Абстракция
```python
from abc import ABC, abstractmethod

class Professor(ABC):
    @abstractmethod
    def evaluate(self):
        pass

class CSProfessor(Professor): 
    def evaluate(self):
        print("Проверка лабораторных работ.")

class MathProfessor(Professor):
    def evaluate(self):
        print("Проверка контрольных по математике.")

CSProfessor().evaluate()
MathProfessor().evaluate()
```
## Примечательные особенности Python ООП

1) Динамическое добавление атрибутов  
Атрибуты могут быть добавлены или изменены после создания объекта.
```python
prof.specialization = "Scheme"
print(prof.specialization) 
```

2) Декораторы классов и методов 
Декораторы позволяют изменять или добавлять функциональность к классам и методам.
```python
def add_repr(cls):
    cls.__repr__ = lambda self: f"{cls.__name__}({self.__dict__})"
    return cls

@add_repr
class MyClass:
    def __init__(self, x):
        self.x = x

obj = MyClass(42)
print(obj)  # MyClass({'x': 42})
```

3) Утиная типизация (Duck Typing)
Python использует подход, при котором неважно, к какому классу принадлежит объект. 
Главное — какие методы и свойства он поддерживает.
```python
class Duck:
    def quack(self):
        print("Quack!")

class Person:
    def quack(self):
        print("I can quack too!")

Duck().quack()    # Quack!
Person().quack()  # I can quack too!
```

## 5. Важнейшие функции для работы с потоками ввода/вывода, строками, регулярными выражениями.

## Работа с потоками ввода/вывода  ##

Python предоставляет удобный интерфейс для работы с файлами и потоками. Основные функции:

- Открытие файлов: `open(file, mode)`{.python}
- Режимы: `r` (чтение), `w` (запись), `a` (добавление), `b` (двоичный режим).
- Чтение данных:
- `read(size)`{.js}  - чтение всего содержимого или указанного количества символов.
- `readline()`{.js}  - чтение одной строки.
- `readlines()`{.js}  - чтение всех строк файла.
- Запись данных: `write(data)`{.js} — запись строки или байтов.
- Закрытие файла: `close()`{.js} или использование контекстного менеджера `with`{.python}.

Пример: чтение и запись файла
```python
with open('file.txt', 'r') as file:
    print(file.read())

with open('file.txt', 'w') as file:
    file.write("Hello, world!\nThis is a test file.")
```

## Работа со строками  ##

Python предоставляет обширный набор методов для обработки строк:

- Модификация строк: `lower()`{.js}, `upper()`{.js}, `capitalize()`{.js}, 
`title()`{.js}, `strip()`{.js}, `replace(old, new)`{.js}.
- Проверка содержимого: `startswith(prefix)`{.js}, `endswith(suffix)`{.js}, `isalnum()`{.js}, 
`isalpha()`{.js}, `isdigit()`{.js}.
- Разделение и объединение: `split(delimiter)`{.js}, `join(iterable)`{.js}.
- Форматирование строк: 
    - Старый стиль: `"%s %d" % ("Hello", 123)`{.python}.
    - Новый стиль: `"{} {}".format("Hello", 123)`{.python}.
    - f-строки: `f"{variable}"`{.python}.

Пример: работа со строками
```python
text = "Python is Awesome!"
clean_text = text.strip().lower()
print(clean_text)  # "python is awesome!"
```

## Регулярные выражения ##

Регулярные выражения используются для поиска и обработки текста c помощью модуля `re`. 
Основные функции:

- `match(pattern, string)`{.js} — проверка, начинается ли строка с шаблона.
- `search(pattern, string)`{.js} — поиск первого совпадения в строке.
- `findall(pattern, string)`{.js} — поиск всех совпадений.
- `sub(pattern, repl, string)`{.js} — замена по шаблону.
- `compile(pattern)`{.js} — компиляция шаблона для повторного использования.

Пример регулярных выражений:
```python
import re

text = "Email: example@mail.com, phone: +123456789"
email = re.search(r'\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b', text)
print(email.group()) 
cleaned_text = re.sub(r'\+\d+', '[hidden]', text)
print(cleaned_text)
```