
### Why will it deadlock if a goroutine acquire a mutex while pinned to its P?
https://groups.google.com/forum/#!topic/golang-nuts/zkkTTEj85lo

##golang 标准库sync提供Mutex、RWMutex,使用起来并不复杂,但有几个地方需要注意:
ref:http://www.pydevops.com/2016/11/21/go-mutex-%E6%BA%90%E7%A0%81%E5%89%96%E6%9E%90/

1. Mutex可以作为结构体的一部分
2. Mutex创建后,以后对Mutex的操作不能复制Mutex,必须实现为pointer-receiver,否则会因复制的关系,导致锁机制失效

 state: state是一个int32类型,由32个二进制位组成,该state被切成3部分使用(通过const 三个常量可以得知)。
 0位表示是否上锁, 1位表示清醒状态,2~31位表示等待队列等待计数。
 如果当前该锁处于未上锁状态,则优先从清醒状态获取任务(从等待队列唤醒一个任务,需要付出的代价是很大的)。
 
 sema: 信号,用于向休眠队列发送锁释放信号,处于自旋状态的g优先可以获得该信号。

 ###mutex

 ```go

 type Mutex struct {
        state int32 		
        sema  uint32  		// 向休眠状态队列发送锁释放信号,处于自旋状态的G最先拿到该信号
}

 const (
	mutexLocked = 1 << iota // mutex is locked  //1
	mutexWoken//2
	mutexStarving//4
	mutexWaiterShift = iota//3
)

 func (m *Mutex) Lock() {
        // 尝试直接获取锁
        if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) { 
                return
        }

        awoke := false		// 自旋标记
        iter := 0		// 自旋计数器
        for {
        	// 获取当前锁的状态
                old := m.state

                // 添加锁状态
                // 0000 | 0001  ==> 0001
                // 00 -休眠队列  
                // 0  -自旋状态 
                // 1  -锁状态
                new := old | mutexLocked

                // 如果锁已经被占用
                // 0 -unlock 
                // 1 -locked
                if old&mutexLocked != 0 {

                	// 尝试进入自旋状态,以便快速切换
                	// 自旋有次数等一系列限制,并非一致循环
                	// 在自旋状态并没有进入休眠状态
                	// for循环进入cpu 空耗状态
                	// awoke标记,自旋的一个标记.当锁被释放后,拿到该标记的g的优先级最高.可以快速获取到锁
                	// 检查是否可以进入自旋状态
                        if runtime_canSpin(iter) {
                                 
                                if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
                                        atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {

                                        // 获取到自旋状态,将自选标记awoke标记为true
                                        awoke = true   
                                }

                                // 进入自旋状态
                                runtime_doSpin()

                                // 尝试多次自旋(自旋有次数限制)
                                iter++
                                continue
                        }

                        //  如果在自旋状态没有获取到锁,则该g进入休眠状态.计数器增加
                        // 累加计数器
                        // new 表示新的state状态
                        new = old + 1<<mutexWaiterShift
                }

                // 如果在自旋状态未获取到锁,进入休眠状态。
                // 需要将自旋状态标记清除
                // 清除清醒标记
                if awoke {
                        if new&mutexWoken == 0 {
                                panic("sync: inconsistent mutex state")
                        }

                        // new = new &^ mutexWoken
                        // new = 4 &^ 2 ==> 0100 &^ 0010 ==> 0000
                        // 清除清醒标记
                        new &^= mutexWoken	
                }

                // 更新状态: (两种可能)
                // 1. 请求锁失败: 自旋结束,更新等待计数
                // 1. 请求锁成功: 添加锁标记
                if atomic.CompareAndSwapInt32(&m.state, old, new) {

                	// 判断是否请求锁成功
                	// 0 表示锁请求成功, 1 表示锁请求失败
                	// old&mutexLocked => 0 & 1 == 0
                	// old&mutexLocked => 1 & 1 == 1
                        if old&mutexLocked == 0 {
                                break
                        }

                        // 锁请求失败,进入休眠状态,等待信号唤醒后重新开始循环
                        runtime_Semacquire(&m.sema)
                        awoke = true
                        iter = 0
                }
        }
}

加锁一共可以分为以下几种情况:
       1. 如果直接获取锁成功,则直接退出Lock函数。
       2. 如果直接获取锁失败,for循环中再次请求锁状态,如果锁未被占用,则再次获取锁。如果此时获取锁成功则退出Lock函数;如果获取锁失败,则将任务g丢到休眠队列,自旋状态标记为true,自旋计数器初始化为0,然后进入新一轮for循环。
       3. 如果直接获取锁失败,for循环中再次请求锁状态,如果锁已经被占用,则检查是否可以进入自旋状态(清醒状态)。如果可以获取,则进入自旋状态。在自旋状态期间如果锁被释放了,则再次获取锁。如果获取锁成功,则退出Lock函数;否则丢进休眠队列,进入下一轮for循环。
       4. 如果直接获取锁失败,for循环中再次请求锁状态,如果锁已经被占用,则检查是否可以进入自旋状态(清醒状态)。如果可以获取并且自旋计数阈值内没有获取到锁,则需要退出自旋状态。然后将该g标记为休眠状态,并将等待队列计数器加1。接着清除自旋标记,然后state状态通过CAS进行更新,最后进入休眠队列,进入下一轮for循环。

       // 不能对未上锁的mutex解锁
// 锁操作与goroutine无关,加解锁操作可分别由不同的G完成
func (m *Mutex) Unlock() {

        // 原子操作
        // 移除锁定标记
        new := atomic.AddInt32(&m.state, -mutexLocked)

        old := new
        for {
                // 当休眠队列内的等待计数为0或者自旋状态计数器为0
                // 退出
                if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken) != 0 {
                        return
                }
                
                // 通过信号唤醒某个等待者
                // 须减少等待计数,并添加清醒标记
                new = (old - 1<<mutexWaiterShift) | mutexWoken
                if atomic.CompareAndSwapInt32(&m.state, old, new) {
                	
                        // 释放锁,发送释放信号
                        runtime_Semrelease(&m.sema)
                        return
                }
                old = m.state
        }
}

 ```