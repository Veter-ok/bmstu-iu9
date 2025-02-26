# Лабораторная работа № 4. Метапрограммирование. Отложенные вычисления

## 1. Продолжения.
Утверждение `(assertion)` — проверка на истинность некоторого условия, заданного программистом. По традиции осуществляется процедурой (функцией) с именем `assert`. Включается в код во время написания кода и отладки с целью установки ограничений на значения и выявления недопустимых значений. Если в процессе выполнения программы указанное условие нарушается, то программа завершается с выводом диагностического сообщения о том, какое условие было нарушено. Если условие соблюдено, то выполнение программы продолжается, никаких сообщений не выводится.

Реализуйте каркас (фреймворк) для отладки с помощью утверждений. Пусть Ваш каркас перед использованием инициализируется вызовом `(use-assertions)`, а сами утверждения записываются в коде ниже в виде `(assert условие)`. Если условие не выполнено, происходит завершение работы программы без возникновения ошибки выполнения и вывод в консоль диагностического сообщения вида `FAILED: условие`. Пример использования каркаса:

```scheme
(use-assertions) ; Инициализация вашего каркаса перед использованием

; Определение процедуры, требующей верификации переданного ей значения:

(define (1/x x)
  (assert (not (zero? x))) ; Утверждение: x ДОЛЖЕН БЫТЬ ≠ 0
  (/ 1 x))

; Применение процедуры с утверждением:

(map 1/x '(1 2 3 4 5)) ; ВЕРНЕТ список значений в программу

(map 1/x '(-2 -1 0 1 2)) ; ВЫВЕДЕТ в консоль сообщение и завершит работу программы
```
Сообщение, которое должно быть выведено при выполнении примера, показанного выше:
```sheme
FAILED: (not (zero? x))
```
**Важно!** Если в программе используются гигиенические макросы и эта программа будет выполнена в среде `guile 1.8.x` (в том числе на сервере тестирования), то следует подключить модуль поддержки таких макросов, написав в начале программы следующую строку: `(use-syntax (ice-9 syncase))`

## 2. Код как данные. Порты ввода-вывода.
Поиск рекурсивных процедур. Напишите процедуру `(select-rec source dest)`, которая принимает два параметра: имя исходного файла на языке Scheme `source` и имя целевого файла `dest`, куда следует выписать только рекурсивные процедуры из первого файла.

Для простоты рассматриваем только глобальные процедуры, процедуру считаем рекурсивной, если в её теле встречается её имя. Рассмотрение сложных случаев (взаимная рекурсия, рекурсия во вложенных процедурах и т.д.) выполнять не нужно.

Статистика текста. Напишите процедуру `(text-stat source)`, которая принимает текстовый файл и подсчитывает в нём количество слов (состоят из букв, дефисов и апострофов), чисел (состят из цифр) и формул (состоят из букв, цифр и математических знаков, например, E=mc2 или H2SO4).

## 3. Мемоизация результатов вычислений.
Реализуйте функцию вычисления n-го “числа трибоначчи” (последовательности чисел, которой первые три числа равны соответственно 0, 0 и 1, а каждое последующее число — сумме предыдущих трех чисел):

Функция
$$
t(n) = 
\begin{cases} 
0, \space n <= 1 \\
1, \space n = 2 \\
t(n - 1) + t(n - 2) + t(n - 3), \space  n > 2
\end{cases}
$$

Область определения функции
Реализуйте версию этой функции с мемоизацией результатов вычислений. Сравните время вычисления значения функций для разных (умеренно больших) значений её аргументов без мемоизации и с мемоизацией. Для точного измерения вычисления рекомендуется использовать команду `REPL Guile time (Guile 2.x)`.

## 4. Отложенные вычисления
Используя средства языка Scheme для отложенных вычислений, реализуйте средства для работы с бесконечными «ленивыми» точечными парами и списками:

Гигиенический макрос `(lazy-cons a b)`, конструирующий ленивую точечную пару вида (значение-a . обещание-вычислить-значение-b). Почему макрос в данном случае предпочтительнее процедуры?  
Процедуру `(lazy-car p)`, возвращающую значение 1-го элемента «ленивой» точечной пары `p`.  
Процедуру `(lazy-cdr p)`, возвращающую значение 2-го элемента «ленивой» точечной пары `p`.

