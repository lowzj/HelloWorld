#include <iostream>
#include <string>
#include <dlfcn.h>
#include <stdlib.h>
#include <unistd.h>

using namespace std;

int main(int argc, char* argv[]) {
  void* handle;
  char* error;

  typedef int(*PRINT)(const string&);
  PRINT p;

  // int flags = RTLD_NOW | RTLD_GLOBAL;
  int flags = RTLD_LAZY;
  if ( !(handle=dlopen("libdlexamp.so", flags)) ) {
    cerr << dlerror() << endl;
    return EXIT_FAILURE;
  }

  *((void**)&p) = dlsym(handle, "print");
  if ( (error=dlerror()) != NULL ) {
    cerr << error << endl;
    return EXIT_FAILURE;
  }

  while (true) {
    p("hello");
    sleep(1);
  }
  dlclose(handle);

  return 0;
}

