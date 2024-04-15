# Quebot

## Description
사용자의 학습을 도와주는 `Quebot`입니다.\
각 사용자는 자신이 정한 주제에 맞는 랜덤 질문을 받고 답변에대한 정답 유사율을 `%`로 확인할 수 있습니다.

## Use case
1. 사용자는 자신이 학습하고 싶은 주제를 입력합니다.
2. 사용자는 랜덤 질문을 받습니다.
3. 사용자는 질문에 대한 답변을 입력합니다.
4. 사용자는 답변의 정답유사율을 확인합니다.
5. 유사한 질문받기 or 다른 질문받기를 선택합니다.
6. 2번으로 돌아갑니다.

## Structure
1. Core server
core server는 질문과 답변을 처리합니다.

2. Data server
data server는 질문과 답변을 저장하고 불러옵니다.

3. Trained server
trained server는 질문과 답변의 정답유사율을 계산합니다.

4. User server
user server는 사용자의 정보를 저장하고 불러옵니다.

## Dataset
질문과 답변의 데이터셋은 `KorQuAD`과 `KorNLI` 를 사용합니다.
학습 데이터는 `Wikipedia`의 데이터를 사용합니다.

