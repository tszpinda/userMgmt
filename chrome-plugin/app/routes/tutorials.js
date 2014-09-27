import AuthRoute from './auth';
import Ember from 'ember';

export default AuthRoute.extend({
    beforeModel: function() {
        if(localStorage.tutorialId)
            this.transitionTo('tutorial.view', localStorage.tutorialId);
        localStorage.removeItem('tutorialId');
    },

	model: function() {
		console.log("route:tutorials");

        //if()
        //this.transitionTo('posts');

	    /*
	    We don't want to show organizations in the list until they
	    are saved to the server
	     */
	     //return this.store.find('tutorial');

		console.log('checking...');
		var _that = this;
		
		if(chrome.tabs)
			return this.domainMatchingEntries();
		else
			return this.store.find('tutorial');
	},

	domainMatchingEntries : function() {
		var _that = this;
		return new Ember.RSVP.Promise(function(resolve, reject) {
			var tabUrl = null;
			console.log('before:chrome.tabs.query');


			chrome.tabs.query({'active': true, 'lastFocusedWindow': true}, function (tabs) {
	       		tabUrl = tabs[0].url;
	       		console.log('url tab is: ', tabUrl);
	       		_that.store.find('tutorial').then(function(tuts){
	       			tuts.forEach(function(obj){
	       				if (obj && tabUrl.indexOf(obj.get('domain')) === -1)
	       				{
	       					tuts.removeObject(obj);
	       				}
	       			});
	    			resolve(tuts);
	       		}).reject(function(e){
	       			alert('error processing query' + e);
	       			reject(e);
	       		});
	       	});
	       	console.log('after:chrome.tabs.query');
		});
	}

});
