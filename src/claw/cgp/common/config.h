//Author: sheppard(ysf1026@gmail.com) 2013-08-06
//      Desc:

#ifndef CLAW_CGP_COMMON_CONFIG_H_
#define CLAW_CGP_COMMON_CONFIG_H_

#include <string>

namespace claw
{
namespace cgp
{

class CommonConfig
{
public:
    bool Init();

    std::string login_ip;
    std::string login_port;
    std::string hall_ip;
    std::string hall_port;
};

class Config
{
public:
    bool Init();

    CommonConfig common;
};

extern Config g_config;

} //namespace cgp
} //namespace claw

#endif //CLAW_CGP_COMMON_CONFIG_H_
