function playAudio(isPlay, source, type) {
    let playerSourceElement = $("#bottom-audio-player-source")
    let playerElement = $("#bottom-audio-player")
    let currentSource = playerSourceElement.attr("src")
    if (currentSource !== source) {
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
