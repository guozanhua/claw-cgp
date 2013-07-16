//Author: sheppard(ysf1026@gmail.com) 2013-07-16
//      Desc:

#include "claw/cgp/login/login.h"
#include <iostream>
#include <google/protobuf/stubs/common.h>
#include <zmq.hpp>
#include "claw/cgp/common/proto/cmd.pb.h"
#include "claw/cgp/login/thread.h"

namespace claw
{
namespace cgp
{

struct ThreadParam
{
    int thread_id;
    zmq::context_t* context;
};

void* worker_routine(void* arg)
{
    ThreadParam* param = (ThreadParam*)arg;
    int thread_id = param->thread_id;
    zmq::context_t* zmq_context = (zmq::context_t*)param->context;

    zmq::socket_t socket(*zmq_context, ZMQ_REP);
    socket.connect("inproc://workers");

    while(true)
    {
        //Wait for next request from client
        proto::Cmd cmd;
        zmq::message_t request;
        socket.recv(&request);
        if(false == cmd.ParseFromString(std::string((char*)request.data(), request.size()-1)))
        {
            std::cout<<"Parse error, id="<<cmd.id()<<" content="<<cmd.content()<<std::endl;
        }
        std::cout<<"Thread("<<thread_id<<") Received request: [id="<<cmd.id()<<" content="<<cmd.content()<<"]"<< std::endl;

        //Do some 'work'
        //sleep (1);

        //Send reply back to client
        zmq::message_t reply(6);
        memcpy((void*)reply.data(), "World", 6);
        socket.send(reply);
    }

    return NULL;
}

int main(int argc, char* argv[])
{
    GOOGLE_PROTOBUF_VERIFY_VERSION;

    std::cout<<"logic_server start..."<<std::endl;
    zmq::context_t context(1);
    zmq::socket_t clients(context, ZMQ_ROUTER);
    clients.bind("tcp://*:5555");
    zmq::socket_t workers(context, ZMQ_DEALER);
    workers.bind("inproc://workers");

    ThreadParam threadParam;
    threadParam.context = &context;

    for(int thread_nbr=0; thread_nbr!=5; thread_nbr++)
    {
        threadParam.thread_id = thread_nbr;
        Thread worker;
        worker.Start((ThreadFunction)worker_routine, (void*)&threadParam);
        std::cout<<"create thread, turn="<<thread_nbr<<std::endl;
    }
    zmq::proxy(clients, workers, NULL);

    google::protobuf::ShutdownProtobufLibrary();
    return 0;
}

} //namespace cgp
} //namespace claw
