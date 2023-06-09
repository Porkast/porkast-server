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
