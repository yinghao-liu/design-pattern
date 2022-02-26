
#define CALLBACK_MAX 10

#define SUCCESS 0
#define ALREADY_REGISTERED -1
#define EXCEED_MAX -2
#define NOT_FOND -3

// 订阅者需要实现
typedef void (*callback)(int flag);


/******* 发布者*********/
static callback g_callback[CALLBACK_MAX] = {NULL};           


int register_callback(const callback cb)
{
	for (int i = 0; i < CALLBACK_MAX; i++)
	{
		if (NULL == g_callback[i])
		{
			g_callback[i] = cb;
			return SUCCESS;
		}
		else if (cb == g_callback[i])
		{
			printf("callback:%p is same as %u\n", cb, i);
			return ALREADY_REGISTERED;
		}
	}	
	return EXCEED_MAX;
}
	
int unregister_callback(const callback cb)
{
	u16 i = 0;
	for (int i = 0; i < CALLBACK_MAX; i++)
	{
		if (cb == g_callback[i])
		{
			printf("find callback:%p is %u\n", cb, i);
			g_callback[i] = NULL;
			return SUCCESS;
		}
	}
	return NOT_FOND;
}