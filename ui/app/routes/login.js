import Ember from 'ember';

export default Ember.Route.extend({
	/*
    The LoginRoute shares the model of the ApplicationRoute, which is
    the current session
   	*/
  	model: function() {
    	return this.modelFor('application');
  	},
  	actions : {
  		login: function() {
  			var _this = this;
  			this.currentModel.get('errors').clear();

  			this.currentModel.save().then(function(session){
  				console.log(session);
  				localStorage.authToken = session.get('authToken');
  				session.set('authToken', null);
  				/*The CSRF_TOKEN changes when the session is reset */
//  				App.CSRF_TOKEN = session.get('csrfToken');

	        	/*
	        	A saved transition allows redirecting the user to an
	        	autenticated page they requested before they were logged in
	         	*/
		        var transition = _this.controllerFor('session').get('savedTransition');
	        	if (transition) {
	        		transition.retry();
	        	}else{	        		
		        	_this.transitionTo('dashboard');
		        }
  			}, function(){				
  				_this.notifier.error("Login failed");
  			});
  		}
  	}
});
