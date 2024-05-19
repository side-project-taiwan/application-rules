const log = console.log
import fetch from 'node-fetch';


const url = 'put your webhook url here';


const projectContent = {
    "title": "Side Project Taiwan 官方網站",
    "description": "Side Project Taiwan 是一個由軟體開發愛好者成立的開源友善社群, 社群成員包含「開發」「設計」「專案管理」等等想參與專案開發的夥伴！",
    "imageUrl": "",
    "tags": ["PM", "UIUX", "前端", "後端"],
    "isSoftDelete": false,
    "github_url": "https://github.com/side-project-taiwan/sideproj.tw",
    "create_at": "",
    "project_started_at": "2024/02/10",
    "isActive": true,
    "owner": { "name": "SPT Team", "personal_github": "https://github.com/side-project-taiwan", "role": "owner" }
}


// Example data to send in the POST request
const dcApiDataFormat = {
    "content": "申請專案",
    "username": "專案申請小助手",
    "avatar_url": "https://path.to.your/avatar.png",
    "embeds": [{
        "title": "列表專案申請內容（請手動修改%7B 為 '{'，%7D 為 '}'，貼到 vscode 內 format 為 json格式",
        "description": JSON.stringify(projectContent),
        "color": 65280,
        // "image": {
        //     "url": "https://fountain.org.tw/upload/upload/repository/74a7f73b7f18d193ddebff71c0b8afeaimage_normal.jpg"
        // }
    }]
};


fetch(url, {
        method: 'POST',
        body: JSON.stringify(dcApiDataFormat),
        headers: { 'Content-Type': 'application/json' },
    })
    // .then((response) => response.json())
    .then((result) => {
        console.log('Success:', result);
    })
    .catch((error) => {
        console.error('Error:', error);
    });