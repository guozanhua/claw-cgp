//Author: sheppard(ysf1026@gmail.com) 2013-08-06
//      Desc:

#include "claw/cgp/common/config.h"
#include <fstream>
#include <glog/logging.h>
#include <json/json.h>

namespace claw
{
namespace cgp
{

bool CommonConfig::Init()
{
    std::ifstream stream;
    stream.open("../config/common.json");

    Json::Value root;
    Json::Reader reader;
    if(!reader.parse(stream, root))
    {
        LOG(ERROR)<<reader.getFormatedErrorMessages();
        return false;
    }

    login_ip = root["net"]["login_ip"].asString();
    login_port = root["net"]["login_port"].asString();
    hall_ip = root["net"]["hall_ip"].asString();
    hall_port = root["net"]["hall_port"].asString();

    return true;
}

bool Config::Init()
{
    if(!common.Init())
        return false;
    return true;
}

} //namespace cgp
} //namespace claw
