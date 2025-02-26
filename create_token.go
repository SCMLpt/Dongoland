package main

import (
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/types"
)

func main() {
	// Algod 클라이언트 설정 (메인넷 노드 URL과 API 토큰 필요)
	algodClient, err := algod.MakeClient("https://mainnet-api.dongoland.io", "your-api-token")
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	// 계정 생성 (테스트용 개인키와 공개키)
	privateKey, publicKey := crypto.GenerateAccount()
	fmt.Printf("Account Address: %s\n", publicKey.Address())

	// ASA (토큰) 생성 매개변수
	params, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("Failed to get suggested params: %v\n", err)
		return
	}

	// 토큰 설정 (DONG)
	assetCreateTxn, err := types.ConstructAssetCreateTxn(
		publicKey.Address(),      // 발행자 주소
		1000000000,              // 총 발행량 (1,000,000,000 DONG)
		6,                       // 소수점 자리 수
		"DONG",                  // 단위 이름
		"Dongoland Token",       // 토큰 이름
		publicKey.Address(),     // 매니저 주소
		publicKey.Address(),     // 예약 주소
		publicKey.Address(),     // 동결 주소
		params,
	)
	if err != nil {
		fmt.Printf("Failed to create asset transaction: %v\n", err)
		return
	}

	// 트랜잭션 서명 및 보내기
	signedTxn, err := crypto.SignTransaction(privateKey, assetCreateTxn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %v\n", err)
		return
	}

	txID, err := algodClient.SendRawTransaction(signedTxn).Do(context.Background())
	if err != nil {
		fmt.Printf("Failed to send transaction: %v\n", err)
		return
	}

	fmt.Printf("Asset created with TXID: %s\n", txID)
}
