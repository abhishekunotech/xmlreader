package main

import(
	xj "github.com/basgys/goxml2json"
	"bufio"
	"os"
	"fmt"
	"encoding/json"
)




type OCS_Software_Response struct {
	IPAddress           string    `json:"IPAddress,omitempty"`
	Name                string    `json:"HostName,omitempty"`
	SoftwareVersion     string    `json:"SoftwareVersion,omitempty"`
	SoftwareName        string    `json:"SoftwareName,omitempty"`
	SoftwareInstallDate time.Time `json:"SoftwareInstallDate,omitempty"`
}

type OCS_Hardware_Response struct {
	IPAddress  string `json:"IPAddress,omitempty"`
	MACAddress string `json:"MACAddress,omitempty"`
	Cores      string `json:"CPUCores,omitempty"`
	Name       string `json:"HostName,omitempty"`
	Memory     int32  `json:"RAM,omitempty"`
	SwapMemory int32  `json:"SwapMemory,omitempty"`
	OSName     string `json:"OperatingSystem,omitempty"`
}

type OCS_Drives_Response struct {
	IPAddress     string `json:"IPAddress,omitempty"`
	Name          string `json:"HostName,omitempty"`
	FileSystem    string `json:"FileSystem,omitempty"`
	VolumeName    string `json:"VolumeName,omitempty"`
	MountLocation string `json:"MountLocation,omitempty"`
}

type OCS_Data_Request struct{
	Request	OCS_Request	`json:"REQUEST,omitempty"`
}

type OCS_Request struct {
	DeviceID string      `json:"DEVICEID,omitempty"`
	Query    string      `json:"QUERY,omitempty"`
	Content  OCS_Content `json:"CONTENT,omitempty"`
}

type OCS_Content struct {
	CPUData     OCS_CPUData      `json:"CPUS,omitempty"`
	Hardware    OCS_Hardware     `json:"HARDWARE,omitempty"`
	Inputs      []OCS_Input      `json:"INPUTS,omitempty"`
	Networks    string    `json:"NETWORKS,omitempty"`
	//BIOS        OCS_BIOS         `json:"BIOS,omitempty"`
	Drives      []OCS_Drive      `json:"DRIVES,omitempty"`
	Softwares   []OCS_Software   `json:"SOFTWARES,omitempty"`
	//Sounds      []OCS_Sound      `json:"SOUNDS,omitempty"`
	//Videos      OCS_Video        `json:"VIDEOS,omitempty"`
	//Controllers []OCS_Controller `json:"CONTROLLERS,omitempty"`
}

type OCS_CPUData struct {
	//DataWidth    string `json:"DATA_WIDTH,omitempty"`
	//L2CacheSize  string `json:"L2CACHESIZE,omitempty"`
	//LogicalCPUs  string `json:"LOGICAL_CPUS,omitempty"`
	Manufacturer string `json:"MANUFACTURER,omitempty"`
	//CPUType      string `json:"TYPE,omitempty"`
	Cores        string `json:"CORES,omitempty"`
	//CPUArch      string `json:"CPUARCH,omitempty"`
	//CurrentSpeed string `json:"CURRENT_SPEED,omitempty"`
}

type OCS_Hardware struct {
	//Checksum           string `json:"CHECKSUM,omitempty"`
	//Description        string `json:"DESCRIPTION,omitempty"`
	//LastLoggedUser     string `json:"LASTLOGGEDUSER,omitempty"`
	Memory             string `json:"MEMORY,omitempty"`
	ProcessorN         string `json:"PROCESSORN,omitempty"`
	Processors         string `json:"PROCESSORS,omitempty"`
	Name               string `json:"NAME,omitempty"`
	OSName             string `json:"OSNAME,omitempty"`
	//ProcessorT         string `json:"PROCESSORT,omitempty"`
	//Swap               string `json:"SWAP,omitempty"`
	//DateLastLoggedUser string `json:"DATELASTLOGGEDUSER,omitempty"`
	IPAddress          string `json:"IPADDR,omitempty"`
	OSComments         string `json:"OSCOMMENTS,omitempty"`
	OSVersion          string `json:"OSVERSION,omitempty"`
}

type OCS_Input struct {
	Description string `json:"DESCRIPTION,omitempty"`
	Type        string `json:"TYPE,omitempty"`
}

