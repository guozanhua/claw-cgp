//Author: sheppard(ysf1026@gmail.com) 2013-07-14
//      Desc:

#include <iostream>
#include <gtest/gtest.h>

int main(int argc, char** argv)
{
    testing::InitGoogleTest(&argc, argv);
    RUN_ALL_TESTS();
#ifdef _WINDOWS_
    system("pause");
#endif
    return 0;
}
