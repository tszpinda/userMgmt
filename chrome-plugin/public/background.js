chrome.browserAction.onClicked.addListener(function(tab) {
  console.log('------------- ++++++++++++++ background.js loaded ------------- ++++++++++++++ ');
  // No tabs or host permissions needed!

  if(!localStorage.authToken) {
  	chrome.tabs.create({url: "options.html"});	
  	localStorage.authToken = "abc";
  }
/*
  chrome.tabs.executeScript({
    file: 'content.js'
  });*/
});
chrome.runtime.onMessage.addListener(
  	function(request, sender, sendResponse) {
        if(!request)
            return;
        //console.log('onMessage: backgrund.js');
  		if(request.logout){
  			localStorage.removeItem('authToken');
  			return;
  		}else if(request.selectElement) {
            var tutorialId = request.tutorialId;
            var stepId = request.stepId;
            var stepText = request.stepText;

            localStorage.step = JSON.stringify({id : stepId,
                                               selector: "",
                                               text: stepText})

            localStorage.tutorialId = tutorialId;

            chrome.tabs.executeScript(null, {file: 'jquery-1.6.4.js'}, function () {
                chrome.tabs.executeScript(null, {file: 'selectorator.js'}, function () {
                    chrome.tabs.executeScript({file: 'content.js'});
                    sendResponse();
                });
            });
        }else if(request.selectedElement) {
            console.log('selected element', request.selectedElement);
            //alert('selected element:' + request.selectedElement);
            var step = JSON.parse(localStorage.step);
            step.selector = request.selectedElement;
            localStorage.removeItem("step");

            updateStep(step, sendResponse);
        }
	}
);	

function updateStep(step, callback) {
//    {"step":{"id":"540c566d421aa990e1000004","text":"Enter your password","selector":"#passwords","no":0,"tutorial":""}}
    console.log("updateStep", step);
    $.ajax({
        type: "PUT",
        contentType: "application/json; charset=utf-8",
        url: 'http://localhost:3000/steps/' + step.id,
        data: JSON.stringify({step: step}),
        dataType: "json",
        headers: {"auth-token":  localStorage.authToken},
        success: function (msg) {
            console.log('step updated', step);
            callback();
        },
        error: function (err){
            console.error(err);
            alert('Error');
            callback();
        }
    });
}

chrome.tabs.getSelected(null,function(tab) {
    var tablink = tab.url;
    console.log(tablink);
});