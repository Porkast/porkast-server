$(function () {
    // let searchClearBtn = $("#searchClearBtn");
    let searchBtn = $("#homeSearchBtn");
    let searchInput = $("#searchInput");
    let searchBtnBottom = $("#homeSearchBtnBottom");
    let searchInputBottom = $("#searchInputBottom");
    let searchOptionElem = $("#searchOptions")
    let searchOptionBottomElem = $("#searchOptionsBottom")
    let isScopeChannel = searchBtn.attr("scope-channel")
    searchOptionElem.change(function () {
        let value = $(this).val()
        searchOptionBottomElem.val(value)
    })

    searchOptionBottomElem.change(function () {
        let value = $(this).val()
        searchOptionElem.val(value)
    })

    searchInput.keypress(function (event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            let optionValue = searchOptionElem.val()
            doSearch(optionValue, isScopeChannel, keyword)
        }
        event.stopPropagation();
    });
    searchBtn.click(function () {
        let keyword = searchInput.val()
        let optionValue = searchOptionElem.val()
        doSearch(optionValue, isScopeChannel, keyword)
    });
    searchInputBottom.keypress(function (event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            let optionValue = searchOptionElem.val()
            doSearch(optionValue, isScopeChannel, keyword)
        }
        event.stopPropagation();
    });
    searchBtnBottom.click(function () {
        let keyword = searchInputBottom.val()
        let optionValue = searchOptionElem.val()
        doSearch(optionValue, isScopeChannel, keyword)
    });

    SetIsKeywordMatchTitleOnly()
    let searchOnlyMatchTitleCheckbox = $('#searchOnlyMatchTitleCheckbox')
    searchOnlyMatchTitleCheckbox.change(function () {
        if ($(this).is(':checked')) {
            SetSearchOnlyMatchTitle(true)
        } else {
            SetSearchOnlyMatchTitle(false)
        }
    })

    let searchOnlyMatchTitleCheckboxBottom = $('#searchOnlyMatchTitleCheckboxBottom')
    searchOnlyMatchTitleCheckboxBottom.change(function () {
        if ($(this).is(':checked')) {
            SetSearchOnlyMatchTitleBottom(true)
        } else {
            SetSearchOnlyMatchTitleBottom(false)
        }
    })

});

function doSearch(optionValue, isScopeChannel, keyword) {
    let sortByDate
    let scope
    if (optionValue === "sortByDate") {
        sortByDate = 1
    } else {
        sortByDate = 0
    }

    if (isScopeChannel) {
        scope = "channel"
    } else {
        scope = "item"
    }

    window.location.href = "/search?q=" + keyword + "&page=1&sortByDate=" + sortByDate + "&scope=" + scope;
}

function SearchFeedChannel() {
    let searchInput = $("#searchInput");
    let keyword = searchInput.val()
    window.location.href = "/search?q=" + keyword + "&page=1&scope=channel";
}

function SearchFeedItem() {
    let searchInput = $("#searchInput");
    let keyword = searchInput.val()
    window.location.href = "/search?q=" + keyword + "&page=1";
}

function SetIsKeywordMatchTitleOnly() {
    let searchInput = $("#searchInput");
    let keyword = searchInput.val()
    let titleMatchOnlyCheckbox = $('#searchOnlyMatchTitleCheckbox')
    if (keyword !== undefined && keyword !== null && keyword !== "") {
        if (keyword.startsWith('"') && keyword.endsWith('"')) {
            titleMatchOnlyCheckbox.prop('checked', true);
        }
    }

    let searchInputBottom = $("#searchInputBottom");
    let keywordBottom = searchInputBottom.val()
    let titleMatchOnlyCheckboxBottom = $('#searchOnlyMatchTitleCheckboxBottom')
    if (keywordBottom !== undefined && keywordBottom !== null && keywordBottom !== "") {
        if (keyword.startsWith('"') && keyword.endsWith('"')) {
            titleMatchOnlyCheckboxBottom.prop('checked', true);
        }
    }
}

function SetSearchOnlyMatchTitle(isMatchOnly) {
    let searchInput = $("#searchInput");
    let keyword = searchInput.val()
    if (keyword !== undefined && keyword !== null && keyword !== "") {
        if (isMatchOnly) {
            if (!keyword.startsWith('"') && !keyword.endsWith('"')) {
                searchInput.val('"' + keyword + '"')
            }
        } else {
            if (keyword.startsWith('"') && keyword.endsWith('"')) {
                let unMatchKeyword = keyword.replaceAll('"', "");
                searchInput.val(unMatchKeyword)

            }
        }
    }
}

function SetSearchOnlyMatchTitleBottom(isMatchOnly) {
    let searchInput = $("#searchInputBottom");
    let keyword = searchInput.val()
    if (keyword !== undefined && keyword !== null && keyword !== "") {
        if (isMatchOnly) {
            if (!keyword.startsWith('"') && !keyword.endsWith('"')) {
                searchInput.val('"' + keyword + '"')
            }
        } else {
            if (keyword.startsWith('"') && keyword.endsWith('"')) {
                let unMatchKeyword = keyword.replaceAll('"', "");
                searchInput.val(unMatchKeyword)

            }
        }
    }
}

function SubscriptSearchKeyword() {
    let searchInput = $("#searchInputBottom");
    let keyword = searchInput.val()
    let searchOptionElem = $("#searchOptions")
    let optionValue = searchOptionElem.val()
    let sortByDate
    if (optionValue === "sortByDate") {
        sortByDate = 1
    } else {
        sortByDate = 0
    }

    let userInfo = getUserInfo()
    if (userInfo === undefined || userInfo === null) {
        ShowToLoginAlert("")
        return
    }
    let userId = userInfo['id']

    if (userId === "") {
        console.log("cannot get user id")
        return
    }

    let postData = {
        userId: userId,
        keyword: keyword,
        sortByDate: sortByDate,
        lang: "zh-cn"
    }

    $.ajax({
        method: 'POST',
        url: '/v1/api/subscription/keyword',
        headers: {
            Authorization: getAuthToken()
        },
        data: JSON.stringify(postData),
        success: function (data) {
            let jsonData = data
            if (jsonData.code !== 0) {
                ShowErrorAlert(jsonData.message)
            } else {
                ShowSuccessAlert(jsonData.message)
            }
        },
        error: function (data) {
        }
    })
}

