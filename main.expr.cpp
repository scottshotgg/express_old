#include <map>
#include <string>
#include <iostream>


struct Any { std::string type; void* data; };
int main() {
int me1018712470 = 5;
Any me = Any{ "int", &me1018712470 };
float mel1018712470 = 5.99;
Any mel = Any{ "float", &mel1018712470 };
std::string hey = "its me";
std::map<std::string, Any>obj;
std::map<std::string, Any>anotherOne;
int a3940063584 = 7;
anotherOne["a"] = Any{ "int", &a3940063584 };
int arrayMeUpScotty[] = { 7, 8, 9 };
bool notMe3108279319 = false;
obj["notMe"] = Any{ "bool", &notMe3108279319 };
obj["anotherOne"] = Any{ "object", &anotherOne };
obj["arrayMeUpScotty"] = Any{ "array", &arrayMeUpScotty };
int a1490497784 = 5;
obj["a"] = Any{ "int", &a1490497784 };
float b673520425 = 67.8;
obj["b"] = Any{ "float", &b673520425 };
std::string thing993103617 = "hey its me";
obj["thing"] = Any{ "string", &thing993103617 };
int bb = 11;
int arrayBrah[] = { 5, 5 };
float thing = 77;
bool thingBool = false;
int hedy = 1;
int* hedyp = &hedy;
std::cout << thing + *(float*)mel.data << std::endl;
std::cout << mel.data << std::endl;
}