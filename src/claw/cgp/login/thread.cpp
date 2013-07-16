//Author: sheppard(ysf1026@gmail.com) 2013-07-16
//      Desc:

#include "claw/cgp/login/thread.h"

#ifdef _LINUX_
#include <signal.h>
#endif

#ifdef _WINDOWS_
#include <process.h>
#endif

namespace claw
{
namespace cgp
{

#ifdef _WINDOWS_

extern "C"
{
    static unsigned int __stdcall ThreadRoutine(void* arg)
    {
        Thread* self = (Thread*)arg;
        self->thread_function_(self->arg_);
        return 0;
    }
}

void Thread::Start(ThreadFunction* thread_function, void* arg)
{
    thread_function_ = thread_function;
    arg_ = arg;
    descriptor = (HANDLE)_beginthreadex(NULL, 0, &::ThreadRoutine, this, 0, NULL);
}

void Thread::Stop()
{
    WaitForSingleObject(descriptor, INFINITE);
    CloseHandle(descriptor);
}

#else //_WINDOWS_

extern "C"
{
    static void* ThreadRoutine(void* arg)
    {
        sigset_t signal_set;
        int rc = sigfillset(&signal_set);
        //TODO:
        pthread_sigmask(SIG_BLOCK, &signal_set, NULL);

        Thread *self = (Thread*)arg;   
        self->thread_function_(self->arg_);
        return NULL;
    }
}

void Thread::Start(ThreadFunction* thread_function, void* arg)
{
    thread_function_ = thread_function;
    arg_ = arg_;
    pthread_create(&descriptor, NULL, ThreadRoutine, this);
}

void Thread::Stop()
{
    pthread_join(descriptor, NULL);
}

#endif //_WINDOWS_

} //namespace cgp
} //namespace claw
