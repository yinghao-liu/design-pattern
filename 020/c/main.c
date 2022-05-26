
// 一个有限状态机，通常用for+switch-case实现
enum State
{
    StateA = 1,
    StateB,
    StateC
};
enum State g_state;

int main()
{
    for (;;)
    {
        switch (g_state)
        {
        case StateA:
            // dosomething
            // change state
            break;
        case StateB:
            // dosomething
            // change state
            break;
        case StateC:
            // dosomething
            // change state
            break;
        default:
            break;
        }
    }
}