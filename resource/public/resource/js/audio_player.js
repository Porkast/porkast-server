function FeedListItemPlayBtnOnClick(event, feedItemId, audioSource, audioType) {
    let playSvgElement = $("#list-item-play-svg-" + feedItemId)
    let pauseSvgElement = $("#list-item-pause-svg-" + feedItemId)
    let isPlay = playSvgElement.hasClass('hidden')
    playAudio(isPlay, audioSource, audioType)
    if (isPlay) {
        playSvgElement.removeClass("hidden")
        pauseSvgElement.addClass("hidden")
    } else {
        pauseSvgElement.removeClass("hidden")
        playSvgElement.addClass("hidden")
    }
    event.stopPropagation();
}

function resetAllPlayButton() {
    let allPlayeBtnElements = $('[id^="list-item-play-svg-"]')
    let allPauseBtnElements = $('[id^="list-item-pause-svg-"]')
    allPlayeBtnElements.removeClass("hidden")
    allPauseBtnElements.addClass("hidden")
}

function playAudio(isPlay, source, type) {
    let playerSourceElement = $("#bottom-audio-player-source")
    let playerElement = $("#bottom-audio-player")
    let currentSource = playerSourceElement.attr("src")
    if (currentSource !== source) {
        resetAllPlayButton()
        playerSourceElement.attr("src", source)
        playerSourceElement.attr("type", type)
        playerElement[0].load()
    }
    if (isPlay) {
        playerElement[0].pause()
    } else {
        playerElement[0].play()
    }
}
