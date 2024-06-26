# Match-Dating
Implement a Match Dating HTTP server for practice system design

<!-- vscode-markdown-toc -->
* 1. [Functional Requirement](#FunctionalRequirement)
* 2. [Non-functional Requirement](#Non-functionalRequirement)
* 3. [如何儲存資料及查詢](#)
* 4. [Project Layout](#ProjectLayout)
* 5. [API](#API)
* 6. [TODO](#TODO)

<!-- vscode-markdown-toc-config
	numbering=true
	autoSave=true
	/vscode-markdown-toc-config -->
<!-- /vscode-markdown-toc -->

##  1. <a name='FunctionalRequirement'></a>Functional Requirement

1. 新增使用者資料且回傳有幾種配對結果

2. 刪除使用者資料

3. 取得最多N種配對結果

以下是配對規則
- 使用者會有這些資訊，姓名、身高、性別、想要約會的次數
- 男生只能跟比他矮的女生配對，反之女生只能跟比他高的男生配對
- 當有配對結果時，男生及女生的想要約會次數就必須減ㄧ，當變零的時候就必須從中移除。

*要思考一下，重複配對的情況*

##  2. <a name='Non-functionalRequirement'></a>Non-functional Requirement

1. 資料不進資料庫或外部快取，只存在本地的快取。


##  3. <a name=''></a>如何儲存資料及查詢

根據需求，使用者資訊必須存放在 in-memory，所以需要設計一個資料結構來儲存使用者資訊。這個資料結構必須有以下特性

1. 依照條件設定去插入且越快越好。
2. 快速搜尋到某個值。
3. 快速刪除某個值。

根據以上特性參考了幾個資料結構及演算法，如下。

* n 為資料數量。
* 有兩個時間複雜度時，前者為平均，後者為最差。

| Data Struct | Algorithm | Search | Insert | Remove |
| - | - | - | - | - |
| Array | 線性搜尋 | O(n) | O(n) | O(n) |
| Tree | Binary Search | O(log n) / O(n) | O(log n) / O(n) | O(log n) / O(n) |
| Tree | AVL | O(log n) | O(log n) | O(log n) | 
| Tree | 紅黑樹 RBT | O(log n) | O(1) / O(log n) | O(1) / O(log n) | 


根據以上表格 AVL 與 RBT 可以盡可能做到 log n 的時間複雜度，但在這邊我會選擇 RBT，原因是AVL在樹的高度平衡做的嚴謹，會比RBT多做幾次旋轉，意味著會比較浪費，故選擇紅黑數來當作我的資料結構。


##  4. <a name='ProjectLayout'></a>Project Layout
```
├── build                -- 放置dockerfile 或者其他建構檔案
├── cmd                  -- 主要的程式 
├── configs              -- 配置檔案 
├── docs                 -- 文件放置
│   └── postman          -- API文件以Postman呈現      
├── internal             -- 私有程式碼     
│   ├── config           -- 配置檔結構與操作       
│   ├── engine           -- 引擎的實現       
│   ├── handler          -- 處理路由
│   ├── model            -- 各個模型      
│   └── server           -- http gin server      
├── pkg                  -- 可以給別人用的
│   └── logger           -- zerolog相關配置       
└── test                 -- 放unit test的
```
##  5. <a name='API'></a>API

請使用 Postman 導入此[文件](./docs/postman/Match-Dating.postman_collection.json)

##  6. <a name='TODO'></a>TODO 
- [ ] QuerySinglePeople API - 取得最多N種配對結果