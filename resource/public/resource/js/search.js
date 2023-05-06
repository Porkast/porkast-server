$(function() {
    // let searchClearBtn = $("#searchClearBtn");
    let searchBtn = $("#homeSearchBtn");
    let searchInput = $("#searchInput");
    let searchBtnBottom = $("#homeSearchBtnBottom");
    let searchInputBottom = $("#searchInputBottom");
    searchInput.keypress(function(event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            window.location.href = "/search?q=" + keyword + "&page=1";
        }
        event.stopPropagation();
    });
    searchBtn.click(function() {
        let keyword = searchInput.val()
        window.location.href = "/search?q=" + keyword + "&page=1";
    });
    searchInputBottom.keypress(function(event) {
        var keycode = (event.keyCode ? event.keyCode : event.which);
        if (keycode == '13') {
            let keyword = $(this).val()
            window.location.href = "/search?q=" + keyword + "&page=1";
        }
        event.stopPropagation();
    });
    searchBtnBottom.click(function() {
        let keyword = searchInputBottom.val()
        window.location.href = "/search?q=" + keyword + "&page=1";
    });
});
