console.log('------------- ++++++++++++++ content.js loaded ------------- ++++++++++++++ ');
//chrome app
//document.body.innerHTML = document.body.innerHTML.replace(new RegExp("uno", "g"), "dos");
function stopEvent(event){
    console.log('stopEvent:', event.type, event);
    event.preventDefault();
    event.stopImmediatePropagation();
    event.stopPropagation();
    event.cancelBubble = true;
    event.bubbles = false;
    event.sp = 'stop';
    return false;
}
function cancelEvent(e,preventDefault){
    if(!e){
        return
    }
    console.log('cancelEvent:', e.type, e);
    if(preventDefault){
        if(e.preventDefault){e.preventDefault()}else{e.returnValue=false}}
    if(e.stopPropagation){e.stopPropagation()}else{e.cancelBubble=true}
}
function init() {
    if($('#tuWin').length > 0)
        return;

    $('body').append('<div id="tuWin" class="tu-window">' +
        'Selected element: <span id="selectedElement"></span><div class="btns">' +
        '<a id="tuCancel" href="#">Cancel</a>' +
        '<a id="tuOk" href="#">Save</a></div>' +
        '</div>');

    var $app = $('*:not(#tuWin *)');
    var selectedElementSelector = null;

    $app.hover(function () {
        $('.tu-active').removeClass('tu-active');
        $(this).addClass('tu-active');
    }, function () {
        $(this).removeClass('tu-active');
    });

    //var elms = document.getElementsByTagName('a');
    $app.on('click', function(e) {
        cancelEvent(e, true);
        return false;
    });
    $app.on('mouseup', function(e) {
        cancelEvent(e, true);
        return false;
    });
    $app.on('mousedown', function(e) {
        cancelEvent(e, true);
        return false;
    });




    //$app.on('click', function ( event ) {
    //$app.click(function ( event ) {

    //});

    $('#tuCancel').click(function ( event ) {
        event.preventDefault();
        console.log('tuCancel:click');
        close();
    });

    $('#tuOk').click(function ( event ) {
        event.preventDefault();
        console.log('tuOk:click');
        close();

        chrome.runtime.sendMessage({selectedElement: selectedElementSelector}, function (response) {
            console.log("Element Selected in content.js - callback:", response);
        });
    });

}

function close() {
    $('*').unbind('*');
    $('*').unbind('click');
    $('*').unbind('hover');
    $('*').unbind('unhover');
    $('*').unbind('mouseup');
    $('*').unbind('mousedown');
    $('*').unbind('mouseenter mouseleave')
    $('#tuWin').remove();
    $('.tu-selected').removeClass('tu-selected');
    $('.tu-active').removeClass('tu-active');
}

init();

