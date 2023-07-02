$(function() {
    if ('serviceWorker' in navigator) {
        navigator.serviceWorker.register("/resource/js/service-worker.js");
    }
})

function setUserInfo(userInfo) {
    userInfo['auth'] = userInfo['token'] + "@@" + userInfo['id']
    localStorage.setItem("userInfo", JSON.stringify(userInfo))
}

function getUserInfo() {
    let userInfo = JSON.parse(localStorage.getItem("userInfo"))
    if (userInfo === undefined || userInfo === null) {
        return null
    }
    let uid = userInfo['id']
    if (uid !== null && uid !== undefined && uid !== '') {
        return userInfo
    }
    return null
}

function getAuthToken() {
    let userInfo = getUserInfo()
    let auth = ""
    if (userInfo !== null) {
        userId = userInfo.id
        token = userInfo.token
    }
    auth = token + "@@" + userId
    return auth
}

function cleanUserInfo() {
    localStorage.removeItem("userInfo")
}

function isEmail(email) {
    return String(email)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,:\s@"]+(\.[^<>()[\]\\.,:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        )
}

function isMobile(mobile) {
    let reg = /^1[3-9]\d{9}$/
    return reg.test(mobile)
}


function share(channelId, itemId) {
    let domain = window.location.host
    let shareUrl = 'http://' + domain + '/share/feed/' + channelId + '/item/' + itemId
    $('#clipboard-temp-holder-' + itemId).text(shareUrl)
    copyToClickBoard('clipboard-temp-holder-' + itemId)

}

function shareChannel(channelId) {
    let domain = window.location.host
    let shareUrl = 'http://' + domain + '/share/feed/channel/' + channelId
    $('#clipboard-temp-holder-' + channelId).text(shareUrl)
    copyToClickBoard('clipboard-temp-holder-' + channelId)

}

function copyToClickBoard(elemId) {
    var content = document.getElementById(elemId).innerHTML;
    navigator.clipboard.writeText(content)
        .then(() => {
            ShowSuccessAlert('已复制到剪贴板')
        })
        .catch(err => {
            console.log('Something went wrong', err);
        })
}

function ShowSuccessAlert(alertMsg) {
    let alertElem = $("#success-alert-elem")
    let alertMsgElem = $("#success-alert-msg")
    alertMsgElem.text(alertMsg)
    alertElem.removeClass("hidden")
    setTimeout(function() {
        alertElem.addClass("hidden")
    }, 5000);
}

function ShowErrorAlert(alertMsg) {
    let alertElem = $("#error-alert-elem")
    let alertMsgElem = $("#erro-alert-msg")
    alertMsgElem.text(alertMsg)
    alertElem.removeClass("hidden")
    setTimeout(function() {
        alertElem.addClass("hidden")
    }, 5000);
}

function ShowToLoginAlert() {
    let alertElem = $("#login-alert-elem")
    let alertMsgElem = $("#login-alert-msg")
    alertMsgElem.text("还没登录，先去登录或者注册吧")
    alertElem.removeClass("hidden")
}

function loginAlertCancel() {
    let alertElem = $("#login-alert-elem")
    alertElem.addClass("hidden")
}

function loginAlertConfirm() {
    let alertElem = $("#login-alert-elem")
    alertElem.addClass("hidden")
    window.location.href = '/login'
}
