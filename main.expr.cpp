#include <map>
#include <string>
struct Any { std::string type; void* data; };
int main() {
std::map<std::string, Any>obj;
std::string thing4176209394 = "hey its me";
obj["thing"] = Any{ "string", &thing4176209394 };
int a1308182024 = 5;
obj["a"] = Any{ "int", &a1308182024 };
float b1143115524 = 67.8;
obj["b"] = Any{ "float", &b1143115524 };
int bb = 11;
int arrayBrah[] = { 5, 5 };
float thing = 77;
int me3923516135 = 5;
Any me = Any{ "int", &me3923516135};
std::string hey = "its me";
}