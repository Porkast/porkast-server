$(function () {
    let loginBtn = $("#login_btn");
    let registerBtn = $("#register_btn")
    let nicknameInput = $("#nickname_input");
    let emailPhoneInput = $("#email_phone_input");
    let passwordInput = $("#password_input");
    let passwordInputVerify = $("#password_verify_input");
    loginBtn.click(function () {
        let emailPhoneInputText = emailPhoneInput.val()
        let passwordInputText = passwordInput.val()
        doLogin(emailPhoneInputText, passwordInputText)
    });
    registerBtn.click(function () {
        let nicknameInputText = nicknameInput.val()
        let emailPhoneInputText = emailPhoneInput.val()
        let passwordInputText = passwordInput.val()
        let passwordVerifyInputText = passwordInputVerify.val()
        doRegister(nicknameInputText, emailPhoneInputText, passwordInputText, passwordVerifyInputText)
    });
});

function doLogin(account, password) {
    let email = ""
    let phone = 0
    if (isEmail(account)) {
        email = account
    } else if (isMobile(account)) {
        phone = account
    } else {
        // TODO: alert account format is not valid
        return
    }
    let postData = {
        password: password,
        email: email,
        phone: parseInt(phone),
    }
    $.ajax({
        method: 'POST',
        url: '/v1/api/user/login',
        data: JSON.stringify(postData),
        success: function (data) {
            let jsonData = data
            console.log(jsonData)
            if (jsonData.code !== 0) {
                console.log("do login failed")
                ShowErrorAlert(jsonData.message)
            } else {
                ShowSuccessAlert(jsonData.message)
                setUserInfo(jsonData.data)
                window.location.href = '/'
            }
        },
        error: function (data) {
            ShowErrorAlert(data.message)
        }
    })
}

function doRegister(nickname, account, password, vPassword) {
    let email = ""
    let phone = 0
    if (isEmail(account)) {
        email = account
    } else if (isMobile(account)) {
        phone = account
    } else {
        // TODO: alert invalid account
        return
    }
    let postData = {
        Nickname: nickname,
        password: password,
        passwordVerify: vPassword,
        email: email,
        phone: parseInt(phone),
    }
    $.ajax({
        method: 'POST',
        url: '/v1/api/user/register',
        data: JSON.stringify(postData),
        success: function (data) {
            let jsonData = data
            if (jsonData.code !== 0) {
                console.log("do register failed")
                console.log(data)
                ShowErrorAlert(jsonData.message)
            } else {
                ShowSuccessAlert(jsonData.message)
                setUserInfo(jsonData.data)
                window.location.href = '/'
            }
        },
        error: function (data) {
        }
    })
}
