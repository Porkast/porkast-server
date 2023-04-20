$(function() {
    $("#bottom-audio-player").on("canplay", function() {
        let duration = $(this)[0].duration
        let totalTime = secondsToHourMunitesSeconds(Math.floor(duration))
        $("#small-bottom-audio-range-duration").text(totalTime)
        $("#bottom-audio-range-duration").text(totalTime)
    })

    $("#bottom-audio-player").on("timeupdate", function() {
        let currentTime = $(this)[0].currentTime
        let duration = $(this)[0].duration
        let formatTime = secondsToHourMunitesSeconds(Math.floor(currentTime))
        $("#small-bottom-audio-range-time").text(formatTime)
        $("#bottom-audio-range-time").text(formatTime)
        $("#bottom-audio-player-layout-small-time").text(formatTime)

        let rangeInput = $("#bottom-audio-range-input")
        let smallRangeInput = $("#small-bottom-audio-range-input")
        let currentRangeVal = caculateRangeInputVal(Math.floor(currentTime), Math.floor(duration))
        rangeInput.val(currentRangeVal)
        smallRangeInput.val(currentRangeVal)
    })

    $("#bottom-audio-range-input").change(function() {
        let onChangedValue = $(this).val()
        let bottomAudioPlayer = $("#bottom-audio-player")
        let duration = bottomAudioPlayer[0].duration
        let targetTime = Math.round((onChangedValue / 10000) * duration)
        setAudioCurrentTime(targetTime)
    })

    $("#small-bottom-audio-range-input").change(function() {
        let onChangedValue = $(this).val()
        let bottomAudioPlayer = $("#bottom-audio-player")
        let duration = bottomAudioPlayer[0].duration
        let targetTime = Math.round((onChangedValue / 10000) * duration)
        setAudioCurrentTime(targetTime)
    })
})

function playOrPause(event, feedItemId, audioSource, audioType, itemTitle, channelTitle, channelImageUrl) {
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
    setBottomAudioAudioInfo(itemTitle, channelTitle, channelImageUrl)
    event.stopPropagation();
}


function forwardAudioCurrentTime(forwardTime) {
    let bottomAudioPlayer = $("#bottom-audio-player")
    let currentTime = bottomAudioPlayer[0].currentTime
    let targetTime = currentTime + forwardTime
    setAudioCurrentTime(targetTime)
}

function backwardAudioCurrentTime(backwardTime) {
    let bottomAudioPlayer = $("#bottom-audio-player")
    let currentTime = bottomAudioPlayer[0].currentTime
    let targetTime = currentTime - backwardTime
    setAudioCurrentTime(targetTime)
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

function setAudioCurrentTime(currentTime) {
    let playerElement = $("#bottom-audio-player")
    playerElement[0].currentTime = currentTime
}

function setBottomAudioAudioInfo(itemTitle, channelTitle, channelImageUrl) {
    let channelTitleElem = $("#bottom-audio-channel-title")
    let itemTitleElem = $("#bottom-audio-item-title")
    let itemImgElem = $("#bottom-audio-channel-img")
    let bottomSmallLayoutItemTitleElem = $("#bottom-audio-player-layout-small-item-title")
    channelTitleElem.text(channelTitle)
    itemTitleElem.text(itemTitle)
    bottomSmallLayoutItemTitleElem.text(itemTitle)
    itemImgElem.attr("src", channelImageUrl)
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

function secondsToHourMunitesSeconds(totalSeconds) {
    let hours = Math.floor(totalSeconds / 3600);
    totalSeconds %= 3600;
    let minutes = Math.floor(totalSeconds / 60);
    let seconds = totalSeconds % 60;

    // If you want strings with leading zeroes:
    minutes = String(minutes).padStart(2, "0");
    hours = String(hours).padStart(2, "0");
    seconds = String(seconds).padStart(2, "0");
    let formatTime = hours + ":" + minutes + ":" + seconds
    return formatTime
}

function caculateRangeInputVal(currentTime, totalTime) {
    return Math.round((currentTime / totalTime) * 10000)
}
