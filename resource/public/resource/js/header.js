$(function() {
    let userInfo = getUserInfo()
    if (userInfo === undefined || userInfo === null) {
        let headerAvatar = $("#header-avatar")
        headerAvatar.hide()
    }
})
