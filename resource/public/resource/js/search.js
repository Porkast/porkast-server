$(function() {
    // let searchClearBtn = $("#searchClearBtn");
    let searchBtn = $("#homeSearchBtn");
    let searchInput = $("#searchInput");
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
});
