function playOrPause(event, feedItemId, audioSource, audioType) {
    let isPlay = false
    let playSvgElement = $("#list-item-play-svg-" + feedItemId)
    if (playSvgElement.hasClass("hidden")) {
        isPlay = true
    }
    let bottomAudioTag = $("#bottom-audio-player")
    bottomAudioTag.attr("current-item-id", feedItemId)
    bottomAudioTag.attr("current-source", audioSource)
    bottomAudioTag.attr("current-type", audioType)
    doPlayOrPauseAudio(isPlay, feedItemId, audioSource, audioType)
    event.stopPropagation();
}

function bottomPlayOrPause() {
    let bottomAudioTag = $("#bottom-audio-player")
    let currentAudioSource = bottomAudioTag.attr("current-source")
    let currentAudioType = bottomAudioTag.attr("current-type")
    let currentFeedItemId = bottomAudioTag.attr("current-item-id")
    let btnPlayBottomPlayImg = $("#bottom-audio-player-play-btn-img")
    let isPlay = false
    if (btnPlayBottomPlayImg.hasClass("hidden")) {
        isPlay = true
    }
    doPlayOrPauseAudio(isPlay, currentFeedItemId, currentAudioSource, currentAudioType)
}

function resetAllPlayButton() {
    let allPlayeBtnElements = $('[id^="list-item-play-svg-"]')
    let allPauseBtnElements = $('[id^="list-item-pause-svg-"]')
    allPlayeBtnElements.removeClass("hidden")
    allPauseBtnElements.addClass("hidden")
}

function doPlayOrPauseAudio(isPlay, feedItemId, source, type) {
    let playerWrapperElement = $("#bottom-audio-player-layout")
    playerWrapperElement.removeClass("hidden")

    let playerSourceElement = $("#bottom-audio-player-source")
    let playerElement = $("#bottom-audio-player")
    let currentSource = playerSourceElement.attr("src")
    if (currentSource !== source) {
        resetAllPlayButton()
        playerSourceElement.attr("src", source)
        playerSourceElement.attr("type", type)
        playerElement[0].load()
    }

    let playSvgElement = $("#list-item-play-svg-" + feedItemId)
    let pauseSvgElement = $("#list-item-pause-svg-" + feedItemId)
    let btnPlayBottomPlayImg = $("#bottom-audio-player-play-btn-img")
    let btnPlayBottomPauseImg = $("#bottom-audio-player-pause-btn-img")
    let btnSmallPlayBottomPlayImg = $("#small-bottom-audio-player-play-btn-img")
    let btnSmallPlayBottomPauseImg = $("#small-bottom-audio-player-pause-btn-img")
    if (isPlay) {
        btnPlayBottomPlayImg.removeClass("hidden")
        btnPlayBottomPauseImg.addClass("hidden")
        btnSmallPlayBottomPlayImg.removeClass("hidden")
        btnSmallPlayBottomPauseImg.addClass("hidden")

        playSvgElement.removeClass("hidden")
        pauseSvgElement.addClass("hidden")
        playerElement[0].pause()
    } else {
        btnPlayBottomPlayImg.addClass("hidden")
        btnPlayBottomPauseImg.removeClass("hidden")
        btnSmallPlayBottomPlayImg.addClass("hidden")
        btnSmallPlayBottomPauseImg.removeClass("hidden")

        pauseSvgElement.removeClass("hidden")
        playSvgElement.addClass("hidden")
        playerElement[0].play()
    }
}

function hideBottomAudio() {
    let bottomAudioPlayerElem = $('#bottom-audio-player-layout')
    let bottomAudioSmallPlayerElem = $('#bottom-audio-player-layout-small')
    bottomAudioPlayerElem.addClass("hidden")
    bottomAudioPlayerElem.addClass("md:hidden")
    bottomAudioSmallPlayerElem.removeClass("hidden")
}

function showBottomAudio() {
    let bottomAudioPlayerElem = $('#bottom-audio-player-layout')
    let bottomAudioSmallPlayerElem = $('#bottom-audio-player-layout-small')
    bottomAudioPlayerElem.removeClass("hidden")
    bottomAudioPlayerElem.removeClass("md:hidden")
    bottomAudioSmallPlayerElem.addClass("hidden")
}