На их основе определите:  
- Процедуру `(lazy-head xs k)`, возвращающую значение k первых элементов «ленивого» списка `xs` в виде списка.
- Процедуру `(lazy-ref xs k)`, возвращающую значение `k`-го элементa «ленивого» списка `xs`.
- Процедуру `(lazy-map proc xs)`, принимающую «ленивый» список xs и формирующую новый «ленивый» список, полученный применением `(proc x)` к каждому `x` из `xs`.
- Процедуру `(lazy-zip xs ys)`, принимающую два «ленивых» списка `xs` и `ys` и формирующую новый «ленивый» список, составленный из списков `(list x y)` для каждых `x` и `y` из `xs` и `ys` соответственно. Поведение `(lazy-zip xs ys) `эквивалентно `(map list xs ys)` с тем отличием, что списки «ленивые»
Бесконечный список, состоящий из единиц, будет описан как:
```scheme
(define ones (lazy-cons 1 ones))
```
Проверка:
```scheme
> (lazy-head ones 5)
(1 1 1 1 1)
```
Для сравнения, бесконечный список единиц на Хаскеле описывается как (см. лекцию):
```Haskell
ones = 1 : ones
```
Проверка:
```Haskell
> take 5 ones
[1,1,1,1,1]
```
«Ленивый» список факториалов. «Переведите» следующий код с Хаскеля на Scheme:
```Haskell
-- это функция: (define naturals (lambda (start) ...))
naturals = \start -> start : (naturals (start + 1))
factorials = 1 : map (\(n, f) -> n * f) (zip (naturals 1) factorials)
```
Используйте примитивы `lazy-cons`, `lazy-map` и `lazy-zip`, определённые ранее.

## 5. Управляющие конструкции.
Используя гигиенические макросы языка Scheme, реализуйте управляющие конструкции, свойственные императивным языкам программирования.

Управляющие конструкции

### 1. Условие unless
Напишите макрос: `(unless cond? expr1 expr2 … exprn)`, который выполняет последовательность выражений `expr1 expr2 … exprn`, если условие `cond?` ложно.
Предполагается, что `unless `возвращает результат последнего вычисленного в нём выражения. `Unless` может быть вложенным.

Пример:
```scheme
; Пусть x = 1
;
(unless (= x 0) (display "x != 0") (newline))
```
В стандартный поток будет выведено:
```scheme
x != 0
```
### 2. Цикл while
Реализуйте макрос `while`, который позволит организовывать циклы с предусловием: `(while cond? expr1 expr2 … exprn)`, где `cond?` — условие, `expr1 expr2 … exprn` — последовательность инструкций, которые должны быть выполнены в теле цикла. Проверка условия осуществляется перед каждой итерацией, тело цикла выполняется, если условие выполняется. Если при входе в цикл условие не выполняется, то тело цикла не будет выполнено ни разу.

Пример применения:
```scheme
(let ((p 0)
      (q 0))
  (while (< p 3)
         (set! q 0)
         (while (< q 3)
                (display (list p q))
                (newline)
                (set! q (+ q 1)))
         (set! p (+ p 1))))
```
Выведет:
```bash
(0 0)
(0 1)
(0 2)
(1 0)
(1 1)
(1 2)
(2 0)
(2 1)
(2 2)
```
**Рекомендация**. Целесообразно разворачивать макрос в вызов анонимной процедуры без аргументов со статической переменной, содержащей анонимную процедуру с проверкой условия, рекурсивным вызовом и телом цикла. Для краткой записи такой процедуры и ее вызова можно использовать встроенную конструкцию `letrec`, которая аналогична `let` и `_let*_`, но допускает рекурсивные определения, например:
```scheme
(letrec ((iter (lambda (i)
                 (if (= i 10)
                     '()
                     (cons i (iter (+ i 1)))))))
  (iter 0))
  => (0 1 2 3 4 5 6 7 8 9)
```
### 3. Вывод «в стиле С++»
Реализуйте макрос для последовательного вывода значений в стандартный поток вывода вида:

`(cout << "a = " << 1 << endl << "b = " << 2 << endl)`
Здесь `cout` — имя макроса, указывающее, что будет осуществляться вывод в консоль `(от console output)`, символы `<<` разделяют значения, `endl` означает переход на новую строку.

Данный пример выведет следующий текст:
```bash
a = 1
b = 2
```