type OCS_Network struct {
	//Description string `json:"DESCRIPTION,omitempty"`
	Type        string `json:"TYPE,omitempty"`
	//TypeMIB     string `json:"TYPEMIB,omitempty"`
	Speed       string `json:"SPEED,omitempty"`
	MACAddr     string `json:"MACADDR,omitempty"`
	//Status      string `json:"STATUS,omitempty"`
	IPAddress   string `json:"IPADDRESS,omitempty"`
	//IPMask      string `json:"IPMASK,omitempty"`
	//IPGateway   string `json:"IPGATEWAY,omitempty"`
	//IPSubnet    string `json:"IPSUBNET,omitempty"`
	//IPDHCP      string `json:"IPDHCP,omitempty"`
	//MTU         string `json:"MTU,omitempty"`
}
/*
type OCS_BIOS struct {
	//SManufacturer string `json:"SMANUFACTURER,omitempty"`
	//SModel        string `json:"SMODEL,omitempty"`
	SSN           string `json:"SSN,omitempty"`
	//Type          string `json:"TYPE,omitempty"`
	//BManufacturer string `json:"BMANUFACTURER,omitempty"`
	//BVersion      string `json:"BVERSION,omitempty"`
	//BDate         string `json:"BDATE,omitempty"`
	//MManufacturer string `json:"MMANUFACTURER,omitempty"`
	//MModel        string `json:"MMODEL,omitempty"`
	//MSN           string `json:"MSN,omitempty"`
	AssetTag      string `json:"ASSETTAG,omitempty"`
}
*/
type OCS_Drive struct {
	//Type       string `json:"TYPE,omitempty"`
	VolumN     string `json:"VOLUMN,omitempty"`
	Filesystem string `json:"FILESYSTEM,omitempty"`
	Free       string `json:"FREE,omitempty"`
	Total      string `json:"TOTAL,omitempty"`
}
/*
type OCS_Sound struct {
	Description  string `json:"DESCRIPTION,omitempty"`
	Manufacturer string `json:"MANUFACTURER,omitempty"`
	Name         string `json:"NAME,omitempty"`
}
*/
type OCS_Software struct {
	Version     string `json:"VERSION,omitempty"`
	//Comments    string `json:"COMMENTS,omitempty"`
	//Filesize    string `json:"FILESIZE,omitempty"`
	//From        string `json:"FROM,omitempty"`
	Installdate string `json:"INSTALLDATE,omitempty"`
	Name        string `json:"NAME,omitempty"`
}
/*
type OCS_Controller struct {
	PCISlot      string `json:"PCISLOT,omitempty"`
	Manufacturer string `json:"MANUFACTURER,omitempty"`
	Name         string `json:"NAME,omitempty"`
	PCIID        string `json:"PCIID,omitempty"`
}

type OCS_Video struct {
	Name       string `json:"NAME,omitempty"`
	Chipset    string `json:"CHIPSET,omitempty"`
	Memory     string `json:"MEMORY,omitempty"`
	Resolution string `json:"RESOLUTION,omitempty"`
}
*/


func main(){
	file, err := os.Open("/tmp/ocsdata")
	handlerError(err)
	dat :=  bufio.NewReader(file)	
	jsonVal, err := xj.Convert(dat)
	handlerError(err)
	var OCSData OCS_Data_Request
	err = json.Unmarshal(jsonVal.Bytes(),&OCSData)
	handlerError(err)









	client, err := elastic.NewClient()
        if err != nil {
                fmt.Println(err.Error())
        }
        exists, err := client.IndexExists(indexName).Do(context.Background())
        if err != nil {
                panic(err)
        }
        if !exists {
                createIndex, err := client.CreateIndex(indexName).Body(indexMapping).Do(context.Background())
                if err != nil {
                        panic(err)
                }
                if !createIndex.Acknowledged {
                         // Not acknowledged
                } else {
                        fmt.Println("Created Index")
                }
        }

        for idx,valx := range DataArr {
                fmt.Println(reflect.TypeOf(valx))
                _, err := client.Index().Index(indexName).Type(docType).Id(strconv.Itoa(idx)).BodyJson(valx).Do(context.Background())
                if err != nil {
                        fmt.Println(idx)
                        fmt.Println(err.Error())
                }

        }





	
}

func handlerError(err error){
	if err != nil{
		fmt.Println(err.Error())
	}
}
