$(function() {
    let userInfo = getUserInfo()
    let loginDiv = $("#login")
    let loginedDiv = $("#logined")
    if (userInfo === undefined || userInfo === null) {
        loginDiv.show()
        loginedDiv.hide()
    } else {
        loginDiv.hide()
        loginedDiv.show()
        let account = ""
        if (userInfo["phone"] === "") {
            account = userInfo["email"]
        } else {
            account = userInfo["phone"]
        }
        SetHeaderUserInfo(userInfo["nickname"], account)
    }
})

function logout() {
    cleanUserInfo()
    window.location.href = '/'
}

function SetHeaderUserInfo(nickname, account) {
    let nicknameElem = $("#header_nickname_text")
    let accountElem = $("#header_account_text")
    nicknameElem.text(nickname)
    accountElem.text("@" + account)
}
