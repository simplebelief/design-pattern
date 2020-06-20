package main

type SubjectImpl struct {
	ObjList []ObjInterface
	message string
}

//订阅主题
func (subi *SubjectImpl) Register(obj ObjInterface) {

	subi.ObjList = append(subi.ObjList, obj)

}

//取消订阅
func (subi *SubjectImpl) DelObserver(obj ObjInterface) {

	for i, _ := range subi.ObjList {
		if subi.ObjList[i] == obj {
			if i == 0 {
				subi.ObjList = subi.ObjList[i+1:]
			} else if i == len(subi.ObjList)-1 {
				subi.ObjList = subi.ObjList[:i]
			} else {
				subi.ObjList = append(subi.ObjList[:i], subi.ObjList[i+1])
			}

			break
		}
	}

}

//通知观察者
func (subi *SubjectImpl) NotifyOberservers() {

	for _, observer := range subi.ObjList {

		observer.Read(subi.message)
	}

}

func (subi *SubjectImpl) Update(message string) {
	subi.message = message

	//发布消息
	subi.NotifyOberservers()
}
