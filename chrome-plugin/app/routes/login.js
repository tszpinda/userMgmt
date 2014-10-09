//import Ember from 'ember';
import UnauthRoute from './unAuth';

export default UnauthRoute.extend({
	/*
    The LoginRoute shares the model of the ApplicationRoute, which is
    the current session
   	*/
  	model: function() {
      console.log("login.js:model()");
    	return this.modelFor('application');
  	},
  	actions : {
      login: function() {
  			var _this = this;
  			this.currentModel.get('errors').clear();
        localStorage.removeItem("authToken");

  			this.currentModel.save().then(function(session){
          console.log("Logged in: ok");
  				localStorage.authToken = session.get('authToken');
  				/*The CSRF_TOKEN changes when the session is reset 
	        	A saved transition allows redirecting the user to an
	        	autenticated page they requested before they were logged in
	         	*/
		        var transition = _this.controllerFor('session').get('savedTransition');
	        	if (transition) {
	        		transition.retry();
	        	}else{	        		
		        	_this.transitionTo('tutorials');
		        }
  			}, function(error){				
          console.log("Unable to login", error);
  				_this.notifier.error("Login failed");
  			});
  		}
  	}
});