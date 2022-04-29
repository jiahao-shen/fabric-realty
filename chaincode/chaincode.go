package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
// func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
// 	fmt.Println("链码初始化")
// 	//初始化默认数据
// 	var accountIds = [6]string{
// 		"5feceb66ffc8",
// 		"6b86b273ff34",
// 		"d4735e3a265e",
// 		"4e07408562be",
// 		"4b227777d4dd",
// 		"ef2d127de37b",
// 	}
// 	var userNames = [6]string{"管理员", "①号业主", "②号业主", "③号业主", "④号业主", "⑤号业主"}
// 	var balances = [6]float64{0, 5000000, 5000000, 5000000, 5000000, 5000000}
// 	//初始化账号数据
// 	for i, val := range accountIds {
// 		account := &model.Account{
// 			AccountId: val,
// 			UserName:  userNames[i],
// 			Balance:   balances[i],
// 		}
// 		// 写入账本
// 		if err := utils.WriteLedger(account, stub, model.AccountKey, []string{val}); err != nil {
// 			return shim.Error(fmt.Sprintf("%s", err))
// 		}
// 	}
// 	return shim.Success(nil)
// }

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")

	dataNames := []string{"姓名", "性别", "青霉素过敏", "阿司匹林过敏", "姓名", "性别", "青霉素过敏", "阿司匹林过敏"}
	dataIDs := []string{"HospitalA-Name", "HospitalA-Sex", "HospitalA-Internal-Penicillin", "HospitalA-Surgery-Aspirin", "HospitalB-Name", "HospitalB-Sex", "HospitalB-Internal-Penicillin", "HospitalB-Surgery-Aspirin"}
	authors := []string{"HospitalA", "HospitalA", "HospitalA-Internal", "HospitalA-Surgery", "HospitalB", "HospitalB", "HospitalB-Internal", "HospitalB-Surgery"}

	for i, val := range dataNames {
		dataitem := &model.DataItem{
			Name:         val,
			ID:           dataIDs[i],
			Introduction: "",
			Author:       authors[i],
			Type:         "",
			Shared:       "shared",
			Resource:     "",
			Classified:   "",
			Version:      "1.0.0",
			Location:     "localhost",
			Created:      time.Now(),
		}

		if err := utils.WriteLedger(dataitem, stub, model.DataItemKey, []string{dataIDs[i]}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	orgNames := []string{"医院A", "内科", "外科", "医院B", "内科", "外科"}
	orgIDs := []string{"HospitalA", "HospitalA-Internal", "HospitalA-Surgery", "HospitalB", "HospitalB-Internal", "HospitalB-Surgery"}
	superiors := []string{"", "HospitalA", "HospitalA", "", "HospitalB", "HospitalB"}
	subordinates := [][]string{{"HospitalA-Internal", "HospitalA-Surgery"}, nil, nil, {"HospitalB-Internal", "HospitalB-Surgery"}, nil, nil}
	orgDataItems := [][]string{
		{"HospitalA-Name", "HospitalA-Sex"},
		{"HospitalA-Internal-Penicillin"},
		{"HospitalA-Surgery-Aspirin"},
		{"HospitalB-Name", "HospitalB-Sex"},
		{"HospitalB-Internal-Penicillin"},
		{"HospitalB-Surgery-Aspirin"},
	}

	for i, val := range orgNames {
		org := &model.Organization{
			Name:         val,
			ID:           orgIDs[i],
			Type:         "Medical",
			Superior:     superiors[i],
			Subordinates: subordinates[i],
			DataItems:    orgDataItems[i],
			Created:      time.Now(),
		}

		if err := utils.WriteLedger(org, stub, model.OrganizationKey, []string{orgIDs[i]}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	case "queryAccountList":
		return api.QueryAccountList(stub, args)
	case "createRealEstate":
		return api.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return api.QueryRealEstateList(stub, args)
	case "createSelling":
		return api.CreateSelling(stub, args)
	case "createSellingByBuy":
		return api.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return api.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return api.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return api.UpdateSelling(stub, args)
	case "createDonating":
		return api.CreateDonating(stub, args)
	case "queryDonatingList":
		return api.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return api.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return api.UpdateDonating(stub, args)
	case "queryOrganizationList":
		return api.QueryOrganizationList(stub, args)
	case "queryDataItemList":
		return api.QueryDataItemList(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
