$(function () {
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
        SetHeaderUserInfo(userInfo["id"], userInfo["nickname"], account)
    }
})

function logout() {
    cleanUserInfo()
    window.location.href = '/'
}

function SetHeaderUserInfo(userId, nickname, account) {
    let nicknameElem = $("#header_nickname_text")
    let accountElem = $("#header_account_text")
    let accountInfoElem = $("#header_account_info_tag")
    let listenLaterTag = $("#header_listen_later_playlist_tag")
    let userSubList = $("#header_user_sub_list_tag")
    nicknameElem.text(nickname)
    accountElem.text("@" + account)
    listenLaterTag.attr("href", "/listenlater/playlist/" + userId)
    accountInfoElem.attr("href", "/user/info/" + userId)
    userSubList.attr("href", "/user/sub/list/" + userId + "/1")
}
