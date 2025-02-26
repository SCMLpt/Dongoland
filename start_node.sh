#!/bin/bash
# Dongoland 메인넷 노드 시작 스크립트

# 데이터 디렉토리 설정
export ALGORAND_DATA=/path/to/dongoland/data

# 노드 시작
./goal node start -d $ALGORAND_DATA

echo "Dongoland Mainnet node started!"
