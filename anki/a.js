

bWin = navigator.platform.indexOf("Win") == 0;

function playVideo() {
    var videoSrc = document.querySelector("#tv-src").innerHTML;
    py.link("ankiplay" + videoSrc);
}
function playAudio(bPlay) {
    var player = document.getElementById('audiotag');
    var audioSrc = player.getAttribute("src");
    if (bWin) {
        py.link("ankiplay" + audioSrc);
    }
    else {
        if (bPlay) player.play();
        else player.pause();
    }
}
function loopAudio(cb) {
    var bLoop = cb.checked;
    var player = document.getElementById('audiotag');
    if (bLoop)
        player.setAttribute("loop","true");
    else
        player.removeAttribute("loop");
}
function hidex(x) {
    if (x.length > 1) {
        for(var i=0; i<x.length; i++) {
            x[i].setAttribute("style","display:none")
        }
    }else {
        x.setAttribute("style","display:none")
    }
}
function update() {
    var videoSrc = document.querySelector("#tv-src").textContent;
    var tv = document.querySelector("#tv");
    if (bWin) {
        var mobile_only = document.getElementsByClassName("mobile");
        hidex(mobile_only);
    } else {
        tv.innerHTML = '<video width="100%" height="auto" src=' + videoSrc + ' controls="controls" poster="_friends01.jpg"></video>';
        var stage_face = document.querySelector("#stage-face");
        hidex(stage_face);
    }
    var notes = document.querySelector("#notes");
    if (!(notes.textContent)) {
        hidex(notes.parentElement);
    }
}
if (bWin) {
    v=document.querySelector("#debug").innerHTML;
    document.querySelector("#debug").innerHTML=v+"bwin";
    update()
} else {
    window.onload = function() {
        v=document.querySelector("#debug").innerHTML;
        document.querySelector("#debug").innerHTML=v+" bwin-else ";
        update()
    }
}

