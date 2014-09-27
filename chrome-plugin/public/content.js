console.log('------------- ++++++++++++++ content.js loaded ------------- ++++++++++++++ ');
//alert('1')
//document.body.innerHTML = document.body.innerHTML.replace(new RegExp("uno", "g"), "dos");
function init() {
    if($('#tuWin').length > 0)
        return;

    $('body').append('<div id="tuWin" class="tu-window">' +
        'Selected element: <span id="selectedElement"></span><br/>' +
        '<a id="tuCancel" href="#">Cancel</a>' +
        '<a id="tuOk" href="#">Ok</a>' +
        '</div>');

    var $app = $('*:not(#tuWin *)');
    var selectedElementSelector = null;

    $app.hover(function () {
        $('.tu-active').removeClass('tu-active');
        $(this).addClass('tu-active');
    }, function () {
        $(this).removeClass('tu-active');
    });

    $app.unbind('click');
    $app.click(function ( event ) {
        event.preventDefault();

        console.log('clicked', $(this).prop('tagName'));
        $app.removeClass('tu-selected');
        var $this = $(this);
        $this.addClass('tu-selected');

        selectedElementSelector = $this.getSelector({ ignore: { classes: ['tu-active', 'tu-selected'] } });
        selectedElementSelector = selectedElementSelector.join("");
        console.log(selectedElementSelector);

        $("#selectedElement").text(selectedElementSelector);

        return false;
    });

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
    $('*').unbind('mouseenter mouseleave')
    $('#tuWin').remove();
    $('.tu-selected').removeClass('tu-selected');
    $('.tu-active').removeClass('tu-active');
}

init();