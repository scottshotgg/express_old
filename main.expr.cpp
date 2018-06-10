#include <map>
#include <string>
struct Any { std::string type; void* data; };
int main() {
int arrayBrah[] = { 5, 5 };
float thing = 77;
int me3035228544 = 5;
Any me = Any{ "int", &me3035228544};
std::string hey = "its me";
std::map<std::string, Any>obj;
std::map<std::string, Any>anotherOne;
int a1275390956 = 7;
anotherOne["a"] = Any{ "int", &a1275390956 };
int arrayMeUpScotty[] = { 7, 8, 9 };
int a1537738808 = 5;
obj["a"] = Any{ "int", &a1537738808 };
float b2045175289 = 67.8;
obj["b"] = Any{ "float", &b2045175289 };
std::string thing2493503111 = "hey its me";
obj["thing"] = Any{ "string", &thing2493503111 };
bool notMe1441537250 = false;
obj["notMe"] = Any{ "bool", &notMe1441537250 };
obj["anotherOne"] = Any{ "object", &anotherOne };
obj["arrayMeUpScotty"] = Any{ "array", &arrayMeUpScotty };
int bb = 11;
}