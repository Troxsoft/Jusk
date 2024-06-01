#include <iostream>
class Str
{
private:
    std::string _str;

public:
    Str() { this->_str = " "; }

public:
    Str(std::string ss)
    {

        this->_str = ss;
    }

public:
    Str(Str other) { this->_str = other._str; }
    int Len() { return this->_str.length(); }
    bool Eq(Str other) { return this->_str == other._str; }
};