$(function() {
    // let searchClearBtn = $("#searchClearBtn");
    let searchBtn = $("#homeSearchBtn");
    let searchInput = $("#searchInput");
    let searchBtnBottom = $("#homeSearchBtnBottom");
    let searchInputBottom = $("#searchInputBottom");
    let searchOptionElem = $("#search-options")
    searchInput.keypress(function(event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            let optionValue = searchOptionElem.val()
            if (optionValue === "sortByDate") {
                window.location.href = "/search?q=" + keyword + "&page=1&sortByDate=1";
            } else {
                window.location.href = "/search?q=" + keyword + "&page=1";
            }
        }
        event.stopPropagation();
    });
    searchBtn.click(function() {
        let keyword = searchInput.val()
        let optionValue = searchOptionElem.val()
        if (optionValue === "sortByDate") {
            window.location.href = "/search?q=" + keyword + "&page=1&sortByDate=1";
        } else {
            window.location.href = "/search?q=" + keyword + "&page=1";
        }
    });
    searchInputBottom.keypress(function(event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            let optionValue = searchOptionElem.val()
            if (optionValue === "sortByDate") {
                window.location.href = "/search?q=" + keyword + "&page=1&sortByDate=1";
            } else {
                window.location.href = "/search?q=" + keyword + "&page=1";
            }
        }
        event.stopPropagation();
    });
    searchBtnBottom.click(function() {
        let keyword = searchInputBottom.val()
        let optionValue = searchOptionElem.val()
        if (optionValue === "sortByDate") {
            window.location.href = "/search?q=" + keyword + "&page=1&sortByDate=1";
        } else {
            window.location.href = "/search?q=" + keyword + "&page=1";
        }
    });
});

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
