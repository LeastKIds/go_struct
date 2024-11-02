- 비지니스 로직 처리 및 도메인 레이어와의 상호작용
1. domain의 repository로 부터 유저 정보 습득
2. domain의 service로부터 비지니스 로직 실행
3. domain의 repository로 부터 유저 정보 갱신
4. 결과를 adapter/http/handler로 반환
- interface 사용