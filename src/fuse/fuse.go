package fuse

import "time"

type FuseStruct struct {
	Number        int
	SuccessNumber int
	ErrNumber     int
	MaxErrNumber  int
	HalfNumber    int
	Status        bool
	FuseStatus    int
	TimeCycle     int
	LastTime      time.Time
	CloseTime     int
	GoStatus      bool
}

var FushMap map[string]*FuseStruct

func init() {
	FushMap = make(map[string]*FuseStruct, 0)
	go run()
}

func run() {
	for {
		handle()
		time.Sleep(time.Second)
	}
}

func handle() {
	for _, v := range FushMap {
		if !v.GoStatus {
			go handleItem(v)
		}
		v.GoStatus = false
	}
}

func handleItem(fush *FuseStruct) {
	for {
		fush.GoStatus = true
		if fush.FuseStatus == 1 {
			if fush.ErrNumber > fush.MaxErrNumber {
				fush.Status = false
				fush.FuseStatus = 3
				fush.LastTime = time.Now().Add(time.Second * time.Duration(fush.CloseTime))
			} else {
				fush.ErrNumber = 0
				fush.Number = 0
			}
		}

		if fush.FuseStatus == 2 {
			if fush.Number > fush.HalfNumber {
				if fush.ErrNumber > fush.SuccessNumber {
					fush.Status = false
					fush.FuseStatus = 3
					fush.LastTime = time.Now().Add(time.Second * time.Duration(fush.CloseTime))
				} else {
					fush.FuseStatus = 1
					fush.Status = true
				}
			}
		}

		if fush.FuseStatus == 3 {
			if time.Now().After(fush.LastTime) {
				fush.ErrNumber = 0
				fush.Number = 0
				fush.FuseStatus = 2
				fush.Status = true

			}
		}

		time.Sleep(time.Second * time.Duration(fush.TimeCycle))

	}

}

func CreateFuse(name string) *FuseStruct {
	if item, ok := FushMap[name]; ok {
		return item
	} else {
		var fush FuseStruct
		fush.MaxErrNumber = 50
		fush.HalfNumber = 100
		fush.SuccessNumber = 10
		fush.ErrNumber = 0
		fush.Status = true
		fush.FuseStatus = 1
		fush.TimeCycle = 1
		fush.CloseTime = 5
		FushMap[name] = &fush
		return &fush
	}

}

func FushStatus(name string) bool {
	fush := CreateFuse(name)
	fush.Number++
	return fush.Status
}
