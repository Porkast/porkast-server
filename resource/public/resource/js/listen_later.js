function AddToListenLater(channelId, itemId) {

    let userInfo = getUserInfo()
    let userId = userInfo['id']

    if (userId === "") {
        console.log("cannot get user id")
        return
    }

    let postData = {
        UserId: userId,
        ChannelId: channelId,
        ItemId: itemId,
    }
    $.ajax({
        method: 'POST',
        headers: {
            Authorization: getAuthToken()
        },
        url: '/v1/api/listenlater/item',
        data: JSON.stringify(postData),
        success: function (data) {
            let jsonData = data
            if (jsonData.code !== 0) {
                if (jsonData.message === "DB_DATA_ALREADY_EXIST") {
                }
            } else {
            }
        },
        error: function (data) {
        }
    })
}