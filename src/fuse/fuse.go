package fuse

import (
	"time"
)

type FuseStruct struct {
	Number           int
	ErrNumber        int
	MaxErrNumber     int
	HalfErrNumber    int
	HalfBucketNumber int
	HalfNumber       int
	Status           bool
	FuseStatus       int
	TimeCycle        int
	LastTime         time.Time
	CloseTime        int
	GoLastTime       time.Time
}

var FuseMap map[string]*FuseStruct

func init() {
	FuseMap = make(map[string]*FuseStruct, 0)
	go run()
}

func run() {
	for {
		handle()
		time.Sleep(time.Second)
	}
}

func handle() {
	for _, v := range FuseMap {
		if !v.GoLastTime.After(time.Now()) {
			go handleItem(v)
		}
	}
}

func handleItem(fuse *FuseStruct) {
	for {
		// sss, _ := json.Marshal(fuse)
		// fmt.Println("sta", string(sss))

		fuse.GoLastTime = time.Now().Add(time.Second * time.Duration(fuse.TimeCycle*2))
		if fuse.FuseStatus == 1 {
			if fuse.ErrNumber >= fuse.MaxErrNumber {
				fuse.Status = false
				fuse.FuseStatus = 3
				fuse.HalfBucketNumber = fuse.HalfNumber
				fuse.LastTime = time.Now().Add(time.Second * time.Duration(fuse.CloseTime))
			} else {
				fuse.ErrNumber = 0
				fuse.Number = 0
			}
		}

		if fuse.FuseStatus == 2 {
			if fuse.Number >= fuse.HalfNumber {
				if fuse.ErrNumber >= fuse.HalfErrNumber {
					fuse.Status = false
					fuse.FuseStatus = 3
					fuse.HalfBucketNumber = fuse.HalfNumber
					fuse.LastTime = time.Now().Add(time.Second * time.Duration(fuse.CloseTime))
				} else {
					fuse.FuseStatus = 1
					fuse.ErrNumber = 0
					fuse.Number = 0
					fuse.Status = true
				}
			}
		}

		if fuse.FuseStatus == 3 {
			if time.Now().After(fuse.LastTime) {
				fuse.ErrNumber = 0
				fuse.Number = 0
				fuse.FuseStatus = 2
				fuse.Status = true

			}
		}
		time.Sleep(time.Second * time.Duration(fuse.TimeCycle))

	}

}

func CreateFuse(name string) *FuseStruct {
	if item, ok := FuseMap[name]; ok {
		return item
	} else {
		var fuse FuseStruct
		fuse.MaxErrNumber = 10
		fuse.HalfNumber = 10
		fuse.HalfBucketNumber = 10
		fuse.HalfErrNumber = 3
		fuse.ErrNumber = 0
		fuse.Status = true
		fuse.FuseStatus = 1
		fuse.TimeCycle = 1
		fuse.CloseTime = 5
		FuseMap[name] = &fuse
		return &fuse
	}

}

func FuseStatus(name string) bool {
	fuse := CreateFuse(name)
	fuse.Number++
	if fuse.FuseStatus == 2 {
		// sss, _ := json.Marshal(fuse)
		// fmt.Println("FuseStatus 2", string(sss))
		if fuse.HalfBucketNumber > 0 {
			fuse.HalfBucketNumber--
		} else {
			return false
		}
	}
	return fuse.Status
}
