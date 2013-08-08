//Author: sheppard(ysf1026@gmail.com) 2013-07-16
//      Desc:

#include "claw/cgp/login/login.h"
#include <claw/gse/log.h>
#include <claw/gse/thread.h>
#include <google/protobuf/stubs/common.h>
#include <zmq.hpp>
#include "claw/cgp/common/config.h"
#include "claw/cgp/common/proto/cmd.pb.h"

namespace claw
{
namespace cgp
{

Config g_config;

struct ThreadParam
{
    int thread_id;
    zmq::context_t* context;
};

void worker_routine(void* arg)
{
    ThreadParam* param = (ThreadParam*)arg;
    int thread_id = param->thread_id;
    zmq::context_t* zmq_context = (zmq::context_t*)param->context;

    zmq::socket_t socket(*zmq_context, ZMQ_REP);
    socket.connect("inproc://workers");

    while(true)
    {
        //wait for next request from client
        proto::Cmd cmd;
        zmq::message_t request;
        socket.recv(&request);
        if(false == cmd.ParseFromString(std::string((char*)request.data(), request.size()-1)))
        {
            LOG(ERROR)<<"parse error, id="<<cmd.id()<<" content="<<cmd.content();
        }
        LOG(INFO)<<"thread("<<thread_id<<") received request: [id="<<cmd.id()<<" content="<<cmd.content()<<"]";

        //send reply back to client
        zmq::message_t reply(6);
        memcpy((void*)reply.data(), "world", 6);
        socket.send(reply);
    }
}

} //namespace cgp
} //namespace claw

using namespace claw::cgp;

int main(int argc, char* argv[])
{
    GOOGLE_PROTOBUF_VERIFY_VERSION;

    claw::gse::Log log(argv[0], "../log");

    LOG(INFO)<<"logic_server start...";
    zmq::context_t context(1);
    zmq::socket_t clients(context, ZMQ_ROUTER);
    clients.bind(std::string("tcp://*:").append(g_config.common.login_port).c_str());
    zmq::socket_t workers(context, ZMQ_DEALER);
    workers.bind("inproc://workers");

    ThreadParam threadParam;
    threadParam.context = &context;

    for(int thread_nbr=0; thread_nbr!=5; thread_nbr++)
    {
        threadParam.thread_id = thread_nbr;
		claw::gse::Thread worker;
        worker.Start(worker_routine, (void*)&threadParam);
        LOG(INFO)<<"create thread, turn="<<thread_nbr;
    }
    zmq::proxy(clients, workers, NULL);

    google::protobuf::ShutdownProtobufLibrary();
    return 0;
}
