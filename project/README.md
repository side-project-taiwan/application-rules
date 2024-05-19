
## 資料儲存方案
1. 616_demo方案：資料直接寫死前端網頁（或是使用 discord channel 當作 data sourceS

2. 群友建議直接起 Postgres/MySQL DB 並設計好 schema ，以免繞遠路

## 專案列表審核 work flow
1. 前端網頁 post 資料到 discord channel ( 使用 webhookUrl )
2. 審核人員將資料從 channel 上複製下來，貼到 vscode，例如 test.json 內，將 %7B 改成 '{'，將 %7D 改成 '}'
3. format json 檔案，就可得到易讀的 json，並且審核並與審核方確認細節
4. 若審核未通過，按照「審核未通過原因格式」回覆專案 owner
5. 若審核通過，通知專案 owner 已通過，並將 dataFormat 的原本資料格式，新增至 DC 「已過審專案」中

## 前端程式碼相關(先列API但尚未實作，參考程式碼在 discordSdkInGolang.go 中)
GET /api/v1/project-list