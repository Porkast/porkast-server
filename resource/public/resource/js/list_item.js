function FeedListItemPlayBtnOnClick(event, feedItemId, audioSource, audioType) {
    let playSvgElement = $("#list-item-play-svg-" + feedItemId)
    let pauseSvgElement = $("#list-item-pause-svg-" + feedItemId)
    let isPlay = playSvgElement.hasClass('hidden')
    if (isPlay) {
        playSvgElement.removeClass("hidden")
        pauseSvgElement.addClass("hidden")
    } else {
        pauseSvgElement.removeClass("hidden")
        playSvgElement.addClass("hidden")
    }
    playAudio(isPlay, audioSource, audioType)
    event.stopPropagation();
}
