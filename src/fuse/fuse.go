package fuse

import (
	"encoding/json"
	"fmt"
	"time"
)

type FuseStruct struct {
	Number            int
	HalfSuccessNumber int
	HalfBucketNumber  int
	ErrNumber         int
	MaxErrNumber      int
	HalfNumber        int
	Status            bool
	FuseStatus        int
	TimeCycle         int
	LastTime          time.Time
	CloseTime         int
	GoStatus          bool
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
		if !v.GoStatus {
			go handleItem(v)
		}
		v.GoStatus = false
	}
}

func handleItem(fuse *FuseStruct) {
	for {
		sss, _ := json.Marshal(fuse)
		fmt.Println("sta", string(sss))

		fuse.GoStatus = true
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
				if fuse.ErrNumber >= fuse.HalfSuccessNumber {
					fuse.Status = false
					fuse.FuseStatus = 3
					fuse.HalfBucketNumber = fuse.HalfNumber
					fuse.LastTime = time.Now().Add(time.Second * time.Duration(fuse.CloseTime))
				} else {
					fuse.FuseStatus = 1
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
		sssend, _ := json.Marshal(fuse)
		fmt.Println("end", string(sssend))
		// fmt.Println("fuse.FuseStatus", fuse.FuseStatus)
		// fmt.Println("fuse.FuseStatus", fuse.FuseStatus)
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
		fuse.HalfSuccessNumber = 3
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
		sss, _ := json.Marshal(fuse)
		fmt.Println("FuseStatus 2", string(sss))
		if fuse.HalfBucketNumber > 0 {
			fuse.HalfBucketNumber--
		} else {
			return false
		}
	}
	return fuse.Status
}
