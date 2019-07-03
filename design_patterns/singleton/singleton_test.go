package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance() //创建第一个实例

	if counter1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}
	expectedCounter := counter1 //保存第一个实例

	counter1.AddOne() //第一个实例调用方法
	currentCount := counter1.GetCount()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance() //创建第二个实例

	println(expectedCounter, counter2)
	if counter2 != expectedCounter { //比较第一个和第二个实例
		t.Error("expected same instance int counter2 but it got a different instance")
	}

	counter2.AddOne()
	currentCount = counter1.GetCount()
	if currentCount != 2 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}
}

func TestGetInstanceConcurrent(t *testing.T) {
	// runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("instance AddOne i=%d", i), func(t *testing.T) {
			obj := GetInstance()
			obj.AddOne()
			t.Logf("GetCount(%d) \n",
				GetInstance().GetCount(),
			)
		})
	}
}

func BenchmarkGetInstanceConcurrent(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.Run(fmt.Sprintf("instance AddOne i=%d", i), func(b *testing.B) {
			obj := GetInstance()
			obj.AddOne()
			b.Logf("GetCount(%d) \n",
				GetInstance().GetCount(),
			)
		})
	}
}

func BenchmarkMutex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			obj := GetInstance()
			obj.AddOne()
			fmt.Printf("GetCount(%d) \n",
				GetInstance().GetCount(),
			)
		}
	})
}
