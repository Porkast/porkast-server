$(function() {
    let loginBtn = $("#login_btn");
    let emailPhoneInput = $("#email_phone_input");
    let passwordInput = $("#password_input");
    loginBtn.click(function() {
        let emailPhoneInputText = emailPhoneInput.val()
        let passwordInputText = passwordInput.val()
        console.log("emailPhoneInputText : ", emailPhoneInputText)
        console.log("passwordInputText : ", passwordInputText)
    });
});

