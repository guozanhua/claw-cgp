//Author: sheppard(ysf1026@gmail.com) 2013-07-16
//      Desc:

#ifndef CLAW_CGP_LOGIN_THREAD_H_
#define CLAW_CGP_LOGIN_THREAD_H_

#ifdef _WINDOWS_
#include <windows.h>
#else
#include <pthread.h>
#endif

namespace claw
{
namespace cgp
{

typedef void (ThreadFunction) (void*);

class Thread
{
public:
    Thread() {}

    void Start(ThreadFunction* function, void* arg);
    void Stop();

    ThreadFunction* thread_function_;
    void* arg_;
        
private:
    Thread(const Thread&);
    const Thread& operator=(const Thread&);

#ifdef _WINDOWS_
    HANDLE descriptor_;
#else
    pthread_t descriptor_;
#endif
};

} //namespace cgp
} //namespace claw

#endif //CLAW_CGP_LOGIN_THREAD_H_
