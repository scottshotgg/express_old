#include <map>
#include <string>
#include <iostream>

using namespace std;

// struct Any { std::string type; void* data; };
    // std::unordered_map<std::type_index, std::string> type_names;
 
    // type_names[std::type_index(typeid(int))] = "int";
    // type_names[std::type_index(typeid(double))] = "double";

    enum AnyType { 
        pointerType, 
        intType, 
        boolType,
        charType,
        floatType,
        stringType,
        structType,
        objectType
    };

class Any {
    // TODO: I'm sure we'll need to add more later on

    private:
        AnyType type;
        void* data;
    
    public:
        Any() {}
        Any(void* value) : type(pointerType), data(value) {}
        Any(int value) : type(intType), data(new int(value)) {}
        Any(bool value) : type(boolType), data(new bool(value)) {}
        Any(char value) : type(charType), data(new char(value)) {}
        Any(float value) : type(floatType), data(new float(value)) {}
        Any(string value) : type(stringType), data(new string(value)) {}
        // TODO: will have to do something special here, maybe code generation?
        // Any(struct value) : type(structType), data(&value) {}
        // Any(map<string, Any> value) : type(objectType), data(value) {}
        // might take this out, kind of unsafe
        Any(AnyType iType, void* iData) : type(iType), data(iData) {}

        AnyType Type(void) const {
            return type;
        }

        void* Value(void) const {
            return data;
        }
};

// map[string]func

ostream& operator<<(ostream& stream, const Any& right){
    // TODO: for right now, instead of doing the map[string]function to figure out the value
    // https://stackoverflow.com/questions/4972795/how-do-i-typecast-with-type-info
    // https://stackoverflow.com/questions/2136998/using-a-stl-map-of-function-pointers
    // cout << "me" <<  << endl;
    // cout << "me" << right.Value() << endl;
    // cout << "me" << (int*)right.Value() << endl;
    // cout << "me" << *(int*)right.Value() << endl;
    
    switch(right.Type()) {
        case intType:
            stream << *(int*)right.Value();
        case stringType:
            stream << *(string*)right.Value();
    }

    // cout << *(int*)right.Value() << endl;
    return stream;
}

// Interger operations
int operator+(const int left, const Any& right) {
    return left + *(int*)right.Value();
}

Any operator+(const Any& left, const int right) {
    return Any{ right + left };
}

int main() {
    // string meString = "hey its me";
    Any me = Any { 7 };
    cout << "meResult " << me+7 << endl;
}