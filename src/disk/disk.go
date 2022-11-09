package disk

import (
	"godevil/src/envt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Disk struct {
	Name       string // /sys/block/sda stored as /dev/sda
	RealPath   string
	Removable  bool   // /sys/block/sda/removable
	SectorSize uint64 // /sys/block/sda/queue/hw_sector_size
	Size       uint64 // /sys/block/sda/queue/hw_sector_size * /sys/block/sda/size
}

var DiskList []*Disk

// Gives the location of all block devices in the sysfs
func getDevices() {
	blockdev_loc := filepath.Join(envt.SysFS, "block")
	err := filepath.Walk(blockdev_loc, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		l_path := strings.Split(blockdev_loc, "/")
		if l_path[len(l_path)-1] != info.Name() {
			devname := filepath.Join(envt.DevFS, info.Name())
			if _, e := os.Stat(devname); e == nil {
				realpath, err := filepath.EvalSymlinks(filepath.Join(blockdev_loc, info.Name()))
				if err != nil {
					panic(err)
				}
				disk := &Disk{
					Name:     devname,
					RealPath: realpath,
				}
				DiskList = append(DiskList, disk)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (d *Disk) getDeviceSector() {
	realpath := d.RealPath
	var sector_size uint64 = 512
	sectorfile := filepath.Join(realpath, "queue", "hw_sector_size")

	if fd, e := os.Open(sectorfile); e == nil {
		defer fd.Close()
		content, err := ioutil.ReadFile(sectorfile)
		if err != nil {
			panic(err)
		}
		str_content := strings.Trim(string(content), "\n")
		sector_size, err = strconv.ParseUint(str_content, 10, 0)
		if err != nil {
			panic(err)
		}
	}
	d.SectorSize = sector_size
}

func (d *Disk) getDeviceNumSectors() {
	realpath := d.RealPath
	var num_sectors uint64 = 0
	sizefile := filepath.Join(realpath, "size")

	if fd, e := os.Open(sizefile); e == nil {
		defer fd.Close()
		content, err := ioutil.ReadFile(sizefile)
		if err != nil {
			panic(err)
		}

		str_content := strings.Trim(string(content), "\n")
		num_sectors, err = strconv.ParseUint(str_content, 10, 0)
		if err != nil {
			panic(err)
		}
	}
	var disk_size uint64 = num_sectors * d.SectorSize / 1024
	d.Size = disk_size
}

func (d *Disk) getDeviceRemovable() {
	realpath := d.RealPath
	var removable bool = false
	removable_file := filepath.Join(realpath, "removable")
	if fd, e := os.Open(removable_file); e == nil {
		defer fd.Close()
		content, err := ioutil.ReadFile(removable_file)
		if err != nil {
			panic(err)
		}
		str_content := strings.Trim(string(content), "\n")
		if i, err := strconv.Atoi(str_content); err == nil {
			if i == 1 {
				removable = true
			}
		}
	}
	d.Removable = removable
}

func populateDeviceProp() {
	getDevices()
	for _, v := range DiskList {
		v.getDeviceSector()
		v.getDeviceNumSectors()
		v.getDeviceRemovable()
	}
}

func Get() []*Disk {
	if DiskList == nil {
		populateDeviceProp()
	}
	return DiskList
}
