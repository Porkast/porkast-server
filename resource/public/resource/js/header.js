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
    }
})

function logout() {
    cleanUserInfo()
    window.location.href = '/'
}
