# SkyEye

選用COLLY

cmd 目錄：

用途：包含應用程序的主入口。在這裡，你會放置一個或多個主要的 .go 文件，用於啟動和運行你的爬蟲程序。
collector 目錄：

用途：存放與 Colly 收集器相關的設定和初始化代碼。這裡可以設置 Colly 收集器的選項、中間件、錯誤處理等。
handlers 目錄：

用途：處理 Colly 收集器的回調函數。每當 Colly 收集器抓取到數據時，會調用這些函數來處理和分析數據。
models 目錄：

用途：定義數據結構和模型。這些模型代表你將從網頁中抓取的數據結構。
storage 目錄：

用途：負責數據的儲存邏輯，包括數據的持久化處理。例如，將爬取的數據儲存到文件、數據庫或其他儲存系統中。
utils 目錄：

用途：存放共用的工具函式，如輔助函式、格式轉換工具等。
configs 目錄：

用途：包含配置文件和相關設定。在這裡，你可以放置爬蟲的配置選項，比如抓取頻率、目標網站等。
docs 目錄：

用途：存放專案文檔，如設計文檔、使用說明等。
tests 目錄：

用途：包含測試代碼，用於對爬蟲的各個部分進行單元測試和集成測試。
這個結構清晰地劃分了爬蟲專案的不同部分，幫助你管理和維護大型爬蟲應用。每個目錄的職責明確，易於理解和擴展。




my-colly-crawler/
├── cmd/
│   └── main.go
├── collector/
│   └── collector.go
├── handlers/
│   └── handler.go
├── models/
│   └── model.go
├── storage/
│   └── storage.go
├── utils/
│   └── utils.go
├── configs/
│   └── config.go
└── go.mod




go get github.com/stretchr/testify