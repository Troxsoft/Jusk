# Jusk
A programing language that compile to c++
*The language is missing many features*

> ### Example for hello world
```rs
@pkg main
@import "@console/console.jk"
pub fn main()Int{
   Console.println("Hello World")
}
```
#### `jusk r hello_world.jk`
### How to install? (require go,g++ installed)
- ## 1 
```bash
git clone https://github.com/Troxsoft/Jusk.git
```
- ## 2
```bash
go build jusk.go
```

- ## 3 *(optional)*
Add to path
> # Progress
>> - [x] Boolean
>> - [x] Single line comment
>> - [x] Multi line comment  
>> # Message Errors
>> - [ ] error with code
>> - [ ] util info
>> - [ ] show line of error
>> - [ ] show package of error
>> - [ ] show filename of error

>>> # Standart Jusk Library
>>> -  Console ⬜⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛ 1%
>>> -  Util ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛ 0%
>>> -  FS ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛ 0%
>>> -  Global functions ⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛⬛ 2%  `Example: toStr`
>> - [x] declare function
>> - [x] call function
>>>  # if
>>> - [x] if
>>> - [x] elif
>>> - [x] else
>> - [x] import package
>> - [x] strings
>> - [x] numbers
>> - [ ] using
>> - [ ] [] for array
>> - [ ] array parse
>> - [ ] array declaration
>> - [x] declare package
>> - [ ] type-check  80%
>>>  # BinaryExpressions:
>>> - [x] +
>>> - [x] -
>>> - [x] /
>>> - [x] *
>>> - [x] ==
>>> - [x] !=
>>> - [x] %
>>> - [x] <
>>> - [x] >
>>> - [x] <=
>>> - [x] >=
>>> - [x] and
>>> - [x] or
>>
>>>  # visibility
>>> - [x] functions
>>> - [ ] global variables
>>>  # C++ 
>>> - [x] @cpp
>>
>> # Data types
>>> # Structs Progress:
>>>  ⬜⬛⬛⬛⬛⬛⬛⬛⬛⬛ 2%

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